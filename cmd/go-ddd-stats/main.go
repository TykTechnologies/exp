package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	if err := start(context.Background()); err != nil {
		log.Fatal(err)
	}
}

type File struct {
	Name string
	Package string
	Size int64
}

type Size struct {
	Size  string
	Count int
}

func start(_ context.Context) error {
	files, err := glob(".", ".go")
	if err != nil {
		return err
	}

	records := []File{}
	for _, filename := range files {
		packageName := path.Base(filename)

		size, _ := filesize(filename)

		records = append(records, File{
			Name: filename,
			Package: packageName,
			Size: size,
		})
	}

	asJSON := slices.Contains(os.Args, "--json")
	groupByPackage := slieces.Contains(os.Args, "--group")
	printStats := slices.Contains(os.Args, "--stats")

	grouped := map[string][]File{}
	for _, stat := range records {
		size := findGroup(stat)
		if _, ok := grouped[size]; !ok {
			grouped[size] = []File{}
		}
		grouped[size] = append(grouped[size], stat)
	}

	keys := maps.Keys(grouped)

	var lastErr error
	sort.Slice(keys, func(i, j int) bool {
		a, err := strconv.ParseInt(keys[i], 10, 64)
		if err != nil {
			lastErr = err
			return false
		}
		b, err := strconv.ParseInt(keys[j], 10, 64)
		if err != nil {
			lastErr = err
			return false
		}
		return a < b
	})
	if lastErr != nil {
		return lastErr
	}

	sizes := make([]int64, 0, len(records))
	for _, stat := range records {
		sizes = append(sizes, stat.Size)
	}

	data := stats.LoadRawData(sizes)
	var median, p80 float64
	median, _ = stats.Median(data)
	p80, _ = stats.Percentile(data, 80)

	packageCount := getPackageCount()

	switch {
	case asJSON:
		if !printStats {
			b, err := json.Marshal(records)
			if err != nil {
				return err
			}
			fmt.Println(string(b))
			return nil
		}

		b, err := json.Marshal(struct {
			Package  string
			Sizes    []Size
			Median   float64
			P80      float64
			Packages int
			Files    int
		}{
			getPackageName(),
			getSizes(keys, grouped),
			median,
			p80,
			packageCount,
			len(records),
		})
		if err != nil {
			return err
		}

		fmt.Println(string(b))
		return nil

	default:
		if !printStats {
			for _, stat := range records {
				fmt.Println(stat.Name, stat.Size)
			}
			return nil
		}

		fmt.Println("Package:", getPackageName())
		for _, key := range keys {
			fmt.Println("Size", key, "KB, Count", len(grouped[key]))
		}
		fmt.Printf("Median size: %.0f\n", median)
		fmt.Printf("Size 80th percentile: %.0f\n", p80)
		fmt.Println("Number of packages:", packageCount)
		fmt.Println("Number of files:", len(records))
		return nil
	}
	return nil
}

func getPackageName() string {
	output, _ := exec.Command("go", "list", ".").CombinedOutput()
	return strings.TrimSpace(string(output))
}

func getPackageCount() int {
	output, _ := exec.Command("go", "list", "./...").CombinedOutput()
	lines := bytes.Split(output, []byte("\n"))
	return len(lines)
}

func getSizes(keys []string, grouped map[string][]File) []Size {
	result := make([]Size, 0, len(keys))
	for _, key := range keys {
		result = append(result, Size{
			Size:  key + " KB",
			Count: len(grouped[key]),
		})
	}
	return result
}

func findGroup(file File) string {
	var increment int64 = 4

	bucket := increment
	for {
		if file.Size > bucket*1024 {
			bucket *= 2
			continue
		}
		return fmt.Sprint(bucket)
	}
}

func filesize(filename string) (int64, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func glob(dir string, ext string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(filename string, f os.FileInfo, err error) error {
		if filepath.Ext(filename) == ext {
			files = append(files, filename)
		}
		return nil
	})

	return files, err
}
