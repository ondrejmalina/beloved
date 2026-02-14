package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

const default_cfg_file string = "beloved.cfg"

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
	path 	string
	beloved []string
}

func getDefaultConfigPath() (string, error) {
	dd, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dd, "beloved")
	return path, nil
}

func createDefaultConfig(path string) error {
	err := os.MkdirAll(path, 0660)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(path, default_cfg_file))
	f.Close()
	if err != nil {
		return err
	}
}

// Reads cfg file and return a slice of beloved paths.
func readConfig(path string) ([]string, error) {
	// TODO: Add validation of correct cfg format
	paths := []string{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))

	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}

	return paths, nil
}

// Searches for config on path. If the path is an empty string, loads the
// config from the OS default config path. If it does not exist on the
// default path, it creates it.
func loadConfig(path string) ([]string, error) {
	if path == "" {
		cfg, err := loadDefaultConfig()
		if err != nil {
			return nil, err
		}
	}

	cfg, err := readConfig(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config on path %s: %w", path, err)
	}
	return cfg, nil

}

func loadDefaultConfig() ([]string, error) {
	dp, err := getDefaultConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get default cfg path: %w", err)
	}

	fp := filepath.Join(dp, default_cfg_file)
	cfg, err := readConfig(fp)

	if err == nil {
		return cfg, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		err = createDefaultConfig(fp)
		if err != nil {
			return nil, fmt.Errorf("failed to create default config: %w", err)
		}
		return []string{}, nil
	}

	return nil, fmt.Errorf("failed to read config at %s: %w", fp, err)
}