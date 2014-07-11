package view

import (
	"bytes"
	"encoding/gob"

	"github.com/ungerik/go-start/errs"
)

///////////////////////////////////////////////////////////////////////////////
// SessionDataStore

type SessionDataStore interface {
	Get(ctx *Context, data interface{}) (ok bool, err error)
	Set(ctx *Context, data interface{}) (err error)
	Delete(ctx *Context) (err error)
}

///////////////////////////////////////////////////////////////////////////////
// CookieSessionDataStore

func NewCookieSessionDataStore() SessionDataStore {
	return &CookieSessionDataStore{"session_"}
}

type CookieSessionDataStore struct {
	cookieNameBase string
}

func (sessionDataStore *CookieSessionDataStore) cookieName(sessionID string) string {
	return sessionDataStore.cookieNameBase + sessionID
}

func (sessionDataStore *CookieSessionDataStore) Get(ctx *Context, data interface{}) (ok bool, err error) {
	sessionID := ctx.Session.ID()
	if sessionID == "" {
		return false, errs.Format("Can't set session data without a session id")
	}

	cookieValue, ok := ctx.Request.SiteCookieBytes(sessionDataStore.cookieName(sessionID))
	if !ok {
		return false, nil
	}

	decoder := gob.NewDecoder(bytes.NewBuffer(cookieValue))
	err = decoder.Decode(data)
	return err == nil, err
}

func (sessionDataStore *CookieSessionDataStore) Set(ctx *Context, data interface{}) (err error) {
	sessionID := ctx.Session.ID()
	if sessionID == "" {
		return errs.Format("Can't set session data without a session id")
	}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	dataBytes := buffer.Bytes()

	if len(dataBytes) > 4000 { // Max HTTP header size is 4094 minus some space for protocol
		return errs.Format("Session %s data size %d is larger than cookie limit of 4000 bytes", sessionID, len(dataBytes))
	}

	ctx.Response.SetSiteCookieBytes(sessionDataStore.cookieName(sessionID), dataBytes)
	return nil
}

func (sessionDataStore *CookieSessionDataStore) Delete(ctx *Context) (err error) {
	sessionID := ctx.Session.ID()
	if sessionID == "" {
		return errs.Format("Can't delete session data without a session id")
	}

	ctx.Response.DeleteSiteCookie(sessionDataStore.cookieName(sessionID))
	return nil
}
