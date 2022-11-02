package data

import (
	"github.com/diegoclair/go_boilerplate/domain/contract"
	"github.com/diegoclair/go_boilerplate/infra/config"
	"github.com/diegoclair/go_boilerplate/infra/data/mysql"
	"github.com/diegoclair/go_boilerplate/infra/logger"
)

// Connect returns a instace of mysql db
func Connect(cfg *config.Config, log logger.Logger) (contract.DataManager, error) {
	return mysql.Instance(cfg, log)
}
