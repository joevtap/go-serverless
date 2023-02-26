package prompt

import (
	"strings"

	"github.com/manifoldco/promptui"
)

type awsRegion struct{}

var AwsRegion awsRegion

func (awsRegion) Run() (int, string, error) {

	regions := []string{
		"af-south-1",
		"ap-east-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-northeast-3",
		"ap-south-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ca-central-1",
		"eu-central-1",
		"eu-north-1",
		"eu-south-1",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"me-south-1",
		"sa-east-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"cn-north-1",
		"cn-northwest-1",
	}

	prompt := promptui.Select{
		Size:              5,
		StartInSearchMode: true,
		Label:             "Select an AWS Region",
		Items:             regions,
		Searcher: func(input string, index int) bool {
			item := regions[index]
			name := strings.Replace(strings.ToLower(item), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		},
		Templates: &promptui.SelectTemplates{
			Active:   "🤔 {{ . | cyan }}",
			Inactive: "   {{ . | cyan }}",
			Selected: "🌎 {{ . | green }}",
		},
	}

	return prompt.Run()
}
