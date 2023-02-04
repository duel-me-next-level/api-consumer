package infrastructure

import (
	"database/sql"

	"github.com/duel-me-next-level/api-consumer/internal/domain/tournament"
	"github.com/go-pg/pg"
)

type tournamentRepository struct {
	DB *pg.DB
}

func NewTournamentRepository(db *pg.DB) tournament.TournamentRepository {
	return &tournamentRepository{DB: db}
}

func (r *tournamentRepository) Create(t *tournament.Tournament) error {
	_, err := r.DB.Model(t).Insert()
	return err
}

func (r *tournamentRepository) GetByID(id int) (*tournament.Tournament, error) {
	t := &tournament.Tournament{}
	err := r.DB.Model(t).Where("id = ?", id).Select()
	if err == sql.ErrNoRows {
		return nil, tournament.ErrTournamentNotFound
	}
	return t, err
}

func (r *tournamentRepository) FindAll() ([]*tournament.Tournament, error) {
	var ts []*tournament.Tournament
	err := r.DB.Model(&ts).Select()
	if err == sql.ErrNoRows {
		return nil, tournament.ErrTournamentNotFound
	}
	return ts, err
}

func (r *tournamentRepository) Update(t *tournament.Tournament) error {
	_, err := r.DB.Model(t).Where("id = ?", t.ID).Update()
	return err
}

func (r *tournamentRepository) Delete(id int) error {
	t := &tournament.Tournament{ID: id}
	_, err := r.DB.Model(t).Where("id = ?", id).Delete()
	return err
}
