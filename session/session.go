package session

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var Session = map[string]string{}
var lock sync.RWMutex

func CreateMailCode(email, value string) {
	ECode := email + "Code"
	ELastSend := email + "LastSend"
	lock.Lock()
	Session[ECode] = value
	Session[ELastSend] = "f"
	lock.Unlock()
	SetExpired(ECode, 300)
	SetExpired(ELastSend, 60)
}

func GetMailCode(email string) string {
	email += "Code"
	if _, ok := Session[email]; ok {
		return Session[email]
	}
	return ""
}
func GetMailLast(email string) string {
	email += "CodeLastSend"
	if _, ok := Session[email]; ok {
		return Session[email]
	}
	return ""
}

func CreateSession(value string, expired int) string {
	id := strconv.Itoa(time.Now().Second() * rand.Int())
	lock.Lock()
	if _, ok := Session[id]; ok {
		CreateSession(value, expired)
	} else {
		Session[id] = value
	}
	lock.Unlock()
	if expired != 0 {
		SetExpired(id, expired)
	}
	return id
}

func GetSession(id string) string {
	if _, ok := Session[id]; ok {
		return Session[id]
	} else {
		return ""
	}
}

func ChangeSession(id string, value string) error {
	if _, ok := Session[id]; ok {
		lock.Lock()
		Session[id] = value
		lock.Unlock()
		return nil
	}
	return errors.New("not found session")
}

func DeleteSession(id string) {
	lock.Lock()
	delete(Session, id)
	lock.Unlock()
}

func SetExpired(id string, expired int) {
	go func() {
		time.Sleep(time.Duration(expired) * time.Second)
		lock.Lock()
		DeleteSession(id)
		lock.Unlock()
	}()
}
