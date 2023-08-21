package main

import (
	"bufio"
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

const (
	levelHigh   = "lightgreen"
	levelMedium = "lightyellow"
	levelLow    = "lightgray"
)

var (
	filePath         = flag.String("f", "", "file path")
	outPath          = flag.String("o", "", "out path")
	difficultyNumber = flag.Int("d", 0, "difficulty number: 0 - easy, 1 - medium, 2 - hard")
	problemShortLink = flag.String("l", "", "problem link suffix")
	problemName      = flag.String("n", "", "problem name")
	cpuUsageRank     = flag.Float64("c", 0.0, "cpu usage rank")
	memoryUsageRank  = flag.Float64("m", 0, "memory usage rank")
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
	file, openErr := os.Open(markdownFile)
	if openErr != nil {
		panic(openErr)
	}

	scanner := bufio.NewScanner(file)
	var lines = []string{""}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	closeErr := file.Close()
	if closeErr != nil {
		panic(closeErr)
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

	commitLineNumber, totalProblemsLineNumber, avgDifficultyLineNumber, avgCpuUsageLineNumber, avgMemoryUsageLineNumber := 5, 6, 8, 9, 10

	lines[commitLineNumber] = fmt.Sprintf("[![Current Commit](https://img.shields.io/badge/%s-last_commit-blue)](https://studio.sunist.work/sunist-c/leetcode)", time.Now().Format("2006.1.2"))
	lines[totalProblemsLineNumber] = fmt.Sprintf("[![Total Problems](https://img.shields.io/badge/%d+_problems-8A2BE2)](https://studio.sunist.work/sunist-c/leetcode)", length)
	lines[avgDifficultyLineNumber] = fmt.Sprintf("[![Average Difficulty](https://img.shields.io/badge/difficulty-%.4f-%s)](https://studio.sunist.work/sunist-c/leetcode)", difficulty, difficultyColor)
	lines[avgCpuUsageLineNumber] = fmt.Sprintf("[![Average Cpu Usage](https://img.shields.io/badge/cpu_usage_rank-%.2f%%25-%s)](https://studio.sunist.work/sunist-c/leetcode)", cpuUsage, cpuUsageColor)
	lines[avgMemoryUsageLineNumber] = fmt.Sprintf("[![Average Memory Usage](https://img.shields.io/badge/memory_usage_rank-%.2f%%25-%s)](https://studio.sunist.work/sunist-c/leetcode)", memoryUsage, memoryUsageColor)

	writeFile, openWriteFileErr := os.Create(markdownFile)
	if openWriteFileErr != nil {
		panic(openWriteFileErr)
	}
	_, writeErr := io.WriteString(writeFile, strings.Join(lines[1:], "\n"))
	if writeErr != nil {
		panic(writeErr)
	}

	closeWriteFileErr := writeFile.Close()
	if closeWriteFileErr != nil {
		panic(closeWriteFileErr)
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
