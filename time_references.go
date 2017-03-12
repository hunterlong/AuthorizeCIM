package AuthorizeCIM

import "time"

func Now() time.Time {
	current_time := time.Now().UTC()
	return current_time
}

func LastWeek() time.Time {
	t := time.Now().UTC().AddDate(0, 0, -1)
	return t
}

func CurrentDate() string {
	current_time := time.Now().UTC()
	return current_time.Format("2006-01-02")
}

func IntervalMonthly() Interval {
	interval := Interval{Length: "1", Unit: "months"}
	return interval
}

func IntervalQuarterly() Interval {
	interval := Interval{Length: "3", Unit: "months"}
	return interval
}

func IntervalWeekly() Interval {
	interval := Interval{Length: "7", Unit: "days"}
	return interval
}

func IntervalYearly() Interval {
	interval := Interval{Length: "365", Unit: "days"}
	return interval
}
