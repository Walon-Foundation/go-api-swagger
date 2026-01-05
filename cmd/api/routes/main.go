package routes

import (
	"github.com/Walon-Foundation/go-gin-doc/cmd/api/config"
	"github.com/Walon-Foundation/go-gin-doc/cmd/models"
)

type route struct {
	Models models.Models
}

func Router()*route{
	model := models.NewModel(config.Db)
	return &route{
		Models: model,
	}
}