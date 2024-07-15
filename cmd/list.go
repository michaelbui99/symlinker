/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/michaelbui99/symlinker/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all links defined in SymlinkerFile",
	Long:  `Lists all links defined in SymlinkerFile. Symlinker does not keep track of state, so they symlinks may or may not already exist.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		files, err := internal.ListFiles(cwd)
		if err != nil {
			return err
		}

		symlinkerFileName, err := internal.FindSymlinkerFile(files)
		if err != nil {
			return err
		}

		symlinkerFile, err := internal.ParseSymlinkerFile(fmt.Sprintf("./%s", symlinkerFileName))
		if err != nil {
			return err
		}

		// TODO: Add support for "MODULE" link type.
		b := color.New(color.FgBlue, color.Bold)
		r := color.New(color.FgRed, color.Bold)
		for _, link := range symlinkerFile.Links {
			r.Printf("%s: ", link.Name)
			b.Printf("\"%s\"", link.Source)
			fmt.Printf(" --> ")
			b.Printf("\"%s\"\n", link.Target)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
