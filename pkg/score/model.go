package score

import (
	"time"

	"github.com/dreinix/gonake/pkg/auth"
)

type Score struct {
	ID    int       `json:"ID"`
	Value int       `json:"value"`
	User  auth.User `json:"user"`
	Date  time.Time `json:"date"`
}
