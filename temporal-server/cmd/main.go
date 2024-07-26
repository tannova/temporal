package cmd

import (
	tmprcli "go.temporal.io/sdk/client"
)

func main() {

	tprCli, err := tmprcli.Dial(tmprcli.Options{
		HostPort: "localhost:",
	})
	if err != nil {
		return
	}

}
