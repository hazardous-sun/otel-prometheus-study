package postgres

import (
	"database/sql"
	"otel-prometheus-study/internal/domain/shared"
	"otel-prometheus-study/internal/domain/store"
	"otel-prometheus-study/internal/logger"
)

type StoreRepository struct {
	connection *sql.DB
}

func NewStoreRepository(db *sql.DB) StoreRepository {
	return StoreRepository{connection: db}
}

func (sr StoreRepository) InsertStore(s store.Store) (store.Store, error) {
	query := `INSERT INTO stores (name) VALUES ($1) RETURNING id`
	name := s.Name()

	var id int
	err := sr.connection.QueryRow(query, name).Scan(&id)
	if err != nil {
		logger.LogError(err, "context", "InsertStore")
		return store.Store{}, err
	}

	newID, err := shared.NewID(id)
	if err != nil {
		return store.Store{}, err
	}
	s.IDValue = newID
	return s, nil
}

func (sr StoreRepository) GetStores() ([]store.Store, error) {
	query := `SELECT id, name FROM stores`
	rows, err := sr.connection.Query(query)
	if err != nil {
		logger.LogError(err, "context", "GetStores")
		return nil, err
	}
	defer rows.Close()

	var stores []store.Store
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		storeObj, err := store.NewStore(id, name)
		if err != nil {
			return nil, err
		}
		stores = append(stores, storeObj)
	}
	return stores, nil
}
