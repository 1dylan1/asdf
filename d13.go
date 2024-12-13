package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type matrix struct {
	grid [2][3]float64
}

func main() {

	var matrices []matrix
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentMatrix matrix
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Button A:") {
			currentMatrix.grid[0][0], currentMatrix.grid[1][0] = parseButton(line)
		} else if strings.HasPrefix(line, "Button B:") {
			currentMatrix.grid[0][1], currentMatrix.grid[1][1] = parseButton(line)
		} else if strings.HasPrefix(line, "Prize:") {
			currentMatrix.grid[0][2], currentMatrix.grid[1][2] = parsePrize(line)
			matrices = append(matrices, currentMatrix)
			lineCount = 0
		}
		lineCount++
	}
	x, y := PartTwo(matrices)
	fmt.Printf("Total: %.1f", (x*3)+(y))
}

func parseButton(line string) (float64, float64) {
	parts := strings.Split(line, ",")
	xPart := strings.TrimSpace(strings.Split(parts[0], "+")[1])
	yPart := strings.TrimSpace(strings.Split(parts[1], "+")[1])
	x, _ := strconv.ParseFloat(xPart, 64)
	y, _ := strconv.ParseFloat(yPart, 64)
	return x, y
}

func parsePrize(line string) (float64, float64) {
	prizePrefix := float64(10000000000000)
	parts := strings.Split(line, ",")
	xPart := strings.TrimSpace(strings.Split(parts[0], "=")[1])
	yPart := strings.TrimSpace(strings.Split(parts[1], "=")[1])
	x, _ := strconv.ParseFloat(xPart, 64)
	y, _ := strconv.ParseFloat(yPart, 64)
	return x + prizePrefix, y + prizePrefix
}

func guassianEliminationShort(matrix [2][3]float64) (float64, float64) {
	// R1 / divisor -> R1
	divisor := matrix[0][0]
	for i := range matrix[0] {
		matrix[0][i] /= divisor
	}
	// R2 - (scalar)R1 -> R2
	scalar := matrix[1][0]
	for i := range matrix[1] {
		matrix[1][i] -= (scalar * matrix[0][i])
	}

	// R2 / (scalar) -> R2
	scalar = matrix[1][1]
	for i := range matrix[1] {
		matrix[1][i] /= scalar
	}

	// R1 - (scalar)R2 -> R1
	scalar = matrix[0][1]
	for i := range matrix[0] {
		matrix[0][i] -= (scalar * matrix[1][i])
	}
	a := roundToHundredth(matrix[0][2])
	b := roundToHundredth(matrix[1][2])
	return a, b
}

func roundToHundredth(val float64) float64 {
	return math.Round(val*100) / 100
}

const epsilon = 1e-2

func PartTwo(grids []matrix) (float64, float64) {
	tokenA, tokenB := 0.0, 0.0
	for _, grid := range grids {
		a, b := guassianEliminationShort(grid.grid)
		if math.Abs(a-math.Round(a)) < epsilon && math.Abs(b-math.Round(b)) < epsilon {
			tokenA += a
			tokenB += b
		}
	}
	return tokenA, tokenB
}

