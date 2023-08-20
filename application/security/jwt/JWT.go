package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

const (
	Access  = "access"
	Refresh = "refresh"
)

type AppToken struct {
	Token          string    `json:"token"`
	TokenType      string    `json:"type"`
	ExpirationTime time.Time `json:"expirationTime"`
}

type Claims struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

var TokenTypeKeyName = map[string]string{
	Access:  "Secure:JWTAccessSecure",
	Refresh: "Secure.JWTRefreshSecure",
}

var TokenTypeExpTime = map[string]string{
	Access:  "Secure.JWTAccessTimeMinute",
	Refresh: "Secure.JWTRefreshTimeHour",
}

func GenerateJWTToken(userId, tokenType string) (appToken *AppToken, err error) {
	JWTSecureKey := viper.GetString(TokenTypeKeyName[tokenType])
	JWTExpTime := viper.GetString(TokenTypeExpTime[tokenType])

	tokenTimeConverted, err := strconv.ParseInt(JWTExpTime, 10, 64)
	if err != nil {
		return
	}

	tokenTimeUnix := time.Duration(tokenTimeConverted)
	switch tokenType {
	case Refresh:
		tokenTimeUnix *= time.Hour
	case Access:
		tokenTimeUnix *= time.Minute
	default:
		err = errors.New("invalid token type")
	}

	if err != nil {
		return
	}

	nowTime := time.Now()
	expirationTokenTime := nowTime.Add(tokenTimeUnix)

	tokenClaims := &Claims{
		ID:   userId,
		Type: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTokenTime),
		},
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenStr, err := tokenWithClaims.SignedString([]byte(JWTSecureKey))
	if err != nil {
		return
	}

	appToken = &AppToken{
		Token:          tokenStr,
		TokenType:      tokenType,
		ExpirationTime: expirationTokenTime,
	}

	return
}

func GetClaimsAndVerifyToken(tokenString, tokenType string) (claims jwt.MapClaims, err error) {
	JWTRefreshSecure := viper.GetString(TokenTypeKeyName[tokenType])
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWTRefreshSecure), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["type"] != tokenType {
			return nil, fmt.Errorf("invalid token type")
		}

		timeExpire := claims["exp"].(float64)
		if time.Now().Unix() > int64(timeExpire) {
			return nil, fmt.Errorf("token expired")
		}

		return claims, nil
	}

	return nil, err
}

func GetJWTConfig() echojwt.Config {
	accessKey := viper.GetString(TokenTypeKeyName[Access])

	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(Claims)
		},

		SigningKey: []byte(accessKey),
	}
}
