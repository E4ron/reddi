package repository

import (
	"github.com/jmoiron/sqlx"
	"reddi/models"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (s *PostPostgres) GetById(id string) (models.Post, error) {

}
func (s *PostPostgres) GetList(page int, limit int) (models.OutputPostList, error) {

}
func (s *PostPostgres) Create(post models.InputPost) (models.OutputPost, error) {

}
func (s *PostPostgres) Update(post models.InputUpdatePost) error {

}
func (s *PostPostgres) Delete(id string) error {

}
