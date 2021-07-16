package a_star_algorithm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	startRow := 0
	startCol := 1
	endRow := 4
	endCol := 3
	graph := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1},
		{0, 0, 0, 0, 0},
	}
	expected := [][]int{{0, 1}, {0, 0}, {1, 0}, {2, 0}, {2, 1}, {3, 1}, {4, 1}, {4, 2}, {4, 3}}
	actual := AStarAlgorithm(startRow, startCol, endRow, endCol, graph)
	require.Equal(t, expected, actual)
}
