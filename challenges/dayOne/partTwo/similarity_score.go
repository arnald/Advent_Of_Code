package parttwo

import (
	p "github.com/arnald/adventofcode/challenges/dayOne/partOne"
)

const filePath = "challenges/dayOne/distances_list.txt"

func SimilarityScore() (int, error) {
	listOne, listTwo, err := p.ReadTxtFile(filePath)
	if err != nil {
		return 0, err
	}

	return calculateScore(listOne, listTwo), nil
}

func calculateScore(listOne, listTwo []int) int {
	valueCounts := make(map[int]int)
	score := 0

	for _, value := range listTwo {
		valueCounts[value]++
	}

	for _, value := range listOne {
		if count, exists := valueCounts[value]; exists {
			score += value * count
		}
	}

	return score
}
