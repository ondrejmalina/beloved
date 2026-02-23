package main

import (
	"fmt"
	"log"
	"os"

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

var addCmd = &cobra.Command{
	Use:   "add [path]",
	Short: "Add a new path",
	Long:  "Add a new path to the beloved",
	Run:   addPath,
}

func init() {
	rootCmd.AddCommand(addCmd)
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

func addPath(cmd *cobra.Command, args []string) {
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
