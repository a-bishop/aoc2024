package day4

import (
	"io"
	"strings"

	"github.com/a-bishop/aoc2024/utils"
)

func buildGrid(lines []string) [][]string {
	g := [][]string{}
	for _, line := range lines {
		g = append(g, strings.Split(line, ""))
	}
	return g
}

func safeAccess(grid [][]string, row, col int) string {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
		// Return empty string if out of bounds
		return ""
	}
	return grid[row][col]
}

func extractRows(grid [][]string) []string {
	var rows []string
	for _, row := range grid {
		rows = append(rows, strings.Join(row, ""))
	}
	return rows
}

func extractColumns(grid [][]string) []string {
	var cols []string
	if len(grid) == 0 {
		return cols
	}
	for col := 0; col < len(grid[0]); col++ {
		var sb strings.Builder
		for row := 0; row < len(grid); row++ {
			sb.WriteString(grid[row][col])
		}
		cols = append(cols, sb.String())
	}
	return cols
}

func extractDiagonals(grid [][]string, topLeftToBottomRight bool) []string {
	var diagonals []string
	rows := len(grid)
	if rows == 0 {
		return diagonals
	}
	cols := len(grid[0])

	// Helper function to extract a diagonal
	extractDiagonal := func(startRow, startCol, rowStep, colStep int) string {
		var sb strings.Builder
		row, col := startRow, startCol
		for row >= 0 && row < rows && col >= 0 && col < cols {
			sb.WriteString(grid[row][col])
			row += rowStep
			col += colStep
		}
		return sb.String()
	}

	// Extract all diagonals
	if topLeftToBottomRight {
		for col := 0; col < cols; col++ {
			diagonals = append(diagonals, extractDiagonal(0, col, 1, 1))
		}
		for row := 1; row < rows; row++ {
			diagonals = append(diagonals, extractDiagonal(row, 0, 1, 1))
		}
	} else {
		for col := cols - 1; col >= 0; col-- {
			diagonals = append(diagonals, extractDiagonal(0, col, 1, -1))
		}
		for row := 1; row < rows; row++ {
			diagonals = append(diagonals, extractDiagonal(row, cols-1, 1, -1))
		}
	}

	return diagonals
}

func countXmas(rows []string) int {
	count := 0
	for _, row := range rows {
		count += strings.Count(row, "XMAS")
		count += strings.Count(row, "SAMX")
	}
	return count
}

func countMas(grid [][]string) int {
	count := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == "A" {
				upR := safeAccess(grid, y-1, x+1)
				upL := safeAccess(grid, y-1, x-1)
				downR := safeAccess(grid, y+1, x+1)
				downL := safeAccess(grid, y+1, x-1)
				if ((upR == "M" && downL == "S") || (upR == "S" && downL == "M")) && ((upL == "M" && downR == "S") || (upL == "S" && downR == "M")) {
					count += 1
				}

			}
		}
	}
	return count
}

func CeresSearch(input io.Reader) (part1 int, part2 int) {
	grid := buildGrid(utils.GetLines(input))

	rows := extractRows(grid)
	columns := extractColumns(grid)
	topLeftToBottomRight := extractDiagonals(grid, true)
	topRightToBottomLeft := extractDiagonals(grid, false)

	p1 := countXmas(rows) + countXmas(columns) + countXmas(topLeftToBottomRight) + countXmas(topRightToBottomLeft)
	p2 := countMas(grid)

	return p1, p2
}
