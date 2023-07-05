package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	 "github.com/thisisrahmat/redditgo"
)

type ThreadStore struct {
	// imbedding a type into our struct 
	*sqlx.DB
}


//This function initialises our new threadstore then returns it 
func NewThreadStore (db *sqlx.DB) *ThreadStore{

	return &ThreadStore{
		DB: db,
	}
}


//thread method quereies a thread 
func (s *ThreadStore) Thread(id uuid.UUID) (redditgo.Thread, error){
	panic("not implemented") //ToDO: Implemented
}

func (s *ThreadStore) Threads() ([]redditgo.Thread, error){
	panic("not implemented") //ToDO: Implemented
}

func (s *ThreadStore) CreateThread(t *redditgo.Thread) error {
	panic("not implemented") //ToDO: Implemented
}

func (s *ThreadStore) UpdateThread(t *redditgo.Thread)error {
	panic("not implemented") //ToDO: Implemented
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	panic("not implemented") //ToDO: Implemented
}
