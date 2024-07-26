package workflow

import tmprcli "go.temporal.io/sdk/client"

type Workflow struct {
	tprCli tmprcli.Client
}
