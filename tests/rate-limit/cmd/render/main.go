package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type Record struct {
	Time     time.Time
	Rate     int
	Count    int
	Offset   time.Duration
	Duration time.Duration
	Status   int
}

func main() {
	if err := start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() error {
	var filename, output, metric string
	var padding time.Duration
	var printInfo bool

	flag.StringVar(&filename, "f", "", "Filename to open")
	flag.StringVar(&output, "o", "", "Output filename")
	flag.StringVar(&metric, "t", "rate", "Render metric (rate, count, duration)")
	flag.BoolVar(&printInfo, "info", false, "Print request info")
	flag.DurationVar(&padding, "p", padding, "Time padding (total duration)")
	flag.Parse()

	// Read JSON data from the file
	jsonData, err := readJSONFile(filename, padding, metric)
	if err != nil {
		return fmt.Errorf("error reading input file %s: %w", filename, err)
	}

	// Extract offset values and calculate seconds
	var xValues []float64
	for _, record := range jsonData {
		xValues = append(xValues, float64(record.Offset)/float64(time.Second))
	}

	// Create a new chart
	graph := chart.Chart{
		TitleStyle: chart.Style{
			Padding:           chart.Box{IsSet: true, Bottom: 50, Top: 0},
			TextLineSpacing:   70,
			TextVerticalAlign: chart.TextVerticalAlignBottom,
		},
		Canvas: chart.Style{
			Padding: chart.Box{IsSet: true, Top: 60},
		},
		XAxis: chart.XAxis{
			Name: "Offset Seconds",
		},
		YAxis: chart.YAxis{
			Name: seriesName(metric),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: drawing.ColorGreen,
				},
				XValues: xValues,
				YValues: getYValues(jsonData, 200, metric),
			},
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: drawing.ColorRed,
				},
				XValues: xValues,
				YValues: getYValues(jsonData, 429, metric),
			},
		},
	}

	if metric != "rate" {
		graph.Series = graph.Series[0:1]
	}

	f, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("Error creating output file %s: %w", output, err)
	}
	defer f.Close()

	// Save the chart as a PNG file
	err = graph.Render(chart.PNG, f)
	if err != nil {
		return fmt.Errorf("Error rendering chart: %w", err)
	}

	if printInfo {
		printStats(path.Base(output), jsonData)
	}

	return nil
}

func seriesName(metric string) string {
	switch metric {
	case "count":
		return "Request count"
	case "duration":
		return "Latency (ms)"
	}
	return "Request rate (req/s)"
}

// readJSONFile reads JSON data from the specified file
func readJSONFile(filename string, padding time.Duration, metric string) ([]Record, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jsonData []Record
	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		return nil, err
	}

	last := jsonData[len(jsonData)-1]
	if last.Offset < padding {
		switch metric {
		case "rate", "duration":
			jsonData = append(jsonData, Record{Time: last.Time.Add(time.Millisecond), Offset: last.Offset + time.Millisecond, Rate: 0, Status: 429})
			left := padding - last.Offset
			jsonData = append(jsonData, Record{Time: last.Time.Add(left), Offset: last.Offset + left, Rate: 0, Status: 429})
		case "count":
			jsonData = append(jsonData, Record{Time: last.Time.Add(time.Millisecond), Offset: last.Offset + time.Millisecond, Rate: 0, Status: 200, Count: last.Count})
			left := padding - last.Offset
			jsonData = append(jsonData, Record{Time: last.Time.Add(left), Offset: last.Offset + left, Rate: 0, Status: 200, Count: last.Count})
		}
	}

	return jsonData, nil
}

// getYValues extracts Rate values from the JSON data
func getYValues(data []Record, status int, metric string) []float64 {
	var yValues []float64
	for _, record := range data {
		switch metric {
		case "rate":
			if record.Status == status {
				yValues = append(yValues, float64(record.Rate))
			} else {
				yValues = append(yValues, float64(0))
			}
		case "duration":
			yValues = append(yValues, 1000*float64(record.Duration.Seconds()))
		case "count":
			yValues = append(yValues, float64(record.Count))
		}
	}
	return yValues
}

// printStats prints the request counts and average latency
func printStats(output string, data []Record) {
	var requestLog []float64
	var total float64
	for _, record := range data {
		var duration = float64(record.Duration.Milliseconds())

		requestLog = append(requestLog, duration)
		total += duration
	}

	perc := Percentile(requestLog, 95)

	fmt.Printf("[INFO] file %s, requests %d, 95th percentile latency %.4f ms, total latency %f ms\n", output[:len(output)-4], len(data), perc, total)
}
