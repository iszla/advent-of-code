package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	d "github.com/iszla/aoc25/days"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var bench = flag.String("bench", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	switch *bench {
	case "01":
		benchmark(&d.Day01{}, readInput("./days/Day01.txt"))
	case "02":
		benchmark(&d.Day02{}, readInput("./days/Day02.txt"))
	case "03":
		benchmark(&d.Day03{}, readInput("./days/Day03.txt"))
	case "04":
		benchmark(&d.Day04{}, readInput("./days/Day04.txt"))
	case "05":
		benchmark(&d.Day05{}, readInput("./days/Day05.txt"))
	case "06":
		benchmark(&d.Day06{}, readInput("./days/Day06.txt"))
	case "07":
		benchmark(&d.Day07{}, readInput("./days/Day07.txt"))

	default:
		tt := time.Now()
		printDay(&d.Day01{}, readInput("./days/Day01.txt"))
		printDay(&d.Day02{}, readInput("./days/Day02.txt"))
		printDay(&d.Day03{}, readInput("./days/Day03.txt"))
		printDay(&d.Day04{}, readInput("./days/Day04.txt"))
		printDay(&d.Day05{}, readInput("./days/Day05.txt"))
		printDay(&d.Day06{}, readInput("./days/Day06.txt"))
		printDay(&d.Day07{}, readInput("./days/Day07.txt"))
		fmt.Printf("Total execution time: %s\n", time.Since(tt))
	}
}

func printDay(day d.Day, input []string) {
	sp01 := time.Now()
	p01 := day.P01(input)
	ep01 := time.Since(sp01)
	sp02 := time.Now()
	p02 := day.P02(input)
	ep02 := time.Since(sp02)
	fmt.Printf("%T\n", day)
	fmt.Printf("Answer: %d Execution time: %s\n", p01, ep01)
	fmt.Printf("Answer: %d Execution time: %s\n", p02, ep02)
	fmt.Println("|---|---|---|---|")
}

func readInput(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("failed to read file: %s", err))
	}

	return strings.Split(string(data), "\n")
}

func benchmark(day d.Day, input []string) {
	iter := 100
	fmt.Printf("Benchmarking %T\nRunning %d iterations\n", day, iter)
	tt := time.Now()
	for range 100 {
		day.P01(input)
		day.P02(input)
	}
	fmt.Printf("Total execution time: %s\n", time.Since(tt))
}
