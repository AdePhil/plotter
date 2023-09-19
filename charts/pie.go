package charts

import (
	"io"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Pie struct {
	Data []opts.PieData
}

func (pie Pie) RenderGraph(w http.ResponseWriter, r *http.Request) {
	pie.drawLineChart(w)
}

func (pieData Pie) drawLineChart(w io.Writer) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{Title: "All countries in the world by population", Left: "center"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Y: "30"}),
	)

	pie.AddSeries("pie", pieData.Data)
	pie.Render(w)
}
