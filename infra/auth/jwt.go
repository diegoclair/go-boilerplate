package auth

import (
	"context"
	"strings"
	"time"

	"github.com/diegoclair/go_boilerplate/infra/logger"
	"github.com/diegoclair/go_utils-lib/v2/resterrors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/chacha20poly1305"
)

type jwtAuth struct {
	jwtPrivateKey string
	log           logger.Logger
}

func newJwtAuth(jwtPrivateKey string, log logger.Logger) (AuthToken, error) {
	if len(jwtPrivateKey) != chacha20poly1305.KeySize {
		return nil, errInvalidPrivateKey
	}

	return &jwtAuth{
		jwtPrivateKey: jwtPrivateKey,
		log:           log,
	}, nil
}

func (a *jwtAuth) CreateAccessToken(ctx context.Context, accountUUID, sessionUUID string) (tokenString string, payload *tokenPayload, err error) {
	return a.createToken(ctx, accountUUID, sessionUUID, accessTokenDurationTime)
}

func (a *jwtAuth) CreateRefreshToken(ctx context.Context, accountUUID, sessionUUID string) (tokenString string, payload *tokenPayload, err error) {
	return a.createToken(ctx, accountUUID, sessionUUID, refreshTokenDurationTime)
}

func (a *jwtAuth) VerifyToken(ctx context.Context, token string) (payload *tokenPayload, err error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, resterrors.NewUnauthorizedError(errInvalidToken.Error())
		}
		return []byte(a.jwtPrivateKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &tokenPayload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && strings.Contains(verr.Inner.Error(), errExpiredToken.Error()) {
			a.log.Errorf(ctx, "expired token: %v", err)
			return nil, resterrors.NewUnauthorizedError(errExpiredToken.Error())
		}
		return nil, resterrors.NewUnauthorizedError(errInvalidToken.Error(), err)
	}

	var ok bool
	payload, ok = jwtToken.Claims.(*tokenPayload)
	if !ok {
		a.log.Errorf(ctx, "could not parse jwt token: %v", err)
		return nil, resterrors.NewUnauthorizedError(errInvalidToken.Error())
	}
	return payload, nil
}

func (a *jwtAuth) createToken(ctx context.Context, accountUUID, sessionUUID string, duration time.Duration) (tokenString string, payload *tokenPayload, err error) {
	key := []byte(a.jwtPrivateKey)
	payload = newPayload(accountUUID, sessionUUID, duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err = token.SignedString(key)
	if err != nil {
		a.log.Errorf(ctx, "error to encrypt token: %v", err)
		return tokenString, payload, resterrors.NewUnauthorizedError(err.Error())
	}

	return tokenString, payload, nil
}
