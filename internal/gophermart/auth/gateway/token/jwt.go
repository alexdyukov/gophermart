package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	jwt "github.com/golang-jwt/jwt/v4"
)

type (
	AuthJWTGateway struct {
		secret   []byte
		lifetime int64
	}

	ClaimsFields map[string]string

	customClaims struct {
		ClaimsFields `json:"claims"`
		jwt.RegisteredClaims
	}
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrCustomClaims = errors.New("assertion to CustomClaims failed")
)

func NewAuthJWTGateway(lifetime int64, secret []byte) *AuthJWTGateway {
	return &AuthJWTGateway{
		lifetime: lifetime,
		secret:   secret,
	}
}

func (tkn *AuthJWTGateway) IssueWithLoginAndID(login, id string) (string, error) {
	cl := make(ClaimsFields)
	cl["ID"] = id
	cl["Login"] = login

	jwtSigned, err := tkn.generate(cl)
	if err != nil {
		return "", err
	}

	return jwtSigned, nil
}

func (tkn *AuthJWTGateway) ValidateWithLoginAndID(jwtString string) (*sharedkernel.User, error) {
	claims, err := tkn.validate(jwtString)
	if err != nil {
		return nil, err
	}

	if login, ok := claims["Login"]; ok {
		if id, ok := claims["ID"]; ok {
			return sharedkernel.RestoreUser(id, login), nil
		}
	}

	return nil, ErrCustomClaims
}

func (tkn *AuthJWTGateway) newCustomClaims(claimsDto ClaimsFields) customClaims {
	expUnix := time.Now().Unix() + tkn.lifetime
	expTime := time.Unix(expUnix, 0)

	return customClaims{
		ClaimsFields: claimsDto,
		RegisteredClaims: jwt.RegisteredClaims{ // nolint:exhaustivestruct // ok
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
}

func (tkn *AuthJWTGateway) generate(claims ClaimsFields) (string, error) {
	cc := tkn.newCustomClaims(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)

	signedToken, err := token.SignedString(tkn.secret)
	if err != nil {
		return "", err // nolint:wrapcheck // ok
	}

	return signedToken, nil
}

func (tkn *AuthJWTGateway) validate(jwtString string) (ClaimsFields, error) {
	token, err := jwt.ParseWithClaims(
		jwtString,
		new(customClaims),
		func(token *jwt.Token) (interface{}, error) {
			return tkn.secret, nil
		},
		jwt.WithValidMethods([]string{
			"HS256",
		}),
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token is expired: %w", err)
		}

		return nil, ErrInvalidToken
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	cc, ok := token.Claims.(*customClaims)
	if !ok {
		return nil, ErrCustomClaims
	}

	return cc.ClaimsFields, nil
}
