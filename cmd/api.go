package cmd

import (
	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/container"
	"github.com/GuilhermeFirmiano/street-market/pkg/settings"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Server is a cmd to setup api server
func Server(command *cobra.Command, args []string) {
	settings := new(settings.Settings)
	err := grok.FromYAML(command.Flag("settings").Value.String(), settings)

	if err != nil {
		logrus.WithError(err).
			Panic("error loading settings")
	}

	app := container.New(settings)

	api := grok.New(
		grok.WithCORS(),
		grok.WithContainer(app),
		grok.WithSettings(settings.Grok),
		grok.WithHealthz(
			grok.HTTPHealthz(
				grok.WithMongo(),
				grok.WithHealthzSettings(settings.Grok),
			),
		),
	)

	api.Run()
}
