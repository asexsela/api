package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Для того что бы отработала функция init
)

// Цель данной модели определить:
// 1. инстанс хранилища
// 2. конструктор хранилища
// 3. публичный метод Open (установка соединения)
// 4. публичный метод Close (закрытие соединения)

//Instance
type Storage struct {
	config *Config
	//Database file descriptor
	db *sql.DB
	//Subfield for repo interfacing (model user)
	userRepository *UserRepository
	//Subfield for repo interfacing (model user)
	articleRepository *ArticleRepository
}

//Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

//Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)

	if err != nil {
		return err
	}

	err = db.Ping()

	if err != nil {
		return err
	}

	storage.db = db

	log.Println("Database connection created successfully!")

	return nil
}

//Close connection method
func (storage *Storage) Close() {
	storage.db.Close()
}

//Public repo for User
func (s *Storage) User() *UserRepository {

	s.userRepository = &UserRepository{
		storage: s,
	}

	if s.userRepository != nil {
		return s.userRepository
	}

	return nil
}

//Public repo for Article
func (s *Storage) Article() *ArticleRepository {

	s.articleRepository = &ArticleRepository{
		storage: s,
	}

	if s.articleRepository != nil {
		return s.articleRepository
	}

	return nil
}
