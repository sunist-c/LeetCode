package xiaochu_game

import (
	"leetcode/cases"
	"testing"
)

type xiaochuGameInput struct {
	board [][]int
	k     int
	n     int
}

func TestXiaoChuGame(t *testing.T) {
	testCases := []cases.SpecialJudgeTestCase[xiaochuGameInput, [][]int]{
		{
			Input: xiaochuGameInput{
				board: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 3, 0, 0},
					{0, 0, 5, 4, 0, 0, 0, 3, 0, 0},
					{1, 0, 5, 4, 5, 0, 2, 2, 3, 0},
					{2, 2, 1, 1, 1, 2, 2, 2, 2, 0},
					{1, 1, 1, 1, 1, 1, 1, 2, 2, 3},
				},
				k: 3,
				n: 6,
			},
			WantOutput: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 5, 4, 0, 0, 0, 0, 0, 0},
				{2, 2, 5, 4, 5, 0, 0, 0, 0, 0},
			},
		},
		{
			Input: xiaochuGameInput{
				board: [][]int{
					{1},
					{2},
				},
				k: 1,
				n: 2,
			},
			WantOutput: [][]int{
				{0},
				{0},
			},
		},
	}

	judgeFunction := func(input xiaochuGameInput, output [][]int, wantOutput [][]int) bool {
		if len(output) != len(wantOutput) {
			return false
		}

		for i, ints := range wantOutput {
			if len(output[i]) != len(ints) {
				return false
			}

			for j, v := range ints {
				if output[i][j] != v {
					return false
				}
			}
		}

		return true
	}

	originFunction := func(input xiaochuGameInput) (output [][]int) {
		return solution(input.n, input.k, input.board)
	}

	t.Run(cases.NewSpecialJudgeTestCases("XiaoChuGame", originFunction, judgeFunction, testCases...).JudgeFunction())
}
