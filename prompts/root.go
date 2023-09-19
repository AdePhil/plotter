package prompts

import (
	"github.com/AdePhil/plotter/types"
	"github.com/manifoldco/promptui"
)

func GetChartType() (string, error) {

	prompt := promptui.Select{
		Label: "Select Chart Type",
		Items: []string{types.BAR, types.LINE, types.SMOOTH_LINE, types.PIE},
	}

	_, result, error := prompt.Run()

	return result, error
}
