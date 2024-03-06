package router

import (
	"github.com/gin-gonic/gin"
	gerror "github.com/gotify/server/v2/error"
	"github.com/hoangndst/cloud-observability-backend/config"
	"github.com/hoangndst/cloud-observability-backend/database"
	"github.com/hoangndst/cloud-observability-backend/model"
)

func Create(db *database.GormDatabase, vInfo *model.VersionInfo, config *config.Configuration) (*gin.Engine, func()) {
	g := gin.New()

	g.RemoteIPHeaders = []string{"X-Forwarded-For"}
	g.SetTrustedProxies(config.Server.TrustedProxies)

	g.NoRoute(gerror.NotFound())
	//authentication := auth.Auth{DB: db}
	return g, nil
}
