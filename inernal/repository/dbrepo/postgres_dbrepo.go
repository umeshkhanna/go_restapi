package dbrepo

import (
	"context"
	"database/sql"
	"go-restapi/inernal/models"
	"time"
)

type PostgresDBRepo struct {
	Db *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.Db
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select
			id, title, mpaa_rating, release_date, runtime,
			description, coalesce(image, ''), created_at,
			updated_at
		from
			movies
		order by 
		    title
	`
	rows, err := m.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var allmovies []*models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.MPAARating,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		allmovies = append(allmovies, &movie)
	}
	return allmovies, nil
}
