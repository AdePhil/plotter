package prompts

import (
	"errors"
	"strconv"
	"strings"

	"github.com/tigergraph/promptui"
)

func GetXData() (string, error) {
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "X axes values x1, x2,.. (numbers|strings)",
		Validate: validate,
	}

	return prompt.Run()
}

func GetSeriesData() (string, error) {
	validate := func(input string) error {
		nums := strings.Split(input, ",")
		for _, numStr := range nums {
			_, err := strconv.Atoi(strings.Trim(numStr, " "))
			if err != nil {
				return errors.New("input is not in the correct format")
			}
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Y axes values x1, x2,.. (numbers)",
		Validate: validate,
	}
	return prompt.Run()
}

func GetPieData() (string, error) {
	validate := func(input string) error {
		nums := strings.Split(input, ",")
		for _, numStr := range nums {
			part := strings.Split(numStr, ":")
			if len(part) != 2 {
				return errors.New("input is not in the correct format")
			}
			_, err := strconv.Atoi(strings.Trim(part[1], " "))
			if err != nil {
				return errors.New("input is not in the correct format")
			}
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter data e.g Mon:30, Tues:40,.. ",
		Validate: validate,
	}
	return prompt.Run()
}

func GetTitle() (string, error) {
	validate := func(input string) error {
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Enter chart title ",
		Validate: validate,
	}
	return prompt.Run()
}

func GetAiPrompt() (string, error) {
	validate := func(input string) error {
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Type a prompt detailing the chart you want",
		Validate: validate,
	}
	return prompt.Run()
}
