package timex

import (
	"fmt"
	"math"
	"time"
)

var (
	TimeLayout string = "2006-01-02 15:04:05"
	DateLayout        = "20060102"
	TimeFormat        = map[string]string{
		"Y-m-d H:i:s": "2006-01-02 15:04:05",
		"Y-m-d":       "2006-01-02",
		"Ymd":         "20060102",
		"H:i:s":       "15:04:05",
		"Y":           "2006",
		"m":           "01",
		"d":           "02",
	}
)

// 获取当前时间GetTimeNow()，用于测试时的时间修改
func GetTimeNow() time.Time {
	//redisConf := redis.RedisKeyConf{
	//	RedisConf: redis.RedisConf{
	//		Host: "122.228.113.235:17006",
	//		Type: "",
	//		Pass: "xiaozi527sport",
	//		Tls:  false,
	//	},
	//}
	//redisClient := redisConf.NewRedis()
	//val, err := redisClient.Get("test:now:gap:seconds")
	//if err != nil {
	//	fmt.Errorf("GetTimeNow test:now:gap error:%s", err.Error())
	//	return time.Now()
	//}
	//gap, _ := strconv.ParseInt(val, 10, 64)
	//return time.Now().Add(time.Second * time.Duration(gap))

	return time.Now()
}

func GetChinaTomorrowAMSeconds() int64 {
	now := GetTimeNow()
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02", now.AddDate(0, 0, 1).Format("2006-01-02"), loc)
	secondsF := t.Sub(GetTimeNow()).Seconds()
	return int64(secondsF)
}

func GetLocalTomorrowAMSeconds() int64 {
	now := GetTimeNow()
	t, _ := time.ParseInLocation("2006-01-02", now.AddDate(0, 0, 1).Format("2006-01-02"), time.Local)
	secondsF := t.Sub(GetTimeNow()).Seconds()
	return int64(secondsF)
}

func GetTodayZero() time.Time {
	t := GetTimeNow()
	zero := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return zero
}

func GetZero(targetTime time.Time) time.Time {
	zero := time.Date(targetTime.Year(), targetTime.Month(), targetTime.Day(), 0, 0, 0, 0, targetTime.Location())
	return zero
}

// 解析时间,"2021-03-17 00:00:00"
func ParseTime(timeStr string) (datetime time.Time) {
	datetime, _ = time.ParseInLocation(TimeLayout, timeStr, time.Local)
	return
}

//eg:20210812170000
func ParseTimeString(timeStr string) (datetime time.Time) {
	datetime, _ = time.ParseInLocation("20060102150405", timeStr, time.Local)
	return
}

func GetDateInterval(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	interval := int(math.Abs(t1.Sub(t2).Hours())/24) + 1
	return interval
}

// 获取过去的天数,dateString格式20060102
func SinceDays(dateString string) (int64, error) {
	registerTime, err := time.Parse("20060102", dateString)
	if err != nil {
		return 0, err
	} else {
		days := int64(math.Ceil(time.Since(registerTime).Hours() / 24))
		if days == 0 {
			days = 1
		}
		return days, nil
	}
}

func IsSameDay(t1, t2 time.Time) bool {
	year1, month1, day1 := t1.Date()
	year2, month2, day2 := t2.Date()
	return day1 == day2 && month1 == month2 && year1 == year2
}

func LastHourStartAndEnd() (int, int64, int64) {
	now := time.Now()
	startOfLastHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()-1, 0, 0, 0, now.Location())
	endOfLastHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location()).Add(-time.Nanosecond)
	startOfLastHourUnix := startOfLastHour.Unix()
	endOfLastHourUnix := endOfLastHour.Unix()
	return startOfLastHour.Hour(), startOfLastHourUnix, endOfLastHourUnix
}

func CurrentHourStartAndEnd() (int, int64, int64) {
	now := time.Now()
	startOfHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	endOfHour := startOfHour.Add(time.Hour - time.Second)

	startOfHourTimestamp := startOfHour.Unix()
	endOfHourTimestamp := endOfHour.Unix()
	return startOfHour.Hour(), startOfHourTimestamp, endOfHourTimestamp
}

func GetNowDateForLocation() string {
	tm := time.Now()
	// 加载时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return ""
	}
	// 将时间转换为指定时区
	tmInLocation := tm.In(location)
	// 将时间格式化为日期字符串
	dateStr := tmInLocation.Format(time.DateOnly)
	return dateStr
}

func ToDate(ts int64) string {
	tm := time.Unix(ts, 0)

	// 加载时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return ""
	}

	// 将时间转换为指定时区
	tmInLocation := tm.In(location)

	// 将时间格式化为日期字符串
	dateStr := tmInLocation.Format(time.DateOnly)
	return dateStr
}

func ToDateTime(ts int64) string {
	tm := time.Unix(ts, 0)

	// 加载时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return ""
	}

	// 将时间转换为指定时区
	tmInLocation := tm.In(location)

	// 将时间格式化为日期字符串
	dateStr := tmInLocation.Format(time.DateTime)
	return dateStr
}

// WeekDays 获取本周所有时间
func WeekDays(t time.Time) []time.Time {
	week := []time.Time{}
	// 计算出这一天距离周日的天数
	offset := int(t.Weekday())
	for i := 0; i < 7; i++ {
		week = append(week, t.AddDate(0, 0, i-offset))
	}
	return week
}

// DaysInMonth 获取某月的天数
func DaysInMonth(t time.Time) int {
	// 获得下个月的同一天
	nextMonth := t.AddDate(0, 1, 0)
	// 然后将日期减少到本月的第一天，也就是下个月的第一天的前一天
	lastDayOfMonth := nextMonth.AddDate(0, 0, -nextMonth.Day())
	// 返回该日的日期
	return lastDayOfMonth.Day()
}
