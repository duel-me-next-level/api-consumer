package infrastructure

import (
	"database/sql"
	"errors"

	"github.com/duel-me-next-level/api-consumer/internal/core/domain/model"
	repository "github.com/duel-me-next-level/api-consumer/internal/core/domain/ports"
	"github.com/go-pg/pg"
)

type tournamentRepository struct {
	DB *pg.DB
}

func NewTournamentRepository(db *pg.DB) repository.TournamentRepository {
	return &tournamentRepository{DB: db}
}

func (r *tournamentRepository) Create(t *model.Tournament) error {
	_, err := r.DB.Model(t).Insert()
	return err
}

func (r *tournamentRepository) GetByID(id int32) (*model.Tournament, error) {
	t := &model.Tournament{}
	err := r.DB.Model(t).Where("id = ?", id).Select()
	if err == sql.ErrNoRows {
		return nil, errors.New("Tournament not found")
	}
	return t, err
}

func (r *tournamentRepository) FindAll() ([]*model.Tournament, error) {
	var ts []*model.Tournament
	err := r.DB.Model(&ts).Select()
	if err == sql.ErrNoRows {
		return nil, errors.New("Tournament not found")
	}
	return ts, err
}

func (r *tournamentRepository) Update(t *model.Tournament) error {
	_, err := r.DB.Model(t).Where("id = ?", t.ID).Update()
	return err
}

func (r *tournamentRepository) Delete(id int32) error {
	t := &model.Tournament{ID: id}
	_, err := r.DB.Model(t).Where("id = ?", id).Delete()
	return err
}
