package cmd

import (
	"fmt"
	"net/http"

	"github.com/AdePhil/plotter/charts"
	"github.com/AdePhil/plotter/prompts"
	"github.com/spf13/cobra"
)

var barCmd = &cobra.Command{
	Use:   "bar",
	Short: "plots a bar chart",
	Run: func(cmd *cobra.Command, args []string) {
		BarRun()
	},
}

func BarRun() {
	yValues, err := prompts.GetSeriesData()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	seriesData, err := ParseToBarValues(yValues)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	xValues, err := prompts.GetXData()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bar := charts.Bar{SeriesData: seriesData, XValues: ParseToXValues(xValues)}
	http.HandleFunc("/", bar.RenderGraph)
	fmt.Println("Check you graph here: http://localhost:6449/")
	http.ListenAndServe(":6449", nil)

}

func init() {
	RootCmd.AddCommand(barCmd)
}
