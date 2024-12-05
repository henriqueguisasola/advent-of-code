package main

import (
	"fmt"
	"strings"
)

func part1() {
	var lines = strings.Split(input, "\n")
	var matrix [][]string
	var counter = 0

	for _, element := range lines {
		var chars = strings.Split(element, "")
		matrix = append(matrix, chars)
	}

	var width = len(matrix[0])
	var height = len(matrix)

	for i, line := range matrix {
		for j, char := range line {
			if char == "X" {
				// Check E
				if j < width-3 && line[j+1]+line[j+2]+line[j+3] == "MAS" {
					counter += 1
				}

				// Check W
				if j > 2 && line[j-1]+line[j-2]+line[j-3] == "MAS" {
					counter += 1
				}

				// Check N
				if i > 2 && matrix[i-1][j]+matrix[i-2][j]+matrix[i-3][j] == "MAS" {
					counter += 1
				}

				// Check S
				if i < height-3 && matrix[i+1][j]+matrix[i+2][j]+matrix[i+3][j] == "MAS" {
					counter += 1
				}

				// Check NE
				if j < width-3 && i > 2 && matrix[i-1][j+1]+matrix[i-2][j+2]+matrix[i-3][j+3] == "MAS" {
					counter += 1
				}

				// Check NW
				if j > 2 && i > 2 && matrix[i-1][j-1]+matrix[i-2][j-2]+matrix[i-3][j-3] == "MAS" {
					counter += 1
				}

				// Check SE
				if j < width-3 && i < height-3 && matrix[i+1][j+1]+matrix[i+2][j+2]+matrix[i+3][j+3] == "MAS" {
					counter += 1
				}

				// Check SW
				if j > 2 && i < height-3 && matrix[i+1][j-1]+matrix[i+2][j-2]+matrix[i+3][j-3] == "MAS" {
					counter += 1
				}
			}
		}
	}

	fmt.Println(counter)
}

func part2() {
	var lines = strings.Split(input, "\n")
	var matrix [][]string
	var counter = 0

	for _, element := range lines {
		var chars = strings.Split(element, "")
		matrix = append(matrix, chars)
	}

	var width = len(matrix[0])
	var height = len(matrix)

	for i, line := range matrix {
		for j, char := range line {
			if char == "M" {
				// Horizontal MAS
				if i < height-2 && matrix[i+2][j] == "M" {
					if j >= 2 && matrix[i][j-2]+matrix[i+1][j-1]+matrix[i+2][j-2] == "SAS" {
						counter += 1
					}

					if j < width-2 && matrix[i][j+2]+matrix[i+1][j+1]+matrix[i+2][j+2] == "SAS" {
						counter += 1
					}
				}

				// Vertical MAS
				if j < width-2 && matrix[i][j+2] == "M" {
					if i >= 2 && matrix[i-2][j]+matrix[i-1][j+1]+matrix[i-2][j+2] == "SAS" {
						counter += 1
					}

					if i < height-2 && matrix[i+2][j]+matrix[i+1][j+1]+matrix[i+2][j+2] == "SAS" {
						counter += 1
					}
				}
			}
		}
	}

	fmt.Println(counter)
}

func main() {
	part1()
	part2()
}
