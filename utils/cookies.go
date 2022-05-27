package utils

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

type Store struct {
	sessions.Store
}

var StoreInstance = toStore(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY"))))


func (store Store) SetCookie(w http.ResponseWriter, r *http.Request, cookieName interface{}, cookieValue interface{}) {
	session, _ := store.Get(r, "session")
	session.Values[cookieName] = cookieValue
	session.Save(r, w)
}

func (store Store) GetCookie(r *http.Request, cookieName string) interface{} {
	session, _ := store.Get(r, "session")
	return session.Values[cookieName]
}

func toStore(store sessions.Store) Store {
	return Store{store}
}
