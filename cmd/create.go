package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/joevtap/go-serverless/utils"
	"github.com/joevtap/go-serverless/utils/prompt"
	"github.com/joevtap/scaffolder"
	"github.com/spf13/cobra"
)

/*
TODO: Add support for custom templates
TODO: Add support for git repository based templates
TODO: Add logs
TODO: Add support for path as project name
*/

var (
	// template  string
	name      string
	awsRegion string
)

type ProjectData map[string]string

func init() {
	// createCmd.Flags().StringVarP(&template, "template", "t", "", "Use a specific template (git repo)")
	createCmd.Flags().StringVarP(&name, "name", "", "", "Name of the project")
	createCmd.Flags().StringVarP(&awsRegion, "aws-region", "r", "", "AWS region to deploy to")
	createCmd.Flags().BoolP("init-git", "", false, "Initialize a git repository")
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:     "create [project-name]",
	Short:   "Used to create a new project using the Serverless Framework",
	Long:    `The "create" command is used to create a new project using the Serverless Framework.`,
	Example: `  gosls create`,
	Args:    cobra.MaximumNArgs(1),
	Run:     createCmdImpl,
}

func createCmdImpl(cmd *cobra.Command, args []string) {
	var err error

	// If the user provided a name in the command, use it
	if len(args) != 0 {
		name = args[0]
	}

	// If the user did not provide a name, prompt for it
	if name == "" {
		name, err = prompt.Name.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v \n", err)
			return
		}

		// If the user did not provide a name in the prompt, use the default
		if name == "" {
			name = "my-project"
		}
	}

	// If the user did not provide an AWS region, prompt for it
	if awsRegion == "" {
		_, awsRegion, err = prompt.AwsRegion.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v \n", err)
			return
		}
	}

	// Return if the project directory already exists
	if utils.PathExists(name) {
		fmt.Printf("Project directory %s already exists", name)
		return
	}

	spinner := utils.Spinner("🚀 Creating project...", "🎉 Project created!")
	spinner.Start()
	defer spinner.Stop()

	// Create the project directory structure
	tplPath := filepath.Join(utils.GetAppPath(), "tpl", "serverless_default.yaml")
	scaffolder.Scaffold(tplPath, name, ProjectData{
		"name":      name,
		"awsRegion": awsRegion,
	})

	// Initialize the project as a Go module
	initCmd := exec.Command("go", "mod", "init", name)
	initCmd.Dir = name
	if err := initCmd.Run(); err != nil {
		fmt.Printf("Error running go mod init: %v", err)
	}

	// Run go mod tidy
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = name
	if err := tidyCmd.Run(); err != nil {
		fmt.Printf("Error running go mod tidy: %v", err)
	}

	// Initialize the project as a git repository
	if cmd.Flag("init-git").Value.String() == "true" {
		gitInitCmd := exec.Command("git", "init")
		gitInitCmd.Dir = name
		if err := gitInitCmd.Run(); err != nil {
			fmt.Printf("Error running git init: %v", err)
		}
	}
}
