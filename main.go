package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
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

func getDefaultCfgPath() string, error {
	dd, err := os.UserConfigDir()
	if err != nil {
		return nil ,err
	}
	path := filepath.Join(dd, "beloved")
	return path, nil
}

func createDefaultCfgFile(path string) string, error {
	err := os.MkdirAll(path, 0660)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(filepath.Join(path, "beloved.cfg"))
	if err != nil {
		return nil, err
	}

	return f, nil
}

func openCfg(path string) (*os.File, error) {
	cfg, err := os.Open(path)
	
	if err != nil {
		return nil, err
	}

	return cfg, nil
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