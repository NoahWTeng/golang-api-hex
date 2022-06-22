package injector

import (
	"api-golang-hex/config"
	"api-golang-hex/src/interface/database/mongodb"
)

func MongodbProvider(config *config.Global) *mongodb.Handler {
	return mongodb.NewMongodbConnection(config.MongodbVariables)
}
