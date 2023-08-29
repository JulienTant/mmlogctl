/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean [input file] [output file]",
	Short: "clean logs",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]

		excludeLinesWith := viper.GetStringSlice("ExcludeLinesWith")
		if len(excludeLinesWith) > 0 {
			fmt.Println("Excluding lines with:")
			for _, s := range excludeLinesWith {
				fmt.Println(" ", s)
			}
		}

		// Open the input file
		input, err := os.Open(inputFile)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer input.Close()

		// Create a new scanner for the input file
		scanner := bufio.NewScanner(input)

		// Create the output file
		output, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer output.Close()

		// Create a new writer for the output file
		writer := bufio.NewWriter(output)

		readLines := 0
		wroteLines := 0
		// Scan each line of the file
		for scanner.Scan() {
			line := scanner.Text()
			readLines++

			// if lines does not have a { or is empty, skip it
			if len(line) == 0 || strings.Index(line, "{") == -1 {
				continue
			}

			// if the line does not start with a {, remove the text before the {
			if line[0] != '{' {
				line = line[strings.Index(line, "{"):]
			}

			// if we find a line to skip, skip it
			for _, skip := range excludeLinesWith {
				if strings.Index(line, strings.TrimSpace(skip)) != -1 {
					line = ""
					break
				}
			}
			if len(line) == 0 {
				continue
			}

			wroteLines++
			fmt.Fprintln(writer, line)
		}

		// Flush the buffer to ensure all data has been written to the file
		writer.Flush()

		fmt.Printf("Read %d lines, wrote %d lines\n", readLines, wroteLines)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().StringArrayP("ExcludeLinesWith", "e", []string{
		"Worker: Job is complete",
		"No notification data available",
		"Notification will be sent",
		"Notification sent",
		"Notification not sent",
		"Notification received",
		"websocket.slow",
	}, "exclude lines containing one of those strings")
	viper.BindPFlag("ExcludeLinesWith", cleanCmd.Flags().Lookup("ExcludeLinesWith"))
}
