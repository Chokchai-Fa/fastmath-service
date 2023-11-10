package repositories

import (
	"database/sql"
	"fastmath/utils/dbpostgres"
	"fastmath/utils/errors"
)

type QuestionRepository interface {
	GetAllQuestion() (*string, error)
}

type questionRepo struct {
	DBPostgres    *dbpostgres.DBPG
	DBPostgresCli *sql.DB
}

func NewQuestionRepositoryDB(dbPostgres *dbpostgres.DBPG, dbPostgresCli *sql.DB) QuestionRepository {
	return &questionRepo{
		DBPostgres:    dbPostgres,
		DBPostgresCli: dbPostgresCli,
	}
}

func (h *questionRepo) GetAllQuestion() (*string, error) {

	print("call repository")

	return nil, errors.NewUnexpectedError()
}
