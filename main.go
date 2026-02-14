package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

const (
	default_cfg_file string = "beloved.cfg"
	default_cfg_folder string = "beloved"
)

var (
	testPath string
)

var rootCmd = &cobra.Command{
	Use:   "beloved",
	Short: "A simple CLI tool for tracking favorite (or 'beloved') filesystem paths",
	Long: `beloved is a CLI tool for tracking of favorite filesystem paths.
You can add or remove favorite paths to it using commands.`,
	Run: greeting,
}

func greeting(cmd *cobra.Command, args []string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Beloved ❤️").
				Options(huh.NewOptions("United States", "Canada", "Mexico")...).
				Value(&testPath),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type config struct {
	path    string // path to the config file
	beloved []string
}

// Searches for config on path. If the path is an empty string, loads the
// config from the OS default config path. If it does not exist on the
// default path, it creates it.
func (c *config) loadConfig() error {
	if c.path == "" {
		err := c.loadDefaultConfig(); err != nil {
			return err
		}
	}

	err := c.readConfig(); err != nil {
		return fmt.Errorf("failed to read config on path %s: %w", c.path, err)
	}
	return nil

}

func (c *config) loadDefaultConfig() error {
	err := c.getDefaultConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get default cfg path: %w", err)
	}

	err = c.readConfig(); err == nil {
		return nil
	}

	if errors.Is(err, os.ErrNotExist) {
		err = c.createDefaultConfig(); err != nil {
			return fmt.Errorf("failed to create default config file: %w", err)
		}
	}

	err = c.readConfig(); err != nil {
		return fmt.Errorf("failed to read config on path %s: %w", c.path, err)
	}
	return nil
}

func (c *config) getDefaultConfigPath() error {
	dd, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	c.path := filepath.Join(dd, default_cfg_folder, default_cfg_file)
	return nil
}

// Reads cfg file from path to configs slice of beloved paths.
func (c *config) readConfig() error {
	// TODO: Add validation of correct cfg format
	data, err := os.ReadFile(c.path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		c.beloved = append(c.beloved, scanner.Text())
	}
	return nil
}

func (c *config) createDefaultConfig() error {
	err := os.MkdirAll(filepath.Dir(c.path), 0660)
	if err != nil {
		return err
	}

	f, err := os.Create(c.path)
	f.Close()
	if err != nil {
		return err
	}
}

