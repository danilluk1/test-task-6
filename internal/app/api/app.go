package api

import (
	db "github.com/danilluk1/test-task-6/db/sqlc"
	"github.com/danilluk1/test-task-6/internal/services/logger"
)

type App struct {
	Store  db.Store
	Logger logger.Logger
}
