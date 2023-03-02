package api

import db "github.com/danilluk1/test-task-6/db/sqlc"

type App struct {
	Store db.Store
}
