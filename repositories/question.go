package repositories

import (
	"database/sql"
	"fastmath/utils/dbpostgres"
)

type QuestionRepository interface {
	GetAllQuestion()
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

func (h *questionRepo) GetAllQuestion() {

	print("call repository")
}
