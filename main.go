package main

import (
	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/cmd"
	"github.com/spf13/cobra"
)

func init() {
	grok.Init()
}

func main() {
	defer func() {
		recover()
	}()

	api := &cobra.Command{
		Use:   "api",
		Short: "Starts street-market-api",
		Run:   cmd.Server,
	}

	root := &cobra.Command{
		Use:   "street-market",
		Short: "Street Market Microservice",
	}

	root.PersistentFlags().String("settings", "config.yaml", "path to config.yaml config file")
	root.AddCommand(api, workers(), healthz())

	root.Execute()
}

func workers() *cobra.Command {
	importStreetMarket := &cobra.Command{
		Use:   "import-data",
		Short: "Import Data",
		Run:   cmd.ImportData,
	}

	importStreetMarket.Flags().String("path", "", "path to import file")
	importStreetMarket.MarkFlagRequired("path")

	workers := &cobra.Command{
		Use:   "workers",
		Short: "Workers",
	}

	workers.AddCommand(importStreetMarket)

	return workers
}

func healthz() *cobra.Command {
	return &cobra.Command{
		Use:   "healthz",
		Short: "Health Check",
		Run:   grok.ConsumerHealthz("settings", grok.WithMongo()),
	}
}
