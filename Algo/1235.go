package algo

import "sort"

type job struct {
	start, end, profit int
}

func maxProfit(jobs []job) int {
	switch {
	case len(jobs) == 0:
		return 0
	case len(jobs) == 1:
		return jobs[0].profit
	case len(jobs) == 2:
		if jobs[0].end > jobs[1].start {
			if jobs[0].profit > jobs[1].profit {
				return jobs[0].profit
			} else {
				return jobs[1].profit
			}
		} else {
			return jobs[0].profit + jobs[1].profit
		}
	}
	sch := make([][2]int, 0)
	for k := range jobs {
		curMax := 0
		for i := 0; i < len(sch); i++ {
			if jobs[k].start >= sch[i][0] {
				if curMax < sch[i][1] {
					curMax = sch[i][1]
				}
				if k < len(jobs)-1 && jobs[k].end <= jobs[k+1].start {
					sch = append(sch[:i], sch[i+1:]...)
					i--
				}
			}
		}
		sch = append(sch, [2]int{jobs[k].end, curMax + jobs[k].profit})
	}
	allMax := 0
	for _, v := range sch {
		if allMax < v[1] {
			allMax = v[1]
		}
	}
	return allMax
}

func JobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([]job, len(profit))
	for k := range profit {
		jobs[k].start = startTime[k]
		jobs[k].end = endTime[k]
		jobs[k].profit = profit[k]
	}
	sort.Slice(jobs, func(i, j int) bool {
		switch {
		case jobs[i].start != jobs[j].start:
			return jobs[i].start < jobs[j].start
		case jobs[i].end != jobs[j].end:
			return jobs[i].end < jobs[j].end
		default:
			return jobs[i].profit < jobs[j].profit
		}
	})
	return maxProfit(jobs)
}
