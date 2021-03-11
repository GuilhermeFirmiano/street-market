package cmd

import (
	"context"
	"sync"

	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/container"
	"github.com/GuilhermeFirmiano/street-market/pkg/models"
	"github.com/GuilhermeFirmiano/street-market/pkg/settings"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ImportData ...
func ImportData(command *cobra.Command, args []string) {
	settings := new(settings.Settings)
	err := grok.FromYAML(command.Flag("settings").Value.String(), settings)

	if err != nil {
		logrus.WithError(err).
			Panic("error loading settings")
	}

	container := container.New(settings)

	result, err := container.FileService.ReadFileCSV(command.Flag("path").Value.String())

	if err != nil {
		logrus.WithError(err).
			Panic("error loading file")
	}

	logrus.Infof("found %d street market to import", len(result))

	wg := new(sync.WaitGroup)
	wg.Add(len(result) - 1)

	for i, data := range result {
		line := i + 1
		if line == 1 {
			continue
		}

		go func(line int, data []string) {
			defer wg.Done()
			model, err := models.ParseCSVToModel(data)

			if err != nil {
				logrus.WithField("line", line).
					WithError(err).
					Error("error when building model")
				return
			}

			_, err = container.StreetMarketService.Create(context.Background(), model)

			if err != nil {
				logrus.WithField("line", line).
					WithField("registry", model.Registry).
					WithError(err).
					Error("error creating street market")
				return
			}

			logrus.WithField("line", line).
				WithField("registry", model.Registry).
				Info("street market created")

		}(line, data)
	}

	wg.Wait()

	logrus.Info("import completed")
}
