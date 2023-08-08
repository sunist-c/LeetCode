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
