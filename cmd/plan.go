/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/michaelbui99/symlinker/internal"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Shows the changes Symlinker will perfrom on 'symlinker up'",
	Long:  `Shows the changes Symlinker will perform on 'symlinker up'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := internal.ParseSymlinkerFile("./SymlinkerFile.yaml")
		if err != nil {
			return err
		}
		log.Printf("%v\n", f)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(planCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// planCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// planCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
