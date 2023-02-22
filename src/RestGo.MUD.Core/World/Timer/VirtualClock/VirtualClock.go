package VirtualClock

import (
	"fmt"
	"time"
)

type virtualClock struct {
	start time.Time
}

var vc = new(virtualClock)

func GetSessionOfDay(vHours int) (string, int) {
	switch {
	case vHours >= 0 && vHours < 3:
		return "凌晨", 0
	case vHours >= 3 && vHours < 6:
		return "拂曉", 25
	case vHours >= 6 && vHours < 9:
		return "早晨", 100
	case vHours >= 9 && vHours < 12:
		return "午前", 100
	case vHours >= 12 && vHours < 15:
		return "午後", 100
	case vHours >= 15 && vHours < 18:
		return "傍晚", 100
	case vHours >= 18 && vHours < 21:
		return "薄暮", 75
	default:
		return "深夜", 0
	}
}

func GetDateString() string {
	y, m, d, h := GetDate()
	session, _ := GetSessionOfDay(h)
	return fmt.Sprintf("德隆斯 第%d年 %s %d日 %d點 %s", y, m, d, h, session)
}

func GetDate() (vyear int, vmonth string, vday int, vhours int) {
	duration := time.Since(vc.start)

	// convert duration to virtual hours
	vhours = int(duration.Minutes() / RealMinutePervHours)

	// calculate virtual days
	vday = vhours / vHoursPerDay

	// calculate virtual months
	vImonth := vday / vDaysPerMonth

	// calculate virtual years
	vyear = vImonth / vMonthsPerYear

	// remainder virtual hours, days, and months
	vhours %= vHoursPerDay
	vday %= vDaysPerMonth
	vImonth %= vMonthsPerYear

	// return the name of the month
	return vyear, monthNames[vImonth], vday, vhours
}

func GetRealDateTime(vyear int, vmonth string, vday int, vhour int) time.Time {
	// find the virtual month index based on the month name
	var vImonth int
	for i, name := range monthNames {
		if name == vmonth {
			vImonth = i
			break
		}
	}

	// calculate the virtual month and year in minutes
	yearofMinis := vyear * vMonthsPerYear * vDaysPerMonth * vHoursPerDay * RealMinutePervHours
	monthofmins := vImonth * vDaysPerMonth * vHoursPerDay * RealMinutePervHours
	dayofmins := vday * vHoursPerDay * RealMinutePervHours
	vhourofmins := vhour * RealMinutePervHours
	durationMins := yearofMinis + monthofmins + dayofmins + vhourofmins

	// add the virtual minutes to the start time
	return vc.start.Add(time.Duration(durationMins) * time.Minute).Add(8 * time.Hour)
}

func init() {
	vc = &virtualClock{
		start: time.Date(1974, 5, 14, 0, 0, 0, 0, time.UTC).Add(8 * time.Hour),
	}
}
