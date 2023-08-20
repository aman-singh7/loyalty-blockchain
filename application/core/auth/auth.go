package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/aman-singh7/loyalty-blockchain/application/security/jwt"
	userRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

type Auth struct {
	AccessToken               string    `json:"access_token"`
	RefreshToken              string    `json:"refresh_token"`
	ExpirationAccessDateTime  time.Time `json:"expirationAccessDateTime"`
	ExpirationRefreshDateTime time.Time `json:"expirationRefreshDateTime"`
}

type Service struct {
	repo *userRepo.Repository
	app  *firebase.App
}

func NewService(repo *userRepo.Repository) *Service {
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	return &Service{
		repo: repo,
		app:  app,
	}
}

func (s *Service) UserLogin(uid int, token string) (*Auth, error) {
	ctx := context.Background()
	client, err := s.app.Auth(ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	_, err = client.VerifyIDTokenAndCheckRevoked(ctx, token)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: user repo
	// userId, err := s.repo.UserExistsWithUid(uid)
	userId := "434"
	if userId == "" {
		// TODO: create user against this id
		// userId, err := s.repo.CreateUserWithUid(uid)
	}
	accessToken, err := jwt.GenerateJWTToken(userId, "access")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	refreshtoken, err := jwt.GenerateJWTToken(userId, "refresh")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return &Auth{
		AccessToken:               accessToken.Token,
		RefreshToken:              refreshtoken.Token,
		ExpirationAccessDateTime:  accessToken.ExpirationTime,
		ExpirationRefreshDateTime: refreshtoken.ExpirationTime,
	}, nil
}

func (s *Service) GenerateAccessTokenFromRefreshToken(refreshToken string) (*Auth, error) {
	claims, err := jwt.GetClaimsAndVerifyToken(refreshToken, "refresh")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: check if user exists
	userId := claims["uid"].(string)
	// userId, err := s.repo.UserExistsWithUid(claims["uid"].(string))
	// if err != nil {
	// 	return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	// }
	accessToken, err := jwt.GenerateJWTToken(userId, "access")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	refreshtoken, err := jwt.GenerateJWTToken(userId, "refresh")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return &Auth{
		AccessToken:               accessToken.Token,
		RefreshToken:              refreshtoken.Token,
		ExpirationAccessDateTime:  accessToken.ExpirationTime,
		ExpirationRefreshDateTime: refreshtoken.ExpirationTime,
	}, nil
}
