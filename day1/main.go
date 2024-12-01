package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	FILENAME = "input.txt"
)

type LocationID int

func parseInput(logfile string) ([][]int, error) {
	result := make([][]int, 2)
	firstList := []int{}
	secondList := []int{}
	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		// log.Fatalf("open file error: %v", err)
		return result, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text() // GET the line string
		inputs := strings.Split(line, "   ")
		firstInput, err := strconv.Atoi(inputs[0])
		if err != nil {
			return result, err
		}
		firstList = append(firstList, firstInput)
		secondInput, err := strconv.Atoi(inputs[1])
		if err != nil {
			return result, err
		}
		secondList = append(secondList, secondInput)
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return result, err
	}

	result[0] = firstList
	result[1] = secondList
	return result, nil
}

func getDistanceSum(a, b []int) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("List a and b have two different sizes. Can not compute distance...")
	}
	slices.Sort(a)
	slices.Sort(b)
	distanceSum := 0
	for i := 0; i < len(a); i++ {
		distance := math.Abs(float64(a[i] - b[i]))
		distanceSum += int(distance)
	}
	return distanceSum, nil
}

func main() {
	lists, err := parseInput(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	res, err := getDistanceSum(lists[0], lists[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
