package container

import (
	"context"

	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/controllers"
	"github.com/GuilhermeFirmiano/street-market/pkg/repositories"
	"github.com/GuilhermeFirmiano/street-market/pkg/services"
	"github.com/GuilhermeFirmiano/street-market/pkg/settings"
	"go.mongodb.org/mongo-driver/mongo"
)

// Container wrapps app dependencies
type Container struct {
	settings *settings.Settings

	MongoClient *mongo.Client

	FileService *services.FileService

	StreetMarketRepository *repositories.StreetMarketRepository
	StreetMarketService    *services.StreetMarketService
	StreetMarketController *controllers.StreetMarketController
}

// New creates a new context instance
func New(settings *settings.Settings) *Container {
	container := new(Container)

	container.settings = settings

	container.MongoClient = grok.NewMongoConnection(
		container.settings.Grok.Mongo.ConnectionString,
		container.settings.Grok.Mongo.CaFilePath,
	)

	container.StreetMarketRepository = repositories.NewStreetMarketRepository(
		container.MongoClient,
		container.settings.Grok.Mongo.Database,
	)

	container.StreetMarketService = services.NewStreetMarketService(
		container.StreetMarketRepository,
	)

	container.StreetMarketController = controllers.NewStreetMarketController(
		container.StreetMarketService,
	)

	container.FileService = services.NewFileService()

	return container
}

// Close finishes context connections
func (ctx *Container) Close() error {
	return ctx.MongoClient.Disconnect(context.Background())
}

// Controllers returns all context controllers
func (ctx *Container) Controllers() []grok.APIController {
	return []grok.APIController{
		ctx.StreetMarketController,
	}
}
