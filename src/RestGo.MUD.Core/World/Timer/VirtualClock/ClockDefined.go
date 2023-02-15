package VirtualClock

// 1 vHours = 5 Mins
// 1 vDays = 24vHours * 5 mins = 2 Hours
// 1 vMonth = 26vDays = 26 * 2 Hours = 2.16667 Days
// 1 vYear = 16 vMonths = 16 * 2.16667 = 1.1555 Months
const (
	RealMinutePervHours = 5
	vHoursPerDay        = 24
	vDaysPerMonth       = 26
	vMonthsPerYear      = 16
)

var monthNames = []string{
	"雄鷹之月",
	"玄晶之月",
	"寒冰之月",
	"九星之月",
	"芙蓉之月",
	"金輪之月",
	"修羅之月",
	"照空之月",
	"紫韻之月",
	"蒼炎之月",
	"緋雁之月",
	"暗石之月",
	"月光之月",
	"黃金之月",
	"極光之月",
	"流星之月",
}
