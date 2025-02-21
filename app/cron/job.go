package cron

import (
	"mangoTV/app/svc"
)

type JobRouter struct {
	ctx *svc.ServiceContext
}

func NewJobRouter(ctx *svc.ServiceContext) *JobRouter {
	return &JobRouter{ctx: ctx}
}

func (j *JobRouter) StartJob() {
	cronLogic := NewCronSrv()

	_, _ = cronLogic.AddFunc("*/5 * * * * ?", j.CronFiveSecondEvent)
	_, _ = cronLogic.AddFunc("*/10 * * * * ?", j.CronTenSecondEvent)
	_, _ = cronLogic.AddFunc("*/20 * * * * ?", j.CronTwentySecondEvent)
	_, _ = cronLogic.AddFunc("*/30 * * * * ?", j.CronThirtySecondEvent)
	_, _ = cronLogic.AddFunc("0 */1 * * * ?", j.CronMinuteEvent)        //每分钟
	_, _ = cronLogic.AddFunc("0 */3 * * * ?", j.CronThreeMinuteEvent)   //每3分钟
	_, _ = cronLogic.AddFunc("0 */5 * * * ?", j.CronFiveMinuteEvent)    //每5分钟
	_, _ = cronLogic.AddFunc("0 */10 * * * ?", j.CronTenMinuteEvent)    //每10分钟
	_, _ = cronLogic.AddFunc("0 0 */1 * * ?", j.CronHourEvent)          //每小时
	_, _ = cronLogic.AddFunc("0 0 */3 * * ?", j.CronEachThreeHourEvent) //每3小时
	_, _ = cronLogic.AddFunc("0 0 1 * * ?", j.CronDay1Event)            //每天

	cronLogic.Start()
}
