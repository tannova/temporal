package workflow

import (
	tmprcli "go.temporal.io/sdk/client"
	"temporal-server/workflow/schedule"

	"temporal-server/workflow/weather"
)

type Workflow struct {
	tprCli    tmprcli.Client
	weatherWF *weather.Weather
	schedule  *schedule.Schedule
}

func NewWorkflow(tprCli tmprcli.Client) *Workflow {
	return &Workflow{
		tprCli:    tprCli,
		weatherWF: weather.New(tprCli),
		schedule:  schedule.New(tprCli),
	}
}
func (wf *Workflow) Init() {
	wf.weatherWF.RegisterWF()
	wf.schedule.RegisterWF()
	wf.schedule.Start()
	wf.schedule.Start123()
}
