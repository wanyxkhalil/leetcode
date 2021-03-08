package array

type Interval struct {
	Start int
	End   int
}

func merge56(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)
	res := make([]Interval, 0)
	res = append(res, intervals[0])
	curlIndex := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start > res[curlIndex].End {
			curlIndex++
			res = append(res, intervals[i])
		} else {
			res[curlIndex].End = max(intervals[i].End, res[curlIndex].End)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func quickSort(a []Interval, lo, hi int) {
	if lo >= hi {
		return
	}

	p := partitionSort(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func partitionSort(a []Interval, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if (a[j].Start < pivot.Start) || (a[j].Start == pivot.Start && a[j].End < pivot.End) {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
