package repository

import (
	"database/sql"
	"go-restapi/inernal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
