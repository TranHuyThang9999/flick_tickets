package usecase

import (
	"errors"
	"flick_tickets/common/log"
	"flick_tickets/configs"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UseCaseJwt struct {
	config             *configs.Configs
	expAccessToken     time.Duration
	expRefreshToken    time.Duration
	userRepositoryPort domain.RepositoryUser
}

func NewUseCaseJwt(cf *configs.Configs, userRepositoryPort domain.RepositoryUser) (*UseCaseJwt, error) {
	expAccessToken, err := time.ParseDuration(cf.ExpireAccess)
	if err != nil {
		return nil, fmt.Errorf("expire access token has wrong format: %s", err)
	}
	expRefreshToken, err := time.ParseDuration(cf.ExpireRefresh)
	if err != nil {
		return nil, fmt.Errorf("expire refresh token has wrong format: %s", err)
	}
	return &UseCaseJwt{
		config:             cf,
		expAccessToken:     expAccessToken,
		expRefreshToken:    expRefreshToken,
		userRepositoryPort: userRepositoryPort,
	}, nil
}

func (u *UseCaseJwt) encrypt(secret string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (u *UseCaseJwt) Decrypt(tokenString string) (*entities.UserJwtClaim, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&entities.UserJwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(u.config.AccessSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*entities.UserJwtClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, nil
}
func (u *UseCaseJwt) generateToken(id int64, role int, userName string) (*entities.JwtToken, error) {
	userClaim := func(expire time.Duration) *entities.UserJwtClaim {
		return &entities.UserJwtClaim{
			Id:       id,
			UserName: userName,
			Role:     role,
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: time.Now().Add(expire).Unix(),
			},
		}
	}

	accessToken, err := u.encrypt(u.config.AccessSecret, userClaim(u.expAccessToken))
	if err != nil {
		log.Error(err, "Error when generating access token")
		return nil, err
	}
	refreshToken, err := u.encrypt(u.config.RefreshSecret, userClaim(u.expRefreshToken))
	if err != nil {
		log.Error(err, "Error when generating refresh token")
		return nil, err
	}

	return &entities.JwtToken{
		AccessToken:  accessToken,
		AtExpires:    int64(u.expAccessToken / time.Second),
		RefreshToken: refreshToken,
		RtExpires:    int64(u.expRefreshToken / time.Second),
	}, nil
}
