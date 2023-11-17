package cmd

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/opts"
)

func ParseToPieValues(input string) ([]opts.PieData, error) {
	data := strings.Split(input, ",")
	results := []opts.PieData{}
	for _, p := range data {
		pieData := strings.Split(p, ":")
		if len(pieData) != 2 {
			return nil, errors.New("input format incorrect")
		}
		num, err := strconv.Atoi(strings.Trim(pieData[1], " "))
		if err != nil {
			return nil, errors.New("input format incorrect")
		}
		results = append(results, opts.PieData{Name: pieData[0], Value: num})
	}
	return results, nil
}
func ParseToPie(data []int, labels []string) []opts.PieData {
	results := []opts.PieData{}
	for i, _ := range data {
		results = append(results, opts.PieData{Name: labels[i], Value: data[i]})
	}
	return results
}
func ParseToBarValues(input string) ([]opts.BarData, error) {
	r1, err := getNumbers(input)
	if err != nil {
		return nil, err
	}

	return TransformType[int, opts.BarData](r1, func(num int) opts.BarData {
		return opts.BarData{Value: num}
	}), nil
}

func ParseToLineValues(input string) ([]opts.LineData, error) {
	nums, err := getNumbers(input)
	if err != nil {
		return nil, err
	}

	return TransformType[int, opts.LineData](nums, func(num int) opts.LineData {
		return opts.LineData{Value: num}
	}), nil
}

func ParseToLine(nums []int) []opts.LineData {
	return TransformType[int, opts.LineData](nums, func(num int) opts.LineData {
		return opts.LineData{Value: num}
	})
}

func TransformType[T comparable, K comparable](values []T, op func(T) K) []K {
	result := []K{}

	for _, v := range values {
		result = append(result, op(v))
	}

	return result
}

func ParseToXValues(input string) []interface{} {
	values := strings.Split(input, ",")
	return TransformType[string, interface{}](values, func(val string) interface{} {
		return interface{}(val)
	})
}

func generateBarData(data []int) []opts.BarData {
	items := []opts.BarData{}
	for _, num := range data {
		items = append(items, opts.BarData{Value: num})
	}
	return items
}

func getNumbers(input string) ([]int, error) {
	values := strings.Split(input, ",")
	result := []int{}

	for _, value := range values {
		num, err := strconv.Atoi(strings.Trim(value, " "))
		if err != nil {
			return nil, errors.New("input not a valid number")
		}
		result = append(result, num)
	}

	return result, nil
}

func extractValues(input string, key string) (value string) {
	regex := regexp.MustCompile(key + `: (.+)`)
	match := regex.FindStringSubmatch(input)
	if len(match) > 1 {
		value = match[1]
	}
	return value
}

func parseList(input string) []string {
	return regexp.MustCompile(`\s*,\s*`).Split(input, -1)
}

func parseIntList(input string) []int {
	var result []int
	values := parseList(input)
	for _, v := range values {
		num, err := strconv.Atoi(strings.Trim(v, " "))
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func ParseAiValues(input string)(string, string, []string, []int, []string,) {
	 chartType := extractValues(input, "Chart Type")
	 colors := parseList(extractValues(input, "Colors"))
	 seriesData := parseIntList(extractValues(input, "Series data"))
	 labels  := parseList(extractValues(input, "Labels"))
	 title      := extractValues(input, "Title")

  return chartType, title, colors, seriesData, labels
}
