package schedule

import (
	"context"
	"fmt"
	tmprcli "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

type Schedule struct {
	tprCli tmprcli.Client
}

func New(tprCli tmprcli.Client) *Schedule {
	return &Schedule{
		tprCli: tprCli,
	}
}
func (w Schedule) RegisterWF() {
	workerCM := worker.New(w.tprCli, "schedule", worker.Options{})
	workerCM.RegisterWorkflow(SimpleWorkflow)
	// Start Temporal
	err := workerCM.Start()
	if err != nil {
		fmt.Println("error running temporal worker: ", err)
	}
}

func (s *Schedule) Start123() {

	ctx := context.Background()

	_, err := s.tprCli.ExecuteWorkflow(ctx, tmprcli.StartWorkflowOptions{
		ID:           "test-cron-123",
		TaskQueue:    "schedule",
		CronSchedule: "*/3 * * * *",
	}, "SimpleWorkflow", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Schedule) Start() {

	ctx := context.Background()

	scheduleHandle, err := s.tprCli.ScheduleClient().Create(ctx, tmprcli.ScheduleOptions{
		ID: "test-5",
		Spec: tmprcli.ScheduleSpec{
			//Calendars: []tmprcli.ScheduleCalendarSpec{
			//	{
			//		Minute: []tmprcli.ScheduleRange{
			//			{
			//				Start: 1,
			//				End:   2,
			//			},
			//		},
			//	},
			//},
		},
		Action: &tmprcli.ScheduleWorkflowAction{
			ID:        "test-5",
			Workflow:  SimpleWorkflow,
			TaskQueue: "schedule",
		},
	})
	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}
	ids := scheduleHandle.GetID()
	log.Println("Schedule created successfully", ids)
}

func SimpleWorkflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("Workflow is running...")
	// Mô phỏng một công việc tốn thời gian
	workflow.Sleep(ctx, time.Second*50)
	workflow.GetLogger(ctx).Info("Workflow completed")
	return nil
}
