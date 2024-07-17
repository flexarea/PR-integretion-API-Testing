package models

import (
	"database/sql"
)

type LogsModel struct {
	DB *sql.DB
}

func (m *LogsModel) Insert(title, branch, destinationBranch, Pr_comment, slackchannel string) (int, error) {
	stmt := `INSERT INTO prinfo (title, branch, destinationBranch, pr_comment, slackchannel, created)
	VALUES ($1,$2,$3,$4,$5,NOW() AT TIME ZONE 'Africa/Kigali')
	`

	//executing query
	result, err := m.DB.Exec(stmt, title, branch, destinationBranch, Pr_comment, slackchannel)

}
