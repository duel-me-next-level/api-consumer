package repository

type TournamentRepository interface {
	GetByID(id int) (*Tournament, error)
	Create(tournament *Tournament) error
	Update(tournament *Tournament) error
	FindAll() ([]*Tournament, error)
	Delete(id int) error
}
