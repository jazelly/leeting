package minSpeedOnTime

import "math"

func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	sd := 0
	for i := 0; i < n; i++ {
		sd += dist[i]
	}

	if hour <= float64(n-1) {
		return -1
	}

	sl := 1
	sr := 10000001
	var res int
	for sl < sr {
		sm := (sl + sr) / 2
		if timeRequired(dist, sm) <= hour {
			res = sm
			sr = sm
		} else {
			sl = sm + 1
		}
	}
	return res

}

func timeRequired(dist []int, speed int) float64 {
	var time float64
	for i := 0; i < len(dist); i++ {
		t := float64(dist[i]) / float64(speed)
		if i < len(dist)-1 {
			t = math.Ceil(t)
		}
		time += t
	}
	return time
}
