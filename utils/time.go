package utils

import "time"

const (
	TimeLayout string = "2006-01-02 15:04:05"
	DateLayout string = "2006-01-02"
)

//获取时间：国内明天上午3点钟
func GetChinaTomorrow3AM() time.Time {
	//注意不推荐下面做法
	//loc, _ := time.LoadLocation("Asia/Shanghai") // CST /todo:这种方法有的机器取不到时区，报错： missing Location in call to Date
	//t, _ := time.ParseInLocation("2006-01-02", now.AddDate(0, 0, 1).Format("2006-01-02"), loc)

	dateString := time.Now().Format("20060102")
	//获取的时间时上午8：00
	t, _ := time.Parse("20060102", dateString)
	h, _ := time.ParseDuration("1h")
	//明天凌晨3点
	t = t.Add(19 * h)
	return t
}

//获取时间：国内明天上午3点钟
func GetChinaTomorrow3AMSeconds() uint64 {
	tomorrow3AM := GetChinaTomorrow3AM()
	secondsF := tomorrow3AM.Sub(time.Now()).Seconds()
	return uint64(secondsF)
}

func GetChinaTomorrowAMSeconds() uint64 {
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02", now.AddDate(0, 0, 1).Format("2006-01-02"), loc)
	secondsF := t.Sub(time.Now()).Seconds()
	return uint64(secondsF)
}

func GetLocalTomorrowAMSeconds() int64 {
	now := time.Now()
	t, _ := time.ParseInLocation("2006-01-02", now.AddDate(0, 0, 1).Format("2006-01-02"), time.Local)
	secondsF := t.Sub(time.Now()).Seconds()
	return int64(secondsF)
}