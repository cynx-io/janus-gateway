package session

import (
	"encoding/gob"
	"github.com/cynx-io/janus-gateway/internal/dependencies/auth0"
	"net/http"
)

type UserSession struct {
	UserID        string `json:"user_id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Authenticated bool   `json:"authenticated"`
}

func init() {
	gob.Register(UserSession{})
}

func GetSession(r *http.Request) (*UserSession, error) {
	session, err := auth0.Store.Get(r, "auth-session")
	if err != nil {
		return nil, err
	}

	userSession := &UserSession{}

	if userID, ok := session.Values["user_id"].(string); ok {
		userSession.UserID = userID
	}
	if email, ok := session.Values["email"].(string); ok {
		userSession.Email = email
	}
	if name, ok := session.Values["name"].(string); ok {
		userSession.Name = name
	}
	if auth, ok := session.Values["authenticated"].(bool); ok {
		userSession.Authenticated = auth
	}
	if state, ok := session.Values["state"].(string); ok && state != "" {
		// Include state for OAuth flow validation
		session.Values["state"] = state
	}

	return userSession, nil
}

func SetSession(w http.ResponseWriter, r *http.Request, userSession *UserSession) error {
	session, err := auth0.Store.Get(r, "auth-session")
	if err != nil {
		return err
	}

	session.Values["user_id"] = userSession.UserID
	session.Values["email"] = userSession.Email
	session.Values["name"] = userSession.Name
	session.Values["authenticated"] = userSession.Authenticated

	return session.Save(r, w)
}

func SetState(w http.ResponseWriter, r *http.Request, state string) error {
	session, err := auth0.Store.Get(r, "auth-session")
	if err != nil {
		return err
	}

	session.Values["state"] = state
	return session.Save(r, w)
}

func GetState(r *http.Request) (string, error) {
	session, err := auth0.Store.Get(r, "auth-session")
	if err != nil {
		return "", err
	}

	if state, ok := session.Values["state"].(string); ok {
		return state, nil
	}
	return "", nil
}

func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := auth0.Store.Get(r, "auth-session")
	if err != nil {
		return err
	}

	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1
	return session.Save(r, w)
}
