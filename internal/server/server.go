package accountApi

import (
	"github.com/gin-gonic/gin"
)

var Server *API

func init() {
	Server = new(API)
	engine := gin.Default()
	SetRoutes(engine)
	Server.Engine = engine
}

type API struct {
	Engine *gin.Engine
}

func (api *API) Run() error {
	err := api.Engine.Run(":3500")
	return err
}
