package main

import "flag"

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

var markdownContent = `# Leetcode

[![Current Commit](https://img.shields.io/badge/%v-last_commit-blue)](https://studio.sunist.work/sunist-c/leetcode)
[![Total Problems](https://img.shields.io/badge/%d+_problems-8A2BE2)](https://studio.sunist.work/sunist-c/leetcode)

[![Average Difficulty](https://img.shields.io/badge/difficulty-%.4f-%s)](https://studio.sunist.work/sunist-c/leetcode)
[![Average Cpu Usage](https://img.shields.io/badge/cpu_usage_rank-%.2f%%25-%s)](https://studio.sunist.work/sunist-c/leetcode)
[![Average Memory Usage](https://img.shields.io/badge/memory_usage_rank-%.2f%%25-%s)](https://studio.sunist.work/sunist-c/leetcode)

这里是` +
	"`sunist-c`" + `的力扣刷题仓库，用于记录并提交刷题的过程。

刷题规约：

1. 每周完成至少三道题
2. 每周至少完成一道中等难度题
3. 每周完成至少一篇题解

使用Gitea-Actions完成自动化测试，刷题日志和链接放到changelog.json中

json格式如下：

` +
	"```json" +
	`
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
` +
	"```\n"
