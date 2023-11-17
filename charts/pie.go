package charts

import (
	"io"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Pie struct {
	Colors []string
	Data  []opts.PieData
	Title string
}

func (pie Pie) RenderGraph(w http.ResponseWriter, r *http.Request) {
	pie.drawLineChart(w)
}

func (pieData Pie) drawLineChart(w io.Writer) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithColorsOpts(pieData.Colors),
	  charts.WithTooltipOpts(opts.Tooltip{Show: true, Formatter: "{b}: {c} ({d}%)"}),
		charts.WithTitleOpts(opts.Title{Title: pieData.Title, Left: "center"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Y: "25"}),
	)

	pie.AddSeries("pie", pieData.Data)
	pie.Render(w)
}
