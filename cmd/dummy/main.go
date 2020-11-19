package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/iot-arch/octopus-adaptors/dummy/pkg/adaptor/log"
	"github.com/iot-arch/octopus-adaptors/dummy/pkg/dummy"
	_ "github.com/iot-arch/octopus-adaptors/dummy/pkg/util/log/handler"
	"github.com/iot-arch/octopus-adaptors/dummy/pkg/util/log/logflag"
	"github.com/iot-arch/octopus-adaptors/dummy/pkg/util/version/verflag"
)

const (
	name        = "dummy"
	description = ``
)

func newCommand() *cobra.Command {
	var c = &cobra.Command{
		Use:  name,
		Long: description,
		RunE: func(cmd *cobra.Command, args []string) error {
			verflag.PrintAndExitIfRequested(name)
			logflag.SetLogger(log.SetLogger)

			return dummy.Run()
		},
	}

	verflag.AddFlags(c.Flags())
	logflag.AddFlags(c.Flags())
	return c
}

func main() {
	var c = newCommand()
	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
