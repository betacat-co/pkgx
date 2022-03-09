package timex

import (
	"time"
)

// CheckToday 验证是否当天
func CheckToday(t time.Time) bool {
	y1, m1, d1 := time.Now().Date()
	y2, m2, d2 := t.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// FormatTimeToString time.Time类型转时间
func FormatTimeToString(time time.Time, layout string) string {
	return time.Format(layout)
}

// TimeStrToTimeDefault 默认格式日期字符串转time
func TimeStrToTimeDefault(str string) time.Time {
	parseTime, _ := time.ParseInLocation(DefaultLayout, str, time.Local)
	return parseTime
}

// TimeStrToTime 指定格式日期字符串转time
func TimeStrToTime(str, layout string) time.Time {
	parseTime, _ := time.ParseInLocation(layout, str, time.Local)
	return parseTime
}

// YearMonthToBirth 根据年月转日期字符串
func YearMonthToBirth(year, month int) (birth string) {

	t := time.Now()

	if year == 0 && month == 0 {
		birth = t.Format(YearMonthDay)
	} else if year != 0 && month == 0 {
		birth = t.AddDate(0-year, 0, 0).Format(YearMonthDay)
	} else if year == 0 && month != 0 {
		birth = t.AddDate(0, 0-month, 0).Format(YearMonthDay)
	} else {
		birth = t.AddDate(0-year, 0-month, 0).Format(YearMonthDay)
	}
	return
}

// TimeToYearMonth 给定日期字符串计算到目前经过的年月
func TimeToYearMonth(date time.Time) (year, month int) {

	// 当前时间
	now := time.Now()

	diff := (now.Year()-date.Year())*12 + (int(now.Month()) - int(date.Month()))
	year = diff / 12
	month = diff % 12

	return
}

// GetCurrentDayAndMonthTime 获取传入时间的当天、当月起始时间
func GetCurrentDayAndMonthTime(t time.Time) (currentDayStart, currentEnd, currentMonthStart, currentMonthEnd time.Time) {

	currentDayStart = t
	currentEnd = t.AddDate(0, 0, 1)

	startTime := t.AddDate(0, 0, -t.Day()+1)
	currentMonthStart = startTime
	currentMonthEnd = startTime.AddDate(0, 1, 0)

	return
}

// GetTimeFromStrDate 根据日期返回具体时间
func GetTimeFromStrDate(layoutDate, date string) (year, month, day int, err error) {
	d, err := time.Parse(layoutDate, date)
	if err != nil {
		return 0, 0, 0, err
	}
	year = d.Year()
	month = int(d.Month())
	day = d.Day()
	return
}

// GetZodiac 通过年份，返回生肖
func GetZodiac(year int) (zodiac string) {
	if year <= 0 {
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}
	if x == 0 {
		zodiac = "牛"
	}
	if x == 11 || x == -1 {
		zodiac = "虎"
	}
	if x == 10 || x == -2 {
		zodiac = "兔"
	}
	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}

func GetAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}

// GetConstellation 通过月份、日返回星座名称
func GetConstellation(month, day int) (star string) {
	if month <= 0 || month >= 13 {
		star = "-1"
	}
	if day <= 0 || day >= 32 {
		star = "-1"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "水瓶座"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "双鱼座"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "白羊座"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "金牛座"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "双子座"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "巨蟹座"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "狮子座"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "处女座"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "天秤座"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "天蝎座"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "射手座"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "魔蝎座"
	}

	return star
}
