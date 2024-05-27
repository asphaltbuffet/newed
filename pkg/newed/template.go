package newed

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	BaseDirName = "_base"
)

type Template struct {
	Name     string
	Dir      string
	Base     bool
	Sections []string
}

type Templates map[string]Template

func (t Templates) Load(dirs ...string) error {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelWarn}))

	for _, d := range dirs {
		entries, err := os.ReadDir(d)
		if err != nil {
			// return fmt.Errorf("reading directory: %w", err)
			fmt.Println("reading directory:", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				if err = t.Add(filepath.Join(d, entry.Name())); err != nil {
					logger.Info("loading templates", "dir", d, "error", err)
				}
			}
		}
	}

	return nil
}

func (t Templates) Add(dir string) error {
	info, err := os.Stat(dir)
	if err != nil {
		return err
	}

	// we only want to process directories as a template
	if !info.IsDir() {
		return fmt.Errorf("add %s: %w", dir, os.ErrInvalid)
	}

	name := filepath.Base(dir)

	subs, hasBase, err := readTemplateContent(dir)
	if err != nil {
		return err
	}

	t[name] = Template{
		Name:     name,
		Base:     has_base,
		Dir:      dir,
		Sections: subs,
	}

	return nil
}

func readTemplateContent(dir string) ([]string, bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, false, err
	}

	var has_base bool
	subs := []string{}

	for _, e := range entries {
		switch {
		case e.IsDir() && e.Name() == BaseDirName:
			has_base = true
		case e.IsDir():
			subs = append(subs, e.Name())
		default:
			// ignore files
		}
	}

	return subs, has_base, nil
}
