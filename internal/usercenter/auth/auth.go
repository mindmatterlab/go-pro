package auth

//go:generate mockgen -self_package github.com/mindmatterlab/go-pro/internal/usercenter/auth -destination mock_auth.go -package auth github.com/mindmatterlab/go-pro/internal/usercenter/auth AuthProvider,AuthzInterface,AuthnInterface

import (
	"context"

	"github.com/google/wire"

	"github.com/mindmatterlab/go-pro/pkg/authn"
)

// ProviderSet is a Wire provider set that creates a new instance of auth.
var ProviderSet = wire.NewSet(NewAuth, wire.Bind(new(AuthProvider), new(*auth)), AuthnProviderSet, AuthzProviderSet)

// AuthProvider is an interface that combines both the AuthnInterface and AuthzInterface interfaces.
type AuthProvider interface {
	AuthnInterface
	AuthzInterface
}

// auth is a struct that implements AuthnInterface and AuthzInterface interfaces.
type auth struct {
	authn AuthnInterface
	authz AuthzInterface
}

// NewAuth is a constructor function that creates a new instance of auth struct.
func NewAuth(authn AuthnInterface, authz AuthzInterface) *auth {
	return &auth{authn: authn, authz: authz}
}

// Verify is a method that implements Verify method of AuthnInterface.
func (a *auth) Verify(accessToken string) (string, error) {
	return a.authn.Verify(accessToken)
}

// Sign is a method that implements Sign method of AuthnInterface.
func (a *auth) Sign(ctx context.Context, userID string) (authn.IToken, error) {
	return a.authn.Sign(ctx, userID)
}

// Authorize is a method that implements Authorize method of AuthzInterface.
func (a *auth) Authorize(rvals ...any) (bool, error) {
	return a.authz.Authorize(rvals...)
}