package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	jwtClaimsFieldExp string = "exp"
	jwtClaimsFieldIat string = "iat"
	defaultSecretKey  string = "abc!@#$%^&*1234567890"
)

type (
	JWTFilter struct {
		opts jwtFilterOptions
	}

	jwtFilterOptions struct {
		secretKey []byte
	}

	JWTFilterOptions interface {
		apply(*jwtFilterOptions)
	}
)

var defaultJwtFilterOptions = jwtFilterOptions{
	secretKey: []byte(defaultSecretKey),
}

type funcJWTFilterOption struct {
	f func(*jwtFilterOptions)
}

func newFuncJWTFilterOption(f func(*jwtFilterOptions)) *funcJWTFilterOption {
	return &funcJWTFilterOption{
		f: f,
	}
}

func (fdo *funcJWTFilterOption) apply(do *jwtFilterOptions) {
	fdo.f(do)
}

func WithSecretKey(secretKey []byte) JWTFilterOptions {
	return newFuncJWTFilterOption(func(o *jwtFilterOptions) {
		o.secretKey = secretKey
	})
}

func NewJWTFilter(opts ...JWTFilterOptions) *JWTFilter {
	jwtFilter := &JWTFilter{
		opts: defaultJwtFilterOptions,
	}

	for _, opt := range opts {
		opt.apply(&jwtFilter.opts)
	}

	return jwtFilter
}

func (j *JWTFilter) GenerateToken(claims jwt.MapClaims) (string, error) {
	if _, ok := claims[jwtClaimsFieldExp]; !ok {
		claims[jwtClaimsFieldExp] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	}

	claims[jwtClaimsFieldIat] = time.Now().Unix()
	return j.renewToken(claims)
}

func (j *JWTFilter) renewToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.opts.secretKey)
}

func (j *JWTFilter) ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.opts.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("error jwt token")
}
