package usecase

import (
	"context"
	"errors"
	"flick_tickets/common/enums"
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/configs"
	"flick_tickets/core/domain"
	"flick_tickets/core/entities"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtUseCase struct {
	config             *configs.Configs
	expAccessToken     time.Duration
	expRefreshToken    time.Duration
	userRepositoryPort domain.RepositoryUser
}

func NewJwtUseCase(cf *configs.Configs, userRepositoryPort domain.RepositoryUser) (*JwtUseCase, error) {
	expAccessToken, err := time.ParseDuration(cf.ExpireAccess)
	if err != nil {
		return nil, fmt.Errorf("expire access token has wrong format: %s", err)
	}
	expRefreshToken, err := time.ParseDuration(cf.ExpireRefresh)
	if err != nil {
		return nil, fmt.Errorf("expire refresh token has wrong format: %s", err)
	}
	return &JwtUseCase{
		config:             cf,
		expAccessToken:     expAccessToken,
		expRefreshToken:    expRefreshToken,
		userRepositoryPort: userRepositoryPort,
	}, nil
}

func (u *JwtUseCase) encrypt(secret string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (u *JwtUseCase) Decrypt(tokenString string) (*entities.UserJwtClaim, error) {
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
func (u *JwtUseCase) generateToken(id int64, userName string) (*entities.JwtToken, error) {
	userClaim := func(expire time.Duration) *entities.UserJwtClaim {
		return &entities.UserJwtClaim{
			Id:       id,
			UserName: userName,
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
func (u *JwtUseCase) LoginUser(ctx context.Context, user_name string, password string) (*entities.ResponseLogin, error) {
	log.Infof("req : ", user_name, password)
	users, err := u.userRepositoryPort.GetAllUserStaffs(ctx, &domain.UsersReqByForm{
		UserName: user_name,
	})
	if err != nil {
		return &entities.ResponseLogin{
			Result: entities.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	log.Infof("user : ", len(users))
	if users == nil {
		return &entities.ResponseLogin{
			Result: entities.Result{
				Code:    enums.USER_NOT_EXIST_CODE,
				Message: enums.USER_NOT_EXIST_MESS,
			},
		}, nil
	}

	err = utils.ComparePassword(users[0].Password, password)
	if err != nil {
		return &entities.ResponseLogin{
			Result: entities.Result{
				Code:    enums.USER_NOT_EXIST_CODE,
				Message: enums.USER_NOT_EXIST_MESS,
			},
		}, nil
	}
	log.Error(err, "error passwors")
	token, err := u.generateToken(users[0].Id, users[0].UserName)
	if err != nil {
		return &entities.ResponseLogin{
			Result: entities.Result{
				Code:    enums.CREATE_TOKEN,
				Message: enums.CREATE_TOKEN_MESS,
			},
		}, nil
	}

	return &entities.ResponseLogin{
		Result: entities.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		JwtToken: token,
	}, nil
}
