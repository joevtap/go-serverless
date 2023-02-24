package cmd

import (
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	template string
)

func init() {
	CreateCmd.Flags().StringVarP(&template, "template", "t", "", "Use a specific template (git repo)")
	rootCmd.AddCommand(CreateCmd)
}

var CreateCmd = &cobra.Command{
	Use:     "create [project-name]",
	Short:   "Used to create a new project using the Serverless Framework",
	Long:    `The "create" command is used to create a new project using the Serverless Framework.`,
	Example: `  gosls create`,
	Args:    cobra.MaximumNArgs(1),
	Run:     CreateCmdImpl,
}

func CreateCmdImpl(cmd *cobra.Command, args []string) {
	var err error
	var name string

	if len(args) != 0 {
		name = args[0]
	} else {
		name, err = namePrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v \n", err)
		}
	}

	fmt.Printf("Project name: %v\n", name)
}

var namePrompt = promptui.Prompt{
	Label:   "Give your project a name",
	Default: "my-project",
	Validate: func(input string) error {
		if len(input) < 3 {
			return fmt.Errorf("project name must be at least 3 characters")
		}

		if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(input) {
			return fmt.Errorf("project name must be alphanumeric")
		}

		return nil
	},
}
