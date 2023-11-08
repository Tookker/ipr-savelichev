package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"ipr-savelichev/internal/config"
	"ipr-savelichev/internal/router/returnvalue"
)

var (
	ErrParse        = errors.New("Ошибка получения срока жизни JWT токена!")
	ErrDecodeToken  = errors.New("Ошибка расшифрования JWT токена!")
	ErrParseToken   = errors.New("Ошибка парсинга JWT токена!")
	ErrInvalidToken = errors.New("JWT токен не действителен!")
	ErrNegativeVal  = errors.New("Получено отрицательное число срока жизни JWT токена!")
	ErrHeader       = errors.New("Ошибка заголовка")
)

var (
	keyJWT = []byte("ЧеловекУкусилПчелуИОнаРазбухла")
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name=JWT
type JWT interface {
	GenerateJWT(login string) (string, error)
	CheckJWT(token string) error
	MiddlewareJWT(next http.Handler) http.Handler
}

type JWTController struct {
	tokenLifeTime int
	logger        *zap.Logger
}

type tokenClaims struct {
	jwt.RegisteredClaims
}

func NewJWTController(config *config.Config, logger *zap.Logger) (JWT, error) {
	res, err := strconv.Atoi(config.JWTLifeTime)
	if err != nil {
		logger.Error(ErrParse.Error())
		return nil, ErrParse
	}

	if res < 0 {
		logger.Error(ErrNegativeVal.Error())
		return nil, ErrNegativeVal
	}

	return &JWTController{
		tokenLifeTime: res,
		logger:        logger,
	}, nil
}

// Проверка на валидность токена
func (j *JWTController) CheckJWT(token string) error {
	res, err := jwt.ParseWithClaims(token,
		&tokenClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrDecodeToken
			}

			return keyJWT, nil
		})

	if err != nil {
		j.logger.Error(ErrDecodeToken.Error())
		return err
	}

	claims, ok := res.Claims.(*tokenClaims)
	if !ok {
		j.logger.Error(ErrParseToken.Error())
		return ErrParseToken
	}

	if !time.Now().Before(claims.ExpiresAt.Time) {
		j.logger.Error(ErrInvalidToken.Error())
		return ErrInvalidToken
	}

	j.logger.Debug("Проверка токена прошла успешно!")
	return nil
}

// Сгенерировать jwt токен
func (j *JWTController) GenerateJWT(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			Issuer:    login,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.tokenLifeTime) * time.Minute)),
		},
	})

	tokenString, err := token.SignedString(keyJWT)
	if err != nil {
		j.logger.Error(err.Error())
		return "", err
	}

	return tokenString, nil
}

func (j *JWTController) MiddlewareJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header["Authorization"]
		if !ok {
			j.logger.Error(ErrHeader.Error())
			returnvalue.NewResponse(w).ResponseErr(http.StatusForbidden, ErrHeader.Error())
			return
		}

		const tokenSize = 1
		if len(token) != tokenSize {
			j.logger.Error(ErrHeader.Error())
			returnvalue.NewResponse(w).ResponseErr(http.StatusForbidden, ErrHeader.Error())
			return
		}

		err := j.CheckJWT(strings.Split(token[0], " ")[1])
		fmt.Println(token, len(token), strings.Split(token[0], " ")[1])
		if err != nil {
			j.logger.Error(err.Error())
			returnvalue.NewResponse(w).ResponseErr(http.StatusUnauthorized, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
