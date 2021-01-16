package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/utils"
	"github.com/go-fed/httpsig"
)

var errIsLocal = errors.New("is local account")
var errIsRemote = errors.New("is remote account")
var errNotLoggedIn = errors.New("not logged in")

type accountContext struct {
	context.Context
	account *models.Actor
}

func (ctx accountContext) Value(key interface{}) interface{} {
	if str, ok := key.(string); ok && str == "account" {
		return ctx.account
	} else {
		return ctx.Context.Value(key)
	}
}

// SetAccount stores an Account pointer into http.Request.Context.
func SetAccount(r *http.Request, account *models.Actor) *http.Request {
	return r.WithContext(accountContext{
		Context: r.Context(),
		account: account,
	})
}

// GetAccount retrieves the Account pointer from http.Request.Context.
func GetAccount(r *http.Request) *models.Actor {
	if account, ok := r.Context().Value("account").(*models.Actor); ok {
		return account
	}
	return nil
}

// Authenticate a middleware
func Authenticate(logger *log.Logger, siteContext db.SiteContext) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if account, err := authHTTPSSignatures(siteContext, r); err == nil {
				SetAccount(r, account)
			} else if account, err := authSession(siteContext, r); err == nil {
				SetAccount(r, account)
			}
			h.ServeHTTP(w, r)
		})
	}
}

func authSession(context db.SiteContext, r *http.Request) (*models.Actor, error) {
	session, _ := context.Store.Get(r, "session")
	actorID, _ := session.Values["user-id"]
	if accountID, ok := actorID.(int64); ok {
		var account models.Actor
		if err := context.DB.Preload("User").First(&account, accountID).Error; err != nil {
			return nil, err
		}
		if !account.IsLocal(context.Settings.LocalDomains) {
			return &account, errIsRemote
		}
		return &account, nil
	}
	return nil, errNotLoggedIn
}

func authHTTPSSignatures(context db.SiteContext, r *http.Request) (*models.Actor, error) {
	verifier, err := httpsig.NewVerifier(r)
	if err != nil {
		return nil, err
	}
	pubKeyID := verifier.KeyId()
	actorID := strings.ReplaceAll(pubKeyID, "#main-actor", "")
	actor, err := dbmanagers.GetActorByURI(db.Context, actorID)
	if err != nil {
		return nil, err
	}
	if actor.IsLocal(context.Settings.LocalDomains) {
		return actor, errIsLocal
	}
	var algo httpsig.Algorithm = httpsig.Algorithm(httpsig.DigestSha256)
	pubKey, err := utils.ParsePublicKey([]byte(actor.GetPublicKey()))
	if err != nil {
		return nil, err
	}
	// The verifier will verify the Digest in addition to the HTTP signature
	return actor, verifier.Verify(pubKey, algo)
}
