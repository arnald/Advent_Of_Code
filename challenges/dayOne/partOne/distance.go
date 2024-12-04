package dayone

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filePath = "challenges/dayOne/distances_list.txt"

// TotalDistance calculates the total distance between two sorted lists.
func TotalDistance() (int, error) {
	listOne, listTwo, err := ReadTxtFile(filePath)
	if err != nil {
		return 0, err
	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	return calculateTotalDistance(listOne, listTwo)
}

// readTxtFile reads and parses two lists of integers from the given file.
func ReadTxtFile(filePath string) ([]int, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var listOne, listTwo []int

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		if len(values) != 2 {
			continue
		}

		first, err1 := strconv.Atoi(values[0])
		second, err2 := strconv.Atoi(values[1])
		if err1 != nil || err2 != nil {
			continue
		}

		listOne = append(listOne, first)
		listTwo = append(listTwo, second)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return listOne, listTwo, nil
}

// calculateTotalDistance computes the total absolute distance between two sorted lists.
func calculateTotalDistance(listOne, listTwo []int) (int, error) {
	if len(listOne) != len(listTwo) {
		return 0, fmt.Errorf("lists must have the same length")
	}

	totalDistance := 0
	for i := 0; i < len(listOne); i++ {
		totalDistance += int(math.Abs(float64(listOne[i] - listTwo[i])))
	}

	return totalDistance, nil
}
