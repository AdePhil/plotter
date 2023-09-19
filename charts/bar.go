package charts

import (
	"io"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Bar struct {
	XValues    []interface{}
	SeriesData []opts.BarData
}

func (bar Bar) RenderGraph(w http.ResponseWriter, r *http.Request) {
	bar.drawBarChart(w)
}

func (barData Bar) drawBarChart(w io.Writer) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "All countries in the world by population", Left: "center"}), charts.WithLegendOpts(opts.Legend{Show: false}))

	bar.SetXAxis(barData.XValues).AddSeries("", barData.SeriesData)
	bar.Render(w)
}
