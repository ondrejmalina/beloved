package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/ondrejmalina/beloved/internal/cfg"
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

var addCmd = &cobra.Command{
	Use:   "add [path]",
	Short: "Add a new path",
	Long:  "Add a new path to the beloved",
	Args:  cobra.MinimumNArgs(1),
	Run:   addPath,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func greeting(cmd *cobra.Command, args []string) {
	c, err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}
	err = c.Load()
	if err != nil {
		log.Fatal(err)
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Beloved ❤️").
				Options(huh.NewOptions(c.Beloved...)...).
				Value(&testPath),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.Chdir(testPath); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\ncd%s ", testPath)
}

func addPath(cmd *cobra.Command, args []string) {
	path := args[0]
	c, err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}
	c.Add(path)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
