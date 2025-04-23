// Package main calculates the size of the modcache for a module.
package main

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"maps"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const (
	// KiB kibibyte.
	KiB int64 = 1024
	// MiB mebibyte.
	MiB = 1024 * KiB
	// GiB gibibyte.
	GiB = 1024 * MiB
	// TiB tebibyte.
	TiB = 1024 * GiB
	// PiB pebibyte.
	PiB = 1024 * TiB
	// EiB exbibyte.
	EiB = 1024 * PiB
)

const (
	otherDirs = "¤¤¤"
	cacheDir  = "cache"
)

// Options CLI options.
type Options struct {
	JSON bool
	Mode string
}

// Report represents sizes information.
type Report struct {
	Size    int64         `json:"size"`
	Files   int64         `json:"files"`
	Dirs    int64         `json:"dirs"`
	Modules []*ModuleSize `json:"modules,omitempty"`
}

func (r Report) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("Size: %s\n", format(r.Size)))
	builder.WriteString(fmt.Sprintf("%d files, %d directories\n", r.Files, r.Dirs))

	return builder.String()
}

// Module information about module.
// https://go.dev/ref/mod#go-mod-download
type Module struct {
	Path     string // module path
	Query    string // version query corresponding to this version
	Version  string // module version
	Error    string // error loading module
	Info     string // absolute path to cached .info file
	GoMod    string // absolute path to cached .mod file
	Zip      string // absolute path to cached .zip file
	Dir      string // absolute path to cached source root directory
	Sum      string // checksum for path, version (as in go.sum)
	GoModSum string // checksum for go.mod (as in go.sum)
	Origin   any    // provenance of module
	Reuse    bool   // reuse of old module info is safe
}

// ModuleSize represents the size of module sources.
type ModuleSize struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func main() {
	opts := Options{}

	flag.StringVar(&opts.Mode, "mode", "download", "Mode: download, tidy")
	flag.BoolVar(&opts.JSON, "json", false, "JSON output format")

	flag.Parse()

	switch opts.Mode {
	case "download", "tidy":
	default:
		log.Fatalf("Unknown mode: %s\n", opts.Mode)
	}

	ctx := context.Background()

	err := run(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, opts Options) error {
	temp, err := os.MkdirTemp("", "modcache-*")
	if err != nil {
		return fmt.Errorf("create temp directory: %w", err)
	}

	defer func() {
		cmd := exec.CommandContext(ctx, "go", "clean", "-modcache")

		cmd.Env = append(os.Environ(), fmt.Sprintf("GOMODCACHE=%s", temp))

		errRun := cmd.Run()
		if errRun != nil {
			log.Printf("%s (%s): %v", strings.Join(cmd.Args, " "), temp, errRun)
		}
	}()

	switch opts.Mode {
	case "tidy":
		return modTidy(ctx, opts, temp)

	default:
		return modDownload(ctx, opts, temp)
	}
}

func modTidy(ctx context.Context, opts Options, temp string) error {
	cmd := exec.CommandContext(ctx, "go", "mod", "tidy")

	cmd.Env = append(os.Environ(), fmt.Sprintf("GOMODCACHE=%s", temp))

	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s (%s): %w", strings.Join(cmd.Args, " "), temp, err)
	}

	root, err := os.OpenRoot(temp)
	if err != nil {
		return fmt.Errorf("open root %s: %w", temp, err)
	}

	report := Report{}

	err = fs.WalkDir(root.FS(), ".", func(_ string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			report.Dirs++
		} else {
			report.Files++
		}

		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("entry info: %w", err)
		}

		report.Size += info.Size()

		return nil
	})
	if err != nil {
		return fmt.Errorf("walk dir %s: %w", temp, err)
	}

	if opts.JSON {
		return json.NewEncoder(os.Stdout).Encode(report)
	}

	fmt.Print(report.String())

	return nil
}

func modDownload(ctx context.Context, opts Options, temp string) error {
	cmd := exec.CommandContext(ctx, "go", "mod", "download", "-json")

	cmd.Env = append(os.Environ(), fmt.Sprintf("GOMODCACHE=%s", temp))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s (%s): %w", strings.Join(cmd.Args, " "), temp, err)
	}

	decoder := json.NewDecoder(bytes.NewReader(output))

	var data []Module

	for decoder.More() {
		var m Module

		err = decoder.Decode(&m) //nolint:musttag // Structure from Go itself.
		if err != nil {
			return fmt.Errorf("decode: %w", err)
		}

		data = append(data, m)
	}

	dirs := make(map[string]*ModuleSize, len(data)+2)

	dirs[otherDirs] = &ModuleSize{Name: "Others"}
	dirs[cacheDir] = &ModuleSize{Name: cacheDir}

	for _, datum := range data {
		rel, _ := filepath.Rel(temp, datum.Dir)
		dirs[rel] = &ModuleSize{Name: rel}
	}

	root, err := os.OpenRoot(temp)
	if err != nil {
		return fmt.Errorf("open root %s: %w", temp, err)
	}

	report := Report{}
	curr := otherDirs

	err = fs.WalkDir(root.FS(), ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			report.Dirs++

			if _, ok := dirs[path]; ok {
				curr = path
			}

			if curr != otherDirs && !strings.HasPrefix(path, curr) {
				curr = otherDirs
			}
		} else {
			report.Files++
		}

		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("entry info: %w", err)
		}

		report.Size += info.Size()

		dirs[curr].Size += info.Size()

		return nil
	})
	if err != nil {
		return fmt.Errorf("walk dir %s: %w", temp, err)
	}

	if opts.JSON {
		sortedDirs := slices.SortedFunc(maps.Values(dirs), func(ms1 *ModuleSize, ms2 *ModuleSize) int {
			return cmp.Compare(ms2.Size, ms1.Size)
		})
		report.Modules = sortedDirs

		return json.NewEncoder(os.Stdout).Encode(report)
	}

	if len(dirs) > 2 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Module", "Size"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")

		sortedDirs := slices.SortedFunc(maps.Values(dirs), func(ms1 *ModuleSize, ms2 *ModuleSize) int {
			return cmp.Compare(ms2.Size, ms1.Size)
		})

		for _, d := range sortedDirs {
			table.Append([]string{d.Name, format(d.Size)})
		}

		table.Render()
	}

	fmt.Print(report.String())

	return nil
}

func format(size int64) string {
	switch {
	case size > EiB:
		return fmt.Sprintf("%.3f EiB", float64(size)/float64(EiB))
	case size > PiB:
		return fmt.Sprintf("%.3f PiB", float64(size)/float64(PiB))
	case size > TiB:
		return fmt.Sprintf("%.3f TiB", float64(size)/float64(TiB))
	case size > GiB:
		return fmt.Sprintf("%.3f GiB", float64(size)/float64(GiB))
	case size > MiB:
		return fmt.Sprintf("%.3f MiB", float64(size)/float64(MiB))
	case size > KiB:
		return fmt.Sprintf("%.3f KiB", float64(size)/float64(KiB))
	case size == 0:
		return ""
	default:
		return fmt.Sprintf("%d B", size)
	}
}
