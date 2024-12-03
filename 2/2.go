package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction int

type Report struct {
	Levels []int
}

func (r *Report) IsSafe() bool {
	return len(GetReportErrors(r.Levels)) < 1
}

func (r *Report) DampenLevels() {
	errors := GetReportErrors(r.Levels)

	if len(errors) > 0 {
		i := errors[0]
		r.Levels = append(r.Levels[:i], r.Levels[i+1:]...)
	}
}

func GetReportErrors(levels []int) []int {
	errors := make([]int, 0)
	// Safe if all r.Levels are in the same direction
	for i := range levels {
		if i == 0 {
			continue
		}

		if len(errors) == 1 {
			// Remove the error and check the report again
		}

		increasing := levels[0] < levels[1]
		if increasing && levels[i] < levels[i-1] {
			errors = append(errors, i)
		} else if !increasing && levels[i] > levels[i-1] {
			errors = append(errors, i)
		}

		diff := abs(levels[i] - levels[i-1])
		if diff < 1 || diff > 3 {
			errors = append(errors, i)
		}
	}
	return errors
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func NewReport(report []string) Report {
	levels := make([]int, 0)
	for _, val := range report {
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("Failed to convert level to int")
		}
		levels = append(levels, i)
	}

	return Report{
		Levels: levels,
	}
}

func readReports() []Report {
	reports := make([]Report, 0)

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Fields(line)
		r := NewReport(s)
		reports = append(reports, r)
	}

	return reports
}

func p1() {
	reports := readReports()
	safeReports := 0

	for _, report := range reports {
		if report.IsSafe() {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}

func p2() {
	reports := readReports()
	safeReports := 0

	for _, report := range reports {
		if !report.IsSafe() {
			fmt.Println(report.Levels, report.IsSafe())
			report.DampenLevels()
			fmt.Println(report.Levels, report.IsSafe())
			if report.IsSafe() {
				safeReports++
			}
		} else {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}

func main() {
	p1()
	p2()
}
