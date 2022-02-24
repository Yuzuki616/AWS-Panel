package session

import (
	"errors"
	"math/rand"
	"time"
)

var Session = map[int]string{}

func CreateSession(value string) int {
	id := time.Now().Second() * rand.Int()
	if _, ok := Session[id]; ok {
		CreateSession(value)
	} else {
		Session[id] = value
	}
	return id
}

func GetSession(id int) string {
	if _, ok := Session[id]; ok {
		return Session[id]
	} else {
		return ""
	}
}

func ChangeSession(id int, value string) error {
	if _, ok := Session[id]; ok {
		Session[id] = value
		return nil
	}
	return errors.New("not found session")
}

func DeleteSession(id int) {
	delete(Session, id)
}
