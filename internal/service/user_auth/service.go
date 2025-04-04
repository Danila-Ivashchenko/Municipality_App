package user_auth

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"time"

	"crypto/rand"
)

const (
	tokenSize                 = 128
	triesToGenerateTokenCount = 5
)

func generateToken(size int) (string, error) {
	byteSize := size / 2

	bytes := make([]byte, byteSize)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(bytes)

	return token, nil
}

func (svc userAuthService) NewUserToken(ctx context.Context, userID int64) (*entity.UserAuthToken, error) {
	token, err := generateToken(tokenSize)
	if err != nil {
		return nil, err
	}

	tokenExists, err := svc.tokenExists(ctx, token)
	if err != nil {
		return nil, err
	}

	for tries := 0; tokenExists; tries++ {
		slog.WarnContext(ctx, "Token generation warn: token exists", slog.Any("user_id", userID))

		tokenExists, err = svc.tokenExists(ctx, token)
		if err != nil {
			return nil, err
		}

		if tries == triesToGenerateTokenCount {
			return nil, errors.New("token generation warn: too many tries")
		}
	}

	expireAt := time.Now().Add(entity.DefaultExpireTime)

	repoData := &repository.CreateUserTokenData{
		Token:    token,
		UserID:   userID,
		ExpireAt: expireAt,
	}

	fmt.Println("token size", len(token))

	return svc.UserAuthTokenRepository.CreateUserToken(ctx, repoData)
}

func (svc userAuthService) tokenExists(ctx context.Context, token string) (bool, error) {
	userToken, err := svc.UserAuthTokenRepository.GetByToken(ctx, token)
	if err != nil {
		return false, err
	}

	return userToken != nil, nil
}

func (svc userAuthService) DeleteUserToken(ctx context.Context, userToken *entity.UserAuthToken) error {
	if userToken == nil {
		return nil
	}

	return svc.UserAuthTokenRepository.DeleteUserTokenByID(ctx, userToken.ID)
}

func (svc userAuthService) DeleteUserTokenByID(ctx context.Context, id int64) error {
	return svc.UserAuthTokenRepository.DeleteUserTokenByID(ctx, id)
}

func (svc userAuthService) DeleteAllUserTokens(ctx context.Context, userID string) error {
	//TODO implement me
	panic("implement me")
}

func (svc userAuthService) UpdateExpireTime(ctx context.Context, userID string, timeDelta time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (svc userAuthService) GetByTokenWithValidation(ctx context.Context, token string) (*entity.UserAuthToken, error) {
	userToken, err := svc.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if userToken == nil {
		return nil, errors.New("token not found")
	}

	if userToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	return userToken, nil
}

func (svc userAuthService) GetByToken(ctx context.Context, token string) (*entity.UserAuthToken, error) {
	return svc.UserAuthTokenRepository.GetByToken(ctx, token)
}

func (svc userAuthService) GetByUserID(ctx context.Context, userID int64) ([]entity.UserAuthToken, error) {
	return svc.UserAuthTokenRepository.GetByUserID(ctx, userID)
}

func (svc userAuthService) DeleteExtraTokens(ctx context.Context, userID int64) error {
	var (
		extraTokens []entity.UserAuthToken
	)

	existTokens, err := svc.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if len(existTokens) > entity.UserTokensCount-1 {
		for i := 0; i < len(existTokens)-(entity.UserTokensCount-1); i++ {
			extraTokens = append(extraTokens, existTokens[i])
		}
	}

	for _, extraToken := range extraTokens {
		err = svc.DeleteUserTokenByID(ctx, extraToken.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
