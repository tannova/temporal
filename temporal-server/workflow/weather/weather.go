package weather

import (
	tmprcli "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type Weather struct {
	tprCli tmprcli.Client
}

func New(tprCli tmprcli.Client) *Weather {
	return &Weather{
		tprCli: tprCli,
	}
}

func (w Weather) RegisterWF() {
	workerCM := worker.New(w.tprCli, "weather", worker.Options{})
	workerCM.RegisterWorkflow(w.GetWeatherWorkflow)
}

func (w Weather) GetWeatherWorkflow(ctx workflow.Context) error {

	return nil
}
