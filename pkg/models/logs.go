package models

import (
	"database/sql"
)

type LogsModel struct {
	DB *sql.DB
}

func (m *LogsModel) Insert(title, branch, destinationBranch, Pr_comment, slackchannel string) (int, error) {

	var lastId int

	stmt := `INSERT INTO prinfo (title, branch, destinationBranch, pr_comment, slackchannel, created)
	VALUES ($1,$2,$3,$4,$5,NOW() AT TIME ZONE 'Africa/Kigali')
	RETURNING id
	`

	//executing query
	err := m.DB.QueryRow(stmt, title, branch, destinationBranch, Pr_comment, slackchannel).Scan(&lastId)

	if err != nil {
		return 0, err
	}

	return lastId, err

}
