package store

import (
	"otel-prometheus-study/internal/domain/store"
	"otel-prometheus-study/internal/infra/postgres"
)

func ListStores() ([]store.Store, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := postgres.NewStoreRepository(db)
	return repo.GetStores()
}
