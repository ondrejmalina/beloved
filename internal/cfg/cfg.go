package cfg

import (
	"bufio"
	"bytes"
	"errors"
	"path/filepath"
	"fmt"
	"os"
)

const (
	default_cfg_file string = "beloved.cfg"
	default_cfg_folder string = "beloved"
)

type Config struct {
	Path string
	Beloved []string
}

func CreateConfig(path string) Config {
	return Config{path, []string{}}
}

// Searches for config on path. If the path is an empty string, loads the
// config from the OS default config path. If it does not exist on the
// default path, it creates it.
func (c *Config) LoadConfig() error {
	if c.path == "" {
		if err := c.loadDefaultConfig(); err != nil {
			return fmt.Errorf("failed to read default config: %w", err)
		}
	}

	if err := c.readConfig(); err != nil {
		return fmt.Errorf("failed to read config on path %s: %w", c.path, err)
	}
	return nil

}

func (c *Config) loadDefaultConfig() error {
	err := c.getDefaultConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get default cfg path: %w", err)
	}

	if err = c.readConfig(); err == nil {
		return nil
	}

	if errors.Is(err, os.ErrNotExist) {
		if err = c.createDefaultConfig(); err != nil {
			return fmt.Errorf("failed to create default config file: %w", err)
		}
	}

	if err = c.readConfig(); err != nil {
		return fmt.Errorf("failed to read config on path %s: %w", c.path, err)
	}
	return nil
}

// Return OS default config path
func (c *Config) getDefaultConfigPath() error {
	dd, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to read os default config path: %w", err)
	}
	c.path = filepath.Join(dd, default_cfg_folder, default_cfg_file)
	return nil
}

// Reads cfg file from path to configs slice of beloved paths.
func (c *Config) readConfig() error {
	// TODO: Add validation of correct cfg format
	data, err := os.ReadFile(c.path)
	if err != nil {
		return fmt.Errorf("failed to read config %s: %w", c.path, err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		c.beloved = append(c.beloved, scanner.Text())
	}
	return nil
}

func (c *Config) createDefaultConfig() error {
	err := os.MkdirAll(filepath.Dir(c.path), 0660)
	if err != nil {
		return fmt.Errorf("failed to create default config folder %s: %w", filepath.Dir(c.path), err)
	}

	f, err := os.Create(c.path)
	if err != nil {
		return fmt.Errorf("failed to create default config file %s: %w", c.path, err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("failed to close the default config file %s: %w", c.path, err)
	}
	return nil
}

