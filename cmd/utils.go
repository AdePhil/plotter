package cmd

import (
	"errors"
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
	r1, err := getNumbers(input)
	if err != nil {
		return nil, err
	}

	return TransformType[int, opts.LineData](r1, func(num int) opts.LineData {
		return opts.LineData{Value: num}
	}), nil
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
