package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	Date            string `json:"date"`
	Name            string `json:"name"`
	Difficulty      string `json:"difficulty"`
	Link            string `json:"link"`
	CpuUsageRank    string `json:"cpu_usage_rank"`
	MemoryUsageRank string `json:"memory_usage_rank"`
}

func main() {
	flag.Parse()
	if *filePath != "" && *outPath != "" {
		inputFile, openErr := os.OpenFile(*filePath, os.O_RDWR, 0666)
		if openErr != nil {
			panic(openErr)
		}
		outputFile, outFileErr := os.OpenFile(*outPath, os.O_RDWR|os.O_CREATE, 0666)
		if outFileErr != nil {
			panic(outFileErr)
		}
		convertFile(inputFile, outputFile)
		return
	}

	if *filePath != "" && *outPath == "" {
		appendFile()
		return
	}
}

func appendFile() {
	jsonPath := path.Join(*filePath, "changelog.json")
	file, openErr := os.OpenFile(jsonPath, os.O_RDWR, 0666)
	if openErr != nil {
		panic(openErr)
	}

	bytes, readErr := io.ReadAll(file)
	if readErr != nil {
		panic(readErr)
	}

	var records []Record
	unmarshalErr := json.Unmarshal(bytes, &records)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	closeErr := file.Close()
	if closeErr != nil {
		panic(closeErr)
	}
	file, openErr = os.Create(jsonPath)
	if openErr != nil {
		panic(openErr)
	}

	key := time.Now().Format("2006-01-02")
	var difficultyString string
	switch *difficultyNumber {
	case 0:
		difficultyString = "easy"
	case 1:
		difficultyString = "medium"
	case 2:
		difficultyString = "hard"
	}

	dto := Record{
		Date:            key,
		Name:            *problemName,
		Difficulty:      difficultyString,
		Link:            fmt.Sprintf("https://leetcode.cn/problems/%s", *problemShortLink),
		CpuUsageRank:    fmt.Sprintf("%.2f%%", *cpuUsageRank*100.0),
		MemoryUsageRank: fmt.Sprintf("%.2f%%", *memoryUsageRank*100.0),
	}

	records = append(records, dto)
	sort.SliceStable(records, func(i, j int) bool {
		iDate, _ := time.Parse("2006-01-02", records[i].Date)
		jDate, _ := time.Parse("2006-01-02", records[j].Date)
		return iDate.After(jDate)
	})

	outBytes, marshalErr := json.MarshalIndent(&records, "", "  ")
	if marshalErr != nil {
		panic(marshalErr)
	}

	_, writeErr := io.WriteString(file, string(outBytes))
	if writeErr != nil {
		panic(writeErr)
	}
	closeErr = file.Close()
	if closeErr != nil {
		panic(closeErr)
	}

	updateMarkdown(statistic(records))
}

func statistic(records []Record) (length int, difficulty float64, cpuUsage float64, memoryUsage float64) {
	total := len(records)
	difficultySum := 0.0
	cpuUsageSum := 0.0
	memoryUsageSum := 0.0
	for _, record := range records {
		cpuUsage, _ := strconv.ParseFloat(strings.TrimSuffix(record.CpuUsageRank, "%"), 64)
		memoryUsage, _ := strconv.ParseFloat(strings.TrimSuffix(record.MemoryUsageRank, "%"), 64)
		switch record.Difficulty {
		case "easy":
			difficultySum += 1.0
		case "medium":
			difficultySum += 4.0
			cpuUsage *= 2.0
			memoryUsage *= 2.0
			total += 1
		case "hard":
			difficultySum += 9.0
			cpuUsage *= 3.0
			memoryUsage *= 3.0
			total += 2
		}
		cpuUsageSum += cpuUsage
		memoryUsageSum += memoryUsage
	}

	return len(records), difficultySum / float64(total), cpuUsageSum / float64(total), memoryUsageSum / float64(total)
}

func updateMarkdown(length int, difficulty float64, cpuUsage float64, memoryUsage float64) {
	markdownFile := path.Join(*filePath, "readme.md")
	file, openErr := os.Create(markdownFile)
	if openErr != nil {
		panic(openErr)
	}

	var difficultyColor, cpuUsageColor, memoryUsageColor string
	if difficulty < 1.1 {
		difficultyColor = levelLow
	} else if difficulty < 2.1 {
		difficultyColor = levelMedium
	} else {
		difficultyColor = levelHigh
	}

	if cpuUsage < 33 {
		cpuUsageColor = levelLow
	} else if cpuUsage < 66 {
		cpuUsageColor = levelMedium
	} else {
		cpuUsageColor = levelHigh
	}

	if memoryUsage < 33 {
		memoryUsageColor = levelLow
	} else if memoryUsage < 66 {
		memoryUsageColor = levelMedium
	} else {
		memoryUsageColor = levelHigh
	}

	_, writeErr := io.WriteString(file, fmt.Sprintf(markdownContent,
		time.Now().Format("2006.1.2"), length, difficulty, difficultyColor, cpuUsage, cpuUsageColor, memoryUsage, memoryUsageColor))
	if writeErr != nil {
		panic(writeErr)
	}

	closeErr := file.Close()
	if closeErr != nil {
		panic(closeErr)
	}
}

func convertFile(inputFile, outputFile *os.File) {
	bytes, readErr := io.ReadAll(inputFile)
	if readErr != nil {
		panic(readErr)
	}

	var records []map[string]Record
	unmarshalErr := json.Unmarshal(bytes, &records)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	var structure []Record
	for _, record := range records {
		for key, value := range record {
			value.Date = key
			structure = append(structure, value)
		}
	}

	outBytes, marshalErr := json.MarshalIndent(&structure, "", "  ")
	if marshalErr != nil {
		panic(marshalErr)
	}

	_, writeErr := io.WriteString(outputFile, string(outBytes))
	if writeErr != nil {
		panic(writeErr)
	}

	closeInputFileErr, closeOutputFileErr := inputFile.Close(), outputFile.Close()
	if closeInputFileErr != nil || closeOutputFileErr != nil {
		panic("close file error")
	}
}
