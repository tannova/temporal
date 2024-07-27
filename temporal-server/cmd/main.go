package main

import (
	"fmt"

	tmprcli "go.temporal.io/sdk/client"

	"temporal-server/workflow"
)

func main() {
	tprCli, err := tmprcli.Dial(tmprcli.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// init & register workflows
	workflow.NewWorkflow(tprCli).Init()

	fmt.Println("start service success")
	select {}
}
