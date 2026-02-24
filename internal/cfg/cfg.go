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

// Check if application config exists in OS default configuration dir.
// Create it if not.
func Init() (*Config, error) {
	cfg := Config{}

	dd, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("user config dir unavailable: %w", err)
	}
	cfg.Path = filepath.Join(dd, defaultCfgFolder, defaultCfgFile)

	if _, err := os.Stat(cfg.Path); errors.Is(err, os.ErrNotExist) {
		if err := cfg.new(); err != nil {
			return nil, fmt.Errorf("failed to create config: %w", err)
		}
	}

	return &cfg, nil
}

func (c *Config) new() error {
	if err := os.MkdirAll(filepath.Dir(c.Path), 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}
	f, err := os.Create(c.Path)
	if err != nil {
		return fmt.Errorf("create file failed: %w", err)
	}
	return f.Close()
}

// Load all paths from the config file into the Config struct
func (c *Config) Load() error {
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

// Add path to the Config file. Create the file if it does not exist
func (c *Config) Add(path string) (int, error) {
	// TODO: Update the permissions
	f, err := os.OpenFile(c.Path, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return 0, fmt.Errorf("failed to open the config file: %w", err)
	}


	n, err := f.Write([]byte(path+"\n"))
	if err != nil {
		return 0, fmt.Errorf("failed to write to the config file: %w", err)
	}

	return n, err
}
