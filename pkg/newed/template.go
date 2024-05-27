package newed

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/asphaltbuffet/newed/internal/config"
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

// New creates a new Templates object from the provided config.
func New(cfg *config.Config) (Templates, error) {
	t := make(Templates)

	if err := t.Load(cfg.GetTemplateDirs()...); err != nil {
		return nil, err
	}

	return t, nil
}

// Load reads the provided directories and adds any valid templates to the Templates object.
func (t Templates) Load(dirs ...string) error {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelWarn}))

	for _, d := range dirs {
		entries, err := os.ReadDir(d)
		if err != nil {
			continue
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

// Add the provided template to Templates object if new; otherwise, replaces existing template
// with the same name.
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
		Base:     hasBase,
		Dir:      filepath.Dir(dir),
		Sections: subs,
	}

	return nil
}

func readTemplateContent(dir string) ([]string, bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, false, err
	}

	names := make([]string, 0, len(entries))
	hasBase := false

	for _, e := range entries {
		switch {
		case e.IsDir() && e.Name() == BaseDirName:
			hasBase = true
			names = append(names, e.Name())
		case e.IsDir():
			names = append(names, e.Name())
		default:
			// ignore files
		}
	}

	return names, hasBase, nil
}

func Expand(templates ...string) []string {
	expanded := []string{}

	for _, t := range templates {
		// optional templates show up as `base+opt1+opt2`
		additions := strings.Split(t, "+")

		base := additions[0]
		// always add the base template
		// FIX: handle if there is no base template? or do we just assume it will fail quietly?
		expanded = append(expanded, filepath.Join(base, BaseDirName))

		for i := 1; i < len(additions); i++ {
			expanded = append(expanded, filepath.Join(base, additions[i]))
		}
	}

	return expanded
}
