package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/config"
)

type Auth struct {
	cfg *config.Config
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func New(cfg *config.Config) *Auth {
	return &Auth{
		cfg: cfg,
	}
}

func (a *Auth) GenerateToken(username, password string) (string, error) {
	claims := Claims{
		Username: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *Auth) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func JWTMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	config := jwtConfig(cfg)
	return echojwt.WithConfig(config)
}

func jwtConfig(cfg *config.Config) echojwt.Config {
	return echojwt.Config{
		SigningKey: []byte(cfg.JWTSecret),
		ContextKey: "token",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"errors": "Ошбика авторизации",
			})
		},
	}
}
