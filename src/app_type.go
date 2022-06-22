package app

import (
	"api-golang-hex/config"
	"api-golang-hex/src/interface/database/mongodb"
)

type AppContainer struct {
	config  *config.Global
	mongodb *mongodb.Handler
}

type Application interface {
	Init() error
}
