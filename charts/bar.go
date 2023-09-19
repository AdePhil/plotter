package charts

import (
	"io"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Bar struct {
	XValues    []interface{}
	SeriesData []opts.BarData
	Title      string
}

func (bar Bar) RenderGraph(w http.ResponseWriter, r *http.Request) {
	bar.drawBarChart(w)
}

func (barData Bar) drawBarChart(w io.Writer) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: barData.Title, Left: "center"}), charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}), charts.WithLegendOpts(opts.Legend{Show: false}))

	bar.SetXAxis(barData.XValues).AddSeries("", barData.SeriesData)
	bar.Render(w)
}
