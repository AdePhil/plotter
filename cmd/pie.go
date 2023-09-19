package cmd

import (
	"fmt"
	"net/http"

	"github.com/AdePhil/plotter/charts"
	"github.com/AdePhil/plotter/prompts"
	"github.com/spf13/cobra"
)

var pieCmd = &cobra.Command{
	Use:   "pie",
	Short: "plots a pie chart",
	Run: func(cmd *cobra.Command, args []string) {
		PieRun()
	},
}

func PieRun() {
	title, _ := prompts.GetTitle()
	input, err := prompts.GetPieData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	pieData, err := ParseToPieValues(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	line := charts.Pie{Title: title, Data: pieData}
	http.HandleFunc("/", line.RenderGraph)
	fmt.Println("Check you graph here: http://localhost:6449/")
	http.ListenAndServe(":6449", nil)

}

func init() {
	RootCmd.AddCommand(pieCmd)
}
