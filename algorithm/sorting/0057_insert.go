package sorting

import "github.com/gopperin/leetcode/algorithm/core"

func Insert57(intervals []Interval, newInterval Interval) []Interval {
	res := make([]Interval, 0)
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	curIndex := 0
	for curIndex < len(intervals) && intervals[curIndex].End < newInterval.Start {
		res = append(res, intervals[curIndex])
		curIndex++
	}

	for curIndex < len(intervals) && intervals[curIndex].Start <= newInterval.End {
		newInterval = Interval{Start: core.Min(newInterval.Start, intervals[curIndex].Start), End: core.Max(newInterval.End, intervals[curIndex].End)}
		curIndex++
	}
	res = append(res, newInterval)

	for curIndex < len(intervals) {
		res = append(res, intervals[curIndex])
		curIndex++
	}
	return res
}
