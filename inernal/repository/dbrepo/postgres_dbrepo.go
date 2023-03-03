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

func (m *PostgresDBRepo) OneMovie(id int) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select 
			id, title, mpaa_rating, release_date, runtime,
			description, coalesce(image, ''), created_at, updated_at
		from
		    movies
		where id=$1
	`
	row := m.Db.QueryRowContext(ctx, query, id)
	var movie models.Movie
	err := row.Scan(
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

	query = `
		select 
		    g.id, g.genre from movies_genres mg
		left join genres g on (mg.genre_id = g.id)
		where mg.movie_id = $1
		order by g.genre
	`
	rows, err := m.Db.QueryContext(ctx, query, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()
	var genres []*models.Genre
	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			&g.ID,
			&g.Genre,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &g)
	}
	movie.Genres = genres
	return &movie, nil
}

func (m *PostgresDBRepo) OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select 
			id, title, mpaa_rating, release_date, runtime,
			description, coalesce(image, ''), created_at, updated_at
		from
		    movies
		where id=$1
	`
	row := m.Db.QueryRowContext(ctx, query, id)
	var movie models.Movie
	err := row.Scan(
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
		return nil, nil, err
	}

	query = `
		select 
		    g.id, g.genre from movies_genres mg
		left join genres g on (mg.genre_id = g.id)
		where mg.movie_id = $1
		order by g.genre
	`
	rows, err := m.Db.QueryContext(ctx, query, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err
	}
	defer rows.Close()
	var genres []*models.Genre
	var genresArray []int
	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			&g.ID,
			&g.Genre,
		)
		if err != nil {
			return nil, nil, err
		}
		genres = append(genres, &g)
		genresArray = append(genresArray, g.ID)
	}
	movie.Genres = genres
	movie.GenresArray = genresArray
	var allGenres []*models.Genre
	query = "select id, genre from genres order by genre"
	grows, err := m.Db.QueryContext(ctx, query)
	for grows.Next() {
		var g models.Genre
		err := grows.Scan(
			&g.ID,
			&g.Genre,
		)
		if err != nil {
			return nil, nil, err
		}
		allGenres = append(allGenres, &g)
	}
	return &movie, allGenres, nil
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

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			id, email, first_name, last_name, password, created_at, updated_at
		from
		    users
		where email=$1
	`
	var user models.User
	row := m.Db.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdateAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *PostgresDBRepo) GetUserById(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			id, email, first_name, last_name, password, created_at, updated_at
		from
		    users
		where id=$1
	`
	var user models.User
	row := m.Db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdateAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
