package cron

import (
	"github.com/robfig/cron/v3"
)

// CronSrv
// 每隔5秒执行一次：*/5 * * * * ?
// 每隔1分钟执行一次：0 */1 * * * ?
// 每天23点执行一次：0 0 23 * * ?
// 每天凌晨1点执行一次：0 0 1 * * ?
// 每月1号凌晨1点执行一次：0 0 1 1 * ?
// 在26分、29分、33分执行一次：0 26,29,33 * * * ?
// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
type CronSrv struct {
	cron  *cron.Cron
	funds map[string]cron.EntryID
	jobs  map[string]cron.EntryID
}

func NewCronSrv() *CronSrv {
	secondParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return &CronSrv{
		cron:  cron.New(cron.WithParser(secondParser), cron.WithChain()),
		funds: make(map[string]cron.EntryID),
	}
}

func (c *CronSrv) Start() {
	c.cron.Start()
}

// AddFunc   例：  spec := "*/5 * * * * ?"， 每5秒执行一次
func (c *CronSrv) AddFunc(spec string, cmd func()) (entryID cron.EntryID, err error) {
	return c.cron.AddFunc(spec, cmd)
}

func (c *CronSrv) AddOnFunc(spec string, cmd func()) (entryID cron.EntryID, err error) {
	var ok bool
	if entryID, ok = c.funds[spec]; ok {
		return entryID, err
	}

	entryID, err = c.AddFunc(spec, cmd)
	c.funds[spec] = entryID
	return entryID, err
}

func (c *CronSrv) AddJob(spec string, cmd func()) (entryID cron.EntryID, err error) {
	return c.AddJob(spec, cmd)
}

func (c *CronSrv) AddOnceJob(spec string, cmd func()) (entryID cron.EntryID, err error) {
	var ok bool
	if entryID, ok = c.jobs[spec]; ok {
		return entryID, err
	}

	entryID, err = c.AddOnFunc(spec, cmd)
	c.jobs[spec] = entryID
	return
}

func (c *CronSrv) Stop() {
	c.cron.Stop()
}
