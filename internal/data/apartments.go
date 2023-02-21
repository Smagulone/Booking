package data

import (
	"database/sql"
	"errors"
	"github.com/Smagulone/Booking/internal/validator"
	"github.com/lib/pq"
	"time"
)

type MovieModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record in the movies table.
func (a Apartment) Insert(apartment *Apartment) error {
	query := `INSERT INTO movies (name, ranking, location, description) 
VALUES ($1, $2, $3, $4)
RETURNING id, created_at`

	args := []any{apartment.Name, apartment.Ranking, apartment.Location, pq.Array(apartment.Description)}

	return a.DB.QueryRow(query, args...).Scan(&apartment.ID, &apartment.CreatedAt)
}

// Add a placeholder method for fetching a specific record from the movies table.
func (a Apartment) Get(id int64) (*Apartment, error) {
	// The PostgreSQL bigserial type that we're using for the movie ID starts
	// auto-incrementing at 1 by default, so we know that no movies will have ID values // less than that. To avoid making an unnecessary database call, we take a shortcut // and return an ErrRecordNotFound error straight away.
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Define the SQL query for retrieving the movie data.
	query := `SELECT id, created_at, title, year, runtime, genres, version
FROM movies
WHERE id = $1`
	// Declare a Movie struct to hold the data returned by the query.
	var apartment Apartment
	// Execute the query using the QueryRow() method, passing in the provided id value // as a placeholder parameter, and scan the response data into the fields of the
	// Movie struct. Importantly, notice that we need to convert the scan target for the // genres column using the pq.Array() adapter function again.
	err := m.DB.QueryRow(query, id).Scan(&apartment.ID, &apartment.CreatedAt, &apartment.Name, &apartment.Ranking, &apartment.Runtime, pq.Array(&apartment.Description), &apartment.Version)
	// Handle any errors. If there was no matching movie found, Scan() will return
	// a sql.ErrNoRows error. We check for this and return our custom ErrRecordNotFound // error instead.
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	// Otherwise, return a pointer to the Movie struct.
	return &apartment, nil
}

// Add a placeholder method for updating a specific record in the movies table.
func (a Apartment) Update(apartment *Apartment) error {
	// Declare the SQL query for updating the record and returning the new version // number.
	query := `
UPDATE movies
SET name = $1, ranking = $2, description = $4, version = version + 1 WHERE id = $5
RETURNING version`
	// Create an args slice containing the values for the placeholder parameters.
	args := []any{apartment.Name,
		apartment.Ranking, pq.Array(apartment.Description), apartment.ID,
	}
	// Use the QueryRow() method to execute the query, passing in the args slice as a // variadic parameter and scanning the new version value into the movie struct. return m.DB.QueryRow(query, args...).Scan(&movie.Version)
}

func (m MovieModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	// Construct the SQL query to delete the record.
	query := `DELETE FROM movies WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	// Call the RowsAffected() method on the sql.Result object to get the number of rows // affected by the query.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// If no rows were affected, we know that the movies table didn't contain a record // with the provided ID at the moment we tried to delete it. In that case we
	// return an ErrRecordNotFound error.
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

type Apartment struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Ranking   int32     `json:"ranking"`

	Location    []string `json:"location"`
	Description []string `json:"description,omitempty"`
	Floors      int32    `json:"floors,omitempty"`
}

func ValidateMovie(v *validator.Validator, apartment *Apartment) {
	v.Check(apartment.Name != "", "title", "must be provided")
	v.Check(len(apartment.Name) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(apartment.Ranking != 0, "year", "must be provided")
	v.Check(validator.Unique(apartment.Description), "genres", "must not contain duplicate values")
}
