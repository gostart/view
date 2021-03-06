package view

import (
	"github.com/ungerik/go-start/errs"
)

func newSession(ctx *Context) *Session {
	session := &Session{
		Tracker:   ctx.Server.SessionTracker,
		DataStore: ctx.Server.SessionDataStore,
		Ctx:       ctx,
	}
	if session.Tracker == nil {
		session.Tracker = new(CookieSessionTracker)
	}
	if session.DataStore == nil {
		session.DataStore = NewCookieSessionDataStore()
	}
	return session
}

type Session struct {
	Tracker   SessionTracker
	DataStore SessionDataStore
	Ctx       *Context

	/*
		Cached user object of the session.
		User won't be set automatically, use user.OfSession(context) instead.

		Example for setting it automatically for every request:

			import "github.com/ungerik/go-start/user"

			Config.OnPreAuth = func(ctx *Context) error {
				user.OfSession(context) // Sets context.User
				return nil
			}
	*/
	// User interface{}

	cachedID string
}

// ID returns the id of the session or an empty string.
// It's valid to call this method on a nil pointer.
func (session *Session) ID() string {
	if session == nil {
		return ""
	}
	if session.cachedID != "" {
		return session.cachedID
	}
	session.cachedID = session.Tracker.ID(session.Ctx)
	return session.cachedID
}

func (session *Session) SetID(id string) {
	if session.Tracker != nil {
		session.Tracker.SetID(session.Ctx, id)
		session.cachedID = id
	}
}

func (session *Session) DeleteID() {
	session.cachedID = ""
	if session.Tracker == nil {
		return
	}
	session.Tracker.DeleteID(session.Ctx)
}

// SessionData returns all session data in out.
func (session *Session) Data(out interface{}) (ok bool, err error) {
	if session.DataStore == nil {
		return false, errs.Format("Can't get session data without gostart/views.Config.SessionDataStore")
	}
	return session.DataStore.Get(session.Ctx, out)
}

// SetSessionData sets all session data.
func (session *Session) SetData(data interface{}) (err error) {
	if session.DataStore == nil {
		return errs.Format("Can't set session data without gostart/views.Config.SessionDataStore")
	}
	return session.DataStore.Set(session.Ctx, data)
}

// DeleteSessionData deletes all session data.
func (session *Session) DeleteData() (err error) {
	if session.DataStore == nil {
		return errs.Format("Can't delete session data without gostart/views.Config.SessionDataStore")
	}
	return session.DataStore.Delete(session.Ctx)
}
