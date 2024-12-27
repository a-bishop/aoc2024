package day5

import (
	"io"
	"strings"

	"github.com/a-bishop/aoc2024/utils"
)

type Exists struct{}

type StringSet map[string]Exists

func getOrderingMap(ordering []string) map[string]StringSet {
	orderingMap := make(map[string]StringSet)
	for _, order := range ordering {
		pageOrders := strings.Split(order, "|")
		if _, ok := orderingMap[pageOrders[0]]; !ok {
			orderingMap[pageOrders[1]] = make(StringSet)
		}
		// for each given later page, add to the set of pages that should come before it
		orderingMap[pageOrders[1]][pageOrders[0]] = Exists{}
	}
	return orderingMap
}

func PrintQueue(input io.Reader) (part1 int, part2 int) {
	ordering, pageUpdates := utils.GetLinesWithSplit(input, "")

	orderingMap := getOrderingMap(ordering)

	count := 0
	for _, pageUpdate := range pageUpdates {
		pages := strings.Split(pageUpdate, ",")
		valid := true
		for idx, page := range pages {
			// there is an ordering rule to consider
			if _, ok := orderingMap[page]; ok {
				for _, laterPage := range pages[idx+1:] {
					// if the second page is meant to come before the first page, it's invalid
					if _, ok := orderingMap[page][laterPage]; ok {
						valid = false
						break
					}
				}
			}
		}
		if valid {
			count += utils.MustAtoi(pages[len(pages)/2])
		}
	}

	return count, 0
}
