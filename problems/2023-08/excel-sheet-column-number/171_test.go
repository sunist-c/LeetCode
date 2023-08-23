package excel_sheet_column_number

import (
	"leetcode/cases"
	"testing"
)

func TestExcelSheetColumnNumber(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[string, int]{
		{
			Input:      "A",
			WantOutput: 1,
		},
		{
			Input:      "AB",
			WantOutput: 28,
		},
		{
			Input:      "ZY",
			WantOutput: 701,
		},
		{
			Input:      "FXSHRXW",
			WantOutput: 2147483647,
		},
	}

	cases.NewCommonTestCases("ExcelSheetColumnNumber", titleToNumber, testCases...).Run(t)
}
