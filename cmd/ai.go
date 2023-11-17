package cmd

import (
	// "github.com/AdePhil/plotter/charts"
	// "github.com/AdePhil/plotter/prompts"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AdePhil/plotter/charts"
	"github.com/AdePhil/plotter/prompts"
	"github.com/AdePhil/plotter/types"
	"github.com/briandowns/spinner"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var aiCmd = &cobra.Command{
	Use:   "ai",
	Short: "Type a prompt detailing the chart you want",
	Run: func(cmd *cobra.Command, args []string) {
		AiRun()
	},
}

func AiRun() {
	user_prompt, _ := prompts.GetAiPrompt()
	token :=  os.Getenv("OPENAI_TOKEN")
	client := openai.NewClient(token)
	const user_input = "Draw a bar chart with color red of my calorie in take from monday - friday with this data 20, 30, 40, 50, 60"
	
	 req := openai.ChatCompletionRequest{
	 	Model:     openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful assistant designed to extract chart information. Extract chart details, emphasizing on these details: color, series data, labels, and title. Show the series data as a comma separated list of numbers. Always make sure that the series data are numbers. Remember convert the units. Show the labels as comma separated list of strings. Always include the word 'Chart' as a suffix to the chart type. Give me only the key details don't plot the graph.  The numbers should not be comma separated. Always return the results in this format:\n" +
				"- Chart Type: value extracted"+
				"- Colors: values extracted"+
				"- Series data: values extracted" +
				"- Labels: values extracted" +
				"- Title: value extracted" +
				"Make sure to capitalize the values except for Colors which should be in lowercase. Do not put string values in quotes. If the user doesn't select chart type default to the value 'Bar Chart'.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: user_prompt,
			},
		},
		Stream: false,
	 	MaxTokens:    200,
		Temperature: 0,
		TopP: 1,
	 }
	
	 fmt.Println()
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)  
	s.Start()
	resp, err := client.CreateChatCompletion(context.Background(), req)
	s.Stop() 
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	response := resp.Choices[0].Message.Content;
	chartType, title, colors, seriesData, labels := ParseAiValues(response)

	labelsInterface := TransformType[string, interface{}](labels, func(val string) interface{} {
		return interface{}(val)
	})

	var handler func(http.ResponseWriter, *http.Request)
		if chartType == types.BAR {
			bar := charts.Bar{Colors: colors, Title: title, SeriesData: generateBarData(seriesData), XValues: labelsInterface}
			handler = bar.RenderGraph
		} else if chartType == types.LINE {
			line := charts.Line{Colors: colors, Title: title, SeriesData: ParseToLine(seriesData), XValues: labelsInterface, Smooth: false}
			handler = line.RenderGraph
		} else if chartType == types.SMOOTH_LINE {
			line := charts.Line{Colors: colors, Title: title, SeriesData: ParseToLine(seriesData), XValues: labelsInterface, Smooth: true}
			handler = line.RenderGraph
		} else if chartType == types.PIE {
			line := charts.Pie{Colors: colors, Title: title, Data: ParseToPie(seriesData, labels)}
			handler = line.RenderGraph
		}

	http.HandleFunc("/", handler)
	// fmt.Println(chartType, colors, seriesData, labels, title)
	fmt.Println("Check you graph here http://localhost:6449/")
	http.ListenAndServe(":6449", nil)
}

func init() {
	RootCmd.AddCommand(aiCmd)
}