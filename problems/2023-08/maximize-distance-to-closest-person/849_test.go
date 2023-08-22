package maximize_distance_to_closest_person

import (
	"leetcode/cases"
	"testing"
)

func TestMaximizeDistanceToClosestPerson(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]int, int]{
		{
			Input:      []int{1, 0, 0, 0, 1, 0, 1},
			WantOutput: 2,
		},
		{
			Input:      []int{1, 0, 0, 0},
			WantOutput: 3,
		},
		{
			Input:      []int{0, 1},
			WantOutput: 1,
		},
	}

	cases.NewCommonTestCases("MaximizeDistanceToClosestPerson", maxDistToClosest, testCases...).Run(t)
}
