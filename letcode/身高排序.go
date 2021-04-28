/**
* @Author: Gosin
* @Date: 2021/4/27 22:10
 */

package main

import (
	"sort"
)

func main() {
	arr3 := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}

	reconstructQueue(arr3)
}

func reconstructQueue(people [][]int) [][]int {

	sort.Slice(people, func(i, j int) bool {
		return (people[i][0] > people[j][0]) || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1])
		people[p[1]] = p
	}

	return people

}
