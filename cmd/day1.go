package cmd

import (
	"fmt"
	"log"

	one "github.com/arnald/adventofcode/challenges/dayOne/partOne"
	two "github.com/arnald/adventofcode/challenges/dayOne/partTwo"
	"github.com/spf13/cobra"
)

var part string

var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Solve challenges for Day 1",
	Long: `Solve challenges for Day 1 of Advent of Code.

You can specify which part of the challenge to run using the --part flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch part {
		case "1":
			totalDistance, err := one.TotalDistance()
			if err != nil {
				log.Fatalf("Error in Day 1 Part 1: %v\n", err)
			}
			fmt.Printf("Day 1 Part 1 - Total Distance: %d\n", totalDistance)

		case "2":
			similarityScore, err := two.SimilarityScore()
			if err != nil {
				log.Fatalf("Error in Day 1 Part 2: %v\n", err)
			}
			fmt.Printf("Day 1 Part 2 - Similarity Score: %d\n", similarityScore)

		default:
			fmt.Println("Invalid part. Use --part=1 or --part=2.")
		}
	},
}

func init() {
	// Add the day1 command to the root command
	rootCmd.AddCommand(day1Cmd)
	day1Cmd.Flags().StringVarP(&part, "part", "p", "", "Specify the part of the challenge to run (1 or 2)")
}
