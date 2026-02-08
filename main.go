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

func getDefaultConfigPath() (string, error) {
	dd, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dd, "beloved")
	return path, nil
}

func createDefaultConfig(path string) (string, error) {
	err := os.MkdirAll(path, 0660)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(filepath.Join(path, default_cfg_file))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f.Name(), nil
}

func openConfig(path string) (*os.File, error) {
	cfg, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return cfg, nil
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
		p, err := getDefaultConfigPath()
		if err != nil {
		}
	}

	return nil, nil
}

func loadDefaultConfig() error {
	p, err := getDefaultConfigPath()
	if err != nil {
		return err
	}

	fp := filepath.Join(p, default_cfg_file)

	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

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
