package models

import (
	"errors"
	"time"
)

var err = errors.New("error")

type PR_INFO struct {
	ID                int
	Title             string
	Branch            string
	DestinationBranch string
	PR_Comment        string
	SlackChannel      string
	Created           time.Time
}
