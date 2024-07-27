package workflow

import (
	tmprcli "go.temporal.io/sdk/client"

	"temporal-server/workflow/weather"
)

type Workflow struct {
	tprCli    tmprcli.Client
	weatherWF *weather.Weather
}

func NewWorkflow(tprCli tmprcli.Client) *Workflow {
	return &Workflow{
		tprCli:    tprCli,
		weatherWF: weather.New(tprCli),
	}
}
func (wf *Workflow) Init() {
	wf.weatherWF.RegisterWF()
}
