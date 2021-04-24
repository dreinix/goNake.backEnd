package score

import (
	"time"

	"github.com/dreinix/gonake/pkg/user"
)

type Score struct {
	ID    int       `json:"ID"`
	Value int       `json:"value"`
	User  user.User `json:"user"`
	Date  time.Time `json:"date"`
}
