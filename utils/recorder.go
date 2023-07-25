package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"time"
)

var markdownContent = `# Leetcode

[![Build Status](https://drone.sunist.cn/api/badges/sunist-c/leetcode/status.svg)](https://drone.sunist.cn/sunist-c/leetcode) 
[![Total Problems](https://img.shields.io/badge/%v-last_commit-blue)](https://code.sunist.cn/sunist-c/leetcode)
[![Total Problems](https://img.shields.io/badge/%d+_problems-8A2BE2)](https://code.sunist.cn/sunist-c/leetcode)

这里是` + "`sunist-c`" + `的力扣刷题仓库，用于记录并提交刷题的过程。

刷题规约：

1. 每周完成至少三道题
2. 每周至少完成一道中等难度题
3. 每周完成至少一篇题解

使用Gitea的Drone完成自动化测试，刷题日志和链接放到changelog.json中

json格式如下：

` + "```json" + `
[
  {
	"date": "2023-06-30",
	"name": "2490. 回环句",
	"difficulty": "easy",
	"link": "https://leetcode.cn/problems/circular-sentence",
	"cpu_usage_rank": "100%%",
	"memory_usage_rank": "7.41%%"
  }
]
` + "```\n"

type Record struct {
	Date            string `json:"date"`
	Name            string `json:"name"`
	Difficulty      string `json:"difficulty"`
	Link            string `json:"link"`
	CpuUsageRank    string `json:"cpu_usage_rank"`
	MemoryUsageRank string `json:"memory_usage_rank"`
}

var (
	filePath         = flag.String("f", "", "file path")
	outPath          = flag.String("o", "", "out path")
	difficultyNumber = flag.Int("d", 0, "difficulty number: 0 - easy, 1 - medium, 2 - hard")
	problemShortLink = flag.String("l", "", "problem link suffix")
	problemName      = flag.String("n", "", "problem name")
	cpuUsageRank     = flag.Float64("c", 0.0, "cpu usage rank")
	memoryUsageRank  = flag.Float64("m", 0, "memory usage rank")
)

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

	updateMarkdown(len(records))
}

func updateMarkdown(length int) {
	markdownFile := path.Join(*filePath, "readme.md")
	file, openErr := os.Create(markdownFile)
	if openErr != nil {
		panic(openErr)
	}

	_, writeErr := io.WriteString(file, fmt.Sprintf(markdownContent, time.Now().Format("2006.1.2"), length))
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
