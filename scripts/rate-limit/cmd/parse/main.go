package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

type Flags struct {
	Filename string
	Kind     string
}

func NewFlags() (*Flags, error) {
	flags := &Flags{}

	flag.StringVar(&flags.Filename, "f", "", "input filename")
	flag.StringVar(&flags.Kind, "t", "", "input kind (csv, json)")
	flag.Parse()

	if flags.Filename == "" {
		return flags, errors.New("No filename defined (-f <file>)")
	}
	if flags.Kind == "" {
		return flags, errors.New("No file type defined (-t <csv|json>)")
	}

	return flags, nil
}

func start() error {
	flags, err := NewFlags()
	if err != nil {
		return err
	}

	switch flags.Kind {
	case "csv":
		return processCSV(flags.Filename)
	case "json":
		return processJSON(flags.Filename)
	default:
		return fmt.Errorf("Unknown file type provided: %q, only json/csv supported", flags.Kind)
	}

	return nil
}

func processCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var records []*Record

	epoch := time.Now().UTC()

	for i, line := range lines {
		if i == 0 {
			// Skip header line
			continue
		}

		timeVal, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			return fmt.Errorf("error parsing time: %v, line %d", err, i)
		}

		offsetVal, err := strconv.ParseFloat(line[7], 64)
		if err != nil {
			return fmt.Errorf("error parsing offset: %v, line %d", err, i)
		}

		statusVal, err := strconv.Atoi(line[6])
		if err != nil {
			return fmt.Errorf("error parsing status code: %v, line %d", err, i)
		}

		record := &Record{
			Offset:   time.Duration(offsetVal * float64(time.Second)),
			Duration: time.Duration(timeVal * float64(time.Second)),
			Status:   statusVal,
		}
		record.Time = epoch.Add(record.Offset)

		records = append(records, record)
	}

	fillRate(records)
	fillCount(records)

	return render(records)
}

type Record struct {
	Time     time.Time
	Rate     int
	Count    int
	Offset   time.Duration
	Duration time.Duration
	Status   int
}

func (r *Record) String() string {
	return fmt.Sprintf("time=%s offset=%.2fs duration=%dns status=%d", r.Time, r.Offset.Seconds(), r.Duration.Nanoseconds(), r.Status)
}

func fillCount(series []*Record) {
	for idx, r := range series {
		r.Count = idx + 1
	}
}

func fillRate(series []*Record) {
	windowStart := func(recs []*Record, from time.Duration) int {
		for idx, r := range recs {
			if r.Offset > from {
				return idx
			}
		}
		return 0
	}

	for idx, r := range series {
		start := windowStart(series, r.Offset-time.Second)
		r.Rate = 1 + (idx - start)
	}
}

func fillOffset(series []*Record) {
	if len(series) == 0 {
		return
	}

	first := series[0].Time

	for _, record := range series {
		record.Offset = record.Time.Sub(first)
	}
}

func processJSON(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	records := []*Record{}

	for {
		record := &Record{}
		err := decoder.Decode(record)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		records = append(records, record)
	}

	fillOffset(records)
	fillRate(records)
	fillCount(records)

	return render(records)
}

func render(records []*Record) error {
	b, err := json.Marshal(records)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
