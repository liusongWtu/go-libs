package cron

import "github.com/robfig/cron/v3"

const (
	//注意DelayIfStillRunning与SkipIfStillRunning是有本质上的区别的，前者DelayIfStillRunning只要时间足够长，所有的任务都会按部就班地完成，只是可能前一个任务耗时过长，导致后一个任务的执行时间推迟了一点。SkipIfStillRunning会跳过一些执行。
	SkipIfStillRunning  Mode = 1
	DelayIfStillRunning Mode = 2
	RunIfStillRunning   Mode = 3
)

var (
	_cron *cron.Cron
)

type Mode int

func init() {
	c := cron.New(cron.WithSeconds(), cron.WithChain())
	c.Start()

	_cron = c

}

type job struct {
	cmd func()
}

func (d *job) Run() {
	d.cmd()
}

// AddJob 添加定时任务
// spec 指定触发时间规则，支持两种格式：
// 第一种：@every 1s表示每秒触发一次，@every后加一个时间间隔，表示每隔多长时间触发一次。例如@every 1h表示每小时触发一次，@every 1m2s表示每隔 1 分 2 秒触发一次。time.ParseDuration()支持的格式都可以用在这里。
// 第二种：1 * * * * * 对应 cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
// 参考：https://darjun.github.io/2020/06/25/godailylib/cron/
func AddJob(spec string, cmd func(), mode Mode) (cron.EntryID, error) {
	prefixJobs := make([]cron.JobWrapper, 0, 2)
	switch mode {
	case SkipIfStillRunning:
		prefixJobs = append(prefixJobs, cron.SkipIfStillRunning(cron.DefaultLogger))
	case DelayIfStillRunning:
		prefixJobs = append(prefixJobs, cron.DelayIfStillRunning(cron.DefaultLogger))
	}

	prefixJobs = append(prefixJobs, cron.Recover(cron.DefaultLogger))
	return _cron.AddJob(
		spec,
		cron.NewChain(prefixJobs...).Then(&job{cmd: cmd}),
	)
}
