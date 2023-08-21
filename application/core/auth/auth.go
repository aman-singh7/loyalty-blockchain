package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/aman-singh7/loyalty-blockchain/application/security/jwt"
	domain "github.com/aman-singh7/loyalty-blockchain/domain/auth"
	"github.com/aman-singh7/loyalty-blockchain/domain/user"
	userRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

type Auth struct {
	AccessToken     string    `json:"accessToken"`
	RefreshToken    string    `json:"refreshToken"`
	ExpAccessToken  time.Time `json:"expAccessToken"`
	ExpRefreshToken time.Time `json:"expRefreshToken"`
}

type Service struct {
	repo *userRepo.Repository
	app  *firebase.App
}

func NewService(repo *userRepo.Repository) *Service {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	return &Service{
		repo: repo,
		app:  app,
	}
}

func (s *Service) CreateUser(request *domain.CreateUserRequest) (*AuthenticatedUser, error) {
	ctx := context.Background()
	client, err := s.app.Auth(ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	_, err = client.VerifyIDTokenAndCheckRevoked(ctx, request.Token)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	userId, _ := s.repo.UserExists(request.UID)
	var domainUser *user.User
	if userId == "" {
		user, err := client.GetUser(ctx, request.UID)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
		}

		dUser := newFirebaseUserMapper(user)
		domainUser, err = s.repo.Create(dUser)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
		}

		userId = domainUser.UID
	} else {
		domainUser, err = s.repo.GetByPUID(request.UID)
		if err != nil {
			return nil, fmt.Errorf("[CreateUser] Error: %v", err)
		}
	}

	accessToken, err := jwt.GenerateJWTToken(userId, "access")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	refreshtoken, err := jwt.GenerateJWTToken(userId, "refresh")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return authUserMapper(domainUser, &Auth{
		AccessToken:     accessToken.Token,
		RefreshToken:    refreshtoken.Token,
		ExpAccessToken:  accessToken.ExpirationTime,
		ExpRefreshToken: refreshtoken.ExpirationTime,
	}), nil
}

func (s *Service) GenerateAccessTokenFromRefreshToken(request *domain.GenerateTokenRequest) (*Auth, error) {
	claims, err := jwt.GetClaimsAndVerifyToken(request.Token, "refresh")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// TODO: check if user exists
	userId := claims["uid"].(string)
	_, err = s.repo.UserExists(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": err.Error()})
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
		AccessToken:     accessToken.Token,
		RefreshToken:    refreshtoken.Token,
		ExpAccessToken:  accessToken.ExpirationTime,
		ExpRefreshToken: refreshtoken.ExpirationTime,
	}, nil
}
