package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: "helloWorld",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Select Day",
			Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
				"Saturday", "Sunday"},
			Size:     7,
			HideHelp: true,
			Templates: &promptui.SelectTemplates{
				Active:   "🤔 {{ . | yellow }}",
				Inactive: "   {{ . | cyan }}",
				Selected: "😄 {{ . | green }}",
			},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %v\n", result)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
