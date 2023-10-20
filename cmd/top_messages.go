/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// topMessagesCmd represents the topMessages command
var topMessagesCmd = &cobra.Command{
	Use:   "top-messages [input file]",
	Short: "list the top messages found in the logs (default 15)",
	Run: func(cmd *cobra.Command, args []string) {
		// read the input file
		inputFile := args[0]

		// read input file line by line
		f, err := os.Open(inputFile)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer f.Close()

		var topTen map[string]int = make(map[string]int)
		scanner := bufio.NewScanner(f)
		lineNb := 0
		for scanner.Scan() {
			line := scanner.Text()
			lineNb++

			lineMap := map[string]any{}
			// parse line as json
			err := json.Unmarshal([]byte(line), &lineMap)
			if err != nil {
				fmt.Printf("Error line %d: %s\n", lineNb, err)
				continue
			}

			msg, exist := lineMap["msg"]
			if !exist {
				continue
			}

			msgStr, ok := msg.(string)
			if !ok {
				continue
			}

			// add to map
			if _, exist := topTen[msgStr]; !exist {
				topTen[msgStr] = 0
			}
			topTen[msgStr]++
		}

		// sort the map based on the values
		// Create slice of key-value pairs
		pairs := make([][2]interface{}, 0, len(topTen))
		for k, v := range topTen {
			pairs = append(pairs, [2]interface{}{k, v})
		}

		// Sort slice based on values
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i][1].(int) > pairs[j][1].(int)
		})

		// Extract sorted keys
		keys := make([]string, len(pairs))
		for i, p := range pairs {
			keys[i] = p[0].(string)
		}

		nb := viper.GetInt("number")
		if nb <= 0 {
			nb = 10
		}
		// Print sorted map
		for i, k := range keys {
			if i >= nb {
				break
			}
			fmt.Printf("#%d - %s: %d\n", i+1, k, topTen[k])
		}

	},
}

func init() {
	rootCmd.AddCommand(topMessagesCmd)

	topMessagesCmd.Flags().IntP("number", "n", 15, "number to display")
	viper.BindPFlag("number", topMessagesCmd.Flags().Lookup("number"))
}
