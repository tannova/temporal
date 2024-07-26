package main

import (
	tmprcli "go.temporal.io/sdk/client"
	"temporal-server/workflow"
)

func main() {
	tprCli, err := tmprcli.Dial(tmprcli.Options{
		HostPort: "localhost:",
	})
	if err != nil {
		return
	}

	workflow.NewWorkflow(tprCli)
}
