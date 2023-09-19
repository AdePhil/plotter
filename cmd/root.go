package cmd

import (
	"fmt"

	"github.com/AdePhil/plotter/prompts"
	"github.com/AdePhil/plotter/types"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "plotter",
	Short: "A simple utility library for plotting graphs",
	Long:  `A simple utility library for plotting graphs`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		chartType, err := prompts.GetChartType()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if chartType == types.BAR {
			BarRun()
		} else if chartType == types.LINE {
			LineRun(false)
		} else if chartType == types.SMOOTH_LINE {
			LineRun(true)
		} else if chartType == types.PIE {
			PieRun()
		}
	},
}
