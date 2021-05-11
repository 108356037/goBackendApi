package sesscookie

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	sessionName = "golang_cookie"
	authKey     = []byte("qwerqwerqwerqwer")
	encKey      = []byte("asdfasdfasdfasdf")

	store      = sessions.NewCookieStore(authKey, encKey)
	appSession *sessions.Session
)

func InitSession(r *http.Request) *sessions.Session {
	// log.Println("session before get", appSession)

	if appSession != nil {
		return appSession
	}

	store.Options = &sessions.Options{
		Path:     "/",      // to match all requests
		MaxAge:   3600 * 1, // 1 hour
		HttpOnly: true,
	}

	session, err := store.Get(r, sessionName)
	appSession = session

	// log.Println("session after get", session)
	if err != nil {
		panic(err)
	}
	return session
}

func SetSessionValue(w http.ResponseWriter, r *http.Request, key, value interface{}) error {
	session := InitSession(r)
	session.Values[key] = value
	fmt.Printf("set session with key %s and value %s\n", key, value)
	if err := session.Save(r, w); err != nil {
		return err
	}
	return nil
}

func GetSessionValue(w http.ResponseWriter, r *http.Request, key interface{}) interface{} {
	session := InitSession(r)
	valWithOutType := session.Values[key]
	// fmt.Printf("valWithOutType: %v\n", valWithOutType)
	value, ok := valWithOutType.(string)
	// log.Println("returned value: ", value)

	if !ok {
		fmt.Printf("cannot get session value by key: %v\n", key)
		return nil
	}

	return value
}
