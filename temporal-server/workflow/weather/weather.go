package weather

import (
	"fmt"
	"go.temporal.io/sdk/activity"
	"time"

	tmprcli "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"temporal-server/model"
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
	workerCM.RegisterWorkflowWithOptions(w.GetWeatherWorkflow, workflow.RegisterOptions{
		Name: "weather-workflow",
	})
	workerCM.RegisterActivityWithOptions(GetWeatherActivity, activity.RegisterOptions{
		Name: "weather-activity",
	})
	// Start Temporal
	err := workerCM.Start()
	if err != nil {
		fmt.Println("error running temporal worker: ", err)
	}
}

func (w Weather) GetWeatherWorkflow(ctx workflow.Context, city string) ([]model.WeatherData, error) {
	fmt.Println("execute GetWeatherWorkflow")
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    10 * time.Second,
			BackoffCoefficient: 2.0,
			MaximumAttempts:    3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var current model.WeatherData
	// start the activities
	err := workflow.ExecuteActivity(ctx, "weather-activity", city).Get(ctx, &current)
	// wait for activities to complete
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var response []model.WeatherData
	// combine results
	response = append(response, current)

	return response, nil
}
