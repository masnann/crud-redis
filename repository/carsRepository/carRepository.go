package carsRepository

import (
	"crud-redis/models"
	"crud-redis/repository"
	"database/sql"
	"errors"
	"log"
)

type CarsRepository struct {
	repo repository.Repository
}

func NewCarsRepository(repo repository.Repository) CarsRepository {
	return CarsRepository{
		repo: repo,
	}
}

var defineColumn = `id, brand, type, color`

func carsDto(rows *sql.Rows) ([]models.Cars, error) {
	var result []models.Cars
	for rows.Next() {
		var val models.Cars
		err := rows.Scan(&val.ID, &val.Brand, &val.Type, &val.Color)
		if err != nil {
			log.Println("Error carsDto :", err)
			return result, err
		}
		result = append(result, val)

	}
	return result, nil

}

func (r CarsRepository) FindCarsByID(id int) (models.Cars, error) {
	var cars models.Cars

	query := `SELECT ` + defineColumn + ` FROM cars WHERE id = $1`

	rows, err := r.repo.DB.Query(query, id)
	if err != nil {
		log.Println("Error query :", err)
		return cars, err
	}
	defer rows.Close()

	data, err := carsDto(rows)
	if len(data) == 0 {
		log.Println("Error carsDto FindCarsByID: ", err)
		return cars, errors.New("cars not found")
	}
	return data[0], nil
}

func (r CarsRepository) GetAllCars() ([]models.Cars, error) {
	query := `SELECT ` + defineColumn + ` FROM cars`

	rows, err := r.repo.DB.Query(query)
	if err != nil {
		log.Println("Error query: ", err)
		return nil, err
	}

	defer rows.Close()

	data, err := carsDto(rows)
	if len(data) == 0 {
		log.Println("Error carsDto")
		return nil, err 
	}
	return data, nil
	
}