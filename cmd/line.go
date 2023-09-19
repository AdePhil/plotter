package cmd

import (
	"fmt"
	"net/http"

	"github.com/AdePhil/plotter/charts"
	"github.com/AdePhil/plotter/prompts"
	"github.com/spf13/cobra"
)

var lineCmd = &cobra.Command{
	Use:   "line",
	Short: "plots a line chart",
	Run: func(cmd *cobra.Command, args []string) {
		LineRun(false)
	},
}

func LineRun(isSmooth bool) {
	title, _ := prompts.GetTitle()
	yValues, err := prompts.GetSeriesData()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	seriesData, err := ParseToLineValues(yValues)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	xValues, err := prompts.GetXData()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	line := charts.Line{Title: title, SeriesData: seriesData, XValues: ParseToXValues(xValues), Smooth: isSmooth}
	http.HandleFunc("/", line.RenderGraph)
	fmt.Println("Check you graph here: http://localhost:6449/")
	http.ListenAndServe(":6449", nil)

}

func init() {
	RootCmd.AddCommand(lineCmd)
}
