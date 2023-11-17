package charts

import (
	"io"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Line struct {
	Colors []string
	XValues    []interface{}
	SeriesData []opts.LineData
	Smooth     bool
	Title      string
}

func (line Line) RenderGraph(w http.ResponseWriter, r *http.Request) {
	line.drawLineChart(w)
}

func (lineData Line) drawLineChart(w io.Writer) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithColorsOpts(line.Colors),
		charts.WithTitleOpts(opts.Title{Title: lineData.Title, Left: "center"}),
		charts.WithLegendOpts(opts.Legend{Show: false}))

	line.SetXAxis(lineData.XValues).AddSeries("Testing", lineData.SeriesData).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: lineData.Smooth}))

	line.Render(w)
}
