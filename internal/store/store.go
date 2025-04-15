package store

import (
	"database/sql"

	_ "github.com/lib/pq"

	"go-rest-api/internal/config"
	"go-rest-api/internal/repositories"
)

type Store struct {
	config 		   *config.StoreConfig
	db 			   *sql.DB
	userRepository *repositories.UserRepository
}

func New(config *config.StoreConfig) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	
	err = db.Ping()
	if err != nil {
		return err
	}
	
	s.db = db
	
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

// Реализация интерфейса StoreInterface
func (s *Store) DB() *sql.DB {
    return s.db
}

func (s *Store) User() *repositories.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	
	s.userRepository = &repositories.UserRepository{
		Store: s,
	}
	
	return s.userRepository
}
