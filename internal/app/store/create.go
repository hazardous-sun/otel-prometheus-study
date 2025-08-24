package store

import (
	"otel-prometheus-study/internal/domain/store"
	"otel-prometheus-study/internal/infra/postgres"
)

func CreateStore(name string) (store.Store, error) {
	db, err := postgres.ConnectDB()
	if err != nil {
		return store.Store{}, err
	}
	defer db.Close()

	repo := postgres.NewStoreRepository(db)
	newStore, err := store.NewStore(0, name)
	if err != nil {
		return store.Store{}, err
	}

	return repo.InsertStore(newStore)
}
