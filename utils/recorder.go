package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"
)

type Record struct {
	Name            string `json:"name"`
	Difficulty      string `json:"difficulty"`
	Link            string `json:"link"`
	CpuUsageRank    string `json:"cpu_usage_rank"`
	MemoryUsageRank string `json:"memory_usage_rank"`
}

func main() {
	difficultyNumber := flag.Int("d", 0, "difficulty number: 0 - easy, 1 - medium, 2 - hard")
	problemShortLink := flag.String("l", "", "problem link suffix")
	problemName := flag.String("n", "", "problem name")
	cpuUsageRank := flag.Float64("c", 0.0, "cpu usage rank")
	memoryUsageRank := flag.Float64("m", 0, "memory usage rank")
	flag.Parse()

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

	record := Record{
		Name:            *problemName,
		Difficulty:      difficultyString,
		Link:            fmt.Sprintf("https://leetcode.cn/problems/%s", *problemShortLink),
		CpuUsageRank:    fmt.Sprintf("%.2g%%", *cpuUsageRank*100.0),
		MemoryUsageRank: fmt.Sprintf("%.2g%%", *memoryUsageRank*100.0),
	}

	output := map[string]Record{
		key: record,
	}

	bytes, _ := json.MarshalIndent(&output, "", "  ")
	fmt.Println(string(bytes))
}
