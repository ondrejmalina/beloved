package cfg

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	defaultCfgFile   = "beloved.cfg"
	defaultCfgFolder = "beloved"
)

type Config struct {
	Path    string
	Beloved []string
}

// New returns a Config with the path set. Use an empty string for default OS path.
func New(path string) *Config {
	return &Config{Path: path}
}

// Load populates the Config. It handles default pathing and file creation.
func (c *Config) Load() error {
	if c.Path == "" {
		dd, err := os.UserConfigDir()
		if err != nil {
			return fmt.Errorf("user config dir unavailable: %w", err)
		}
		c.Path = filepath.Join(dd, defaultCfgFolder, defaultCfgFile)
	}

	err := c.read()
	if errors.Is(err, os.ErrNotExist) {
		if err := c.init(); err != nil {
			return err
		}
		return c.read()
	}
	return err
}

func (c *Config) read() error {
	f, err := os.Open(c.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	// NOTE: Clear existing data to prevent duplicates on re-load
	c.Beloved = []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			c.Beloved = append(c.Beloved, line)
		}
	}
	return scanner.Err()
}

func (c *Config) init() error {
	if err := os.MkdirAll(filepath.Dir(c.Path), 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}
	f, err := os.Create(c.Path)
	if err != nil {
		return fmt.Errorf("create file failed: %w", err)
	}
	return f.Close()
}
