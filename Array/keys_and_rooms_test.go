package Array

import (
	"testing"
	"fmt"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		rooms [][]int
	}{
		{
			name: "test1",
			rooms: [][]int {
				{1},{2},{3},{},
			},
		},
		{
			name: "test2",
			rooms: [][]int{
				{6, 7, 8}, {5, 4, 9}, {}, {8}, {4}, {}, {1, 9, 2, 3}, {7}, {6, 5}, {2, 3, 1},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(canVisitAllRooms(test.rooms))
		})
	}
}
