package context_paylod_parser

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func SetUserToContext(ctx context.Context, user *entity.User) context.Context {
	return context.WithValue(ctx, UserContextPayloadKey, user)
}

func SetMunicipalityToContext(ctx context.Context, municipality *entity.Municipality) context.Context {
	return context.WithValue(ctx, MunicipalityContextPayloadKey, municipality)
}

func SetPassportToContext(ctx context.Context, passport *entity.Passport) context.Context {
	return context.WithValue(ctx, PassportContextPayloadKey, passport)
}

func SetChapterToContext(ctx context.Context, chapter *entity.Chapter) context.Context {
	return context.WithValue(ctx, ChapterContextPayloadKey, chapter)
}

func SetPartitionToContext(ctx context.Context, partition *entity.Partition) context.Context {
	return context.WithValue(ctx, PartitionContextPayloadKey, partition)
}

func SetObjectTemplateToContext(ctx context.Context, objectTemplate *entity.ObjectTemplate) context.Context {
	return context.WithValue(ctx, ObjectTemplatePayloadKey, objectTemplate)
}

func SetUserAuthTokenToContext(ctx context.Context, userAuthToken *entity.UserAuthToken) context.Context {
	return context.WithValue(ctx, UserAuthTokenPayloadKey, userAuthToken)
}

func GetUserFromContext(ctx context.Context) *entity.User {
	userValue, ok := ctx.Value(UserContextPayloadKey).(*entity.User)
	if !ok {
		return nil
	}

	return userValue
}

func GetMunicipalityFromContext(ctx context.Context) *entity.Municipality {
	municipalityValue, ok := ctx.Value(MunicipalityContextPayloadKey).(*entity.Municipality)
	if !ok {
		return nil
	}

	return municipalityValue
}

func GetPassportFromContext(ctx context.Context) *entity.Passport {
	passportValue, ok := ctx.Value(PassportContextPayloadKey).(*entity.Passport)
	if !ok {
		return nil
	}

	return passportValue
}

func GetPartitionFromContext(ctx context.Context) *entity.Partition {
	partitionValue, ok := ctx.Value(PartitionContextPayloadKey).(*entity.Partition)
	if !ok {
		return nil
	}

	return partitionValue
}

func GetChapterFromContext(ctx context.Context) *entity.Chapter {
	chapterValue, ok := ctx.Value(ChapterContextPayloadKey).(*entity.Chapter)
	if !ok {
		return nil
	}

	return chapterValue
}

func GetObjectTemplateFromContext(ctx context.Context) *entity.ObjectTemplate {
	objectTemplateValue, ok := ctx.Value(ObjectTemplatePayloadKey).(*entity.ObjectTemplate)
	if !ok {
		return nil
	}

	return objectTemplateValue
}

func GetUserAuthTokenFromContext(ctx context.Context) *entity.UserAuthToken {
	userAuthTokenValue, ok := ctx.Value(UserAuthTokenPayloadKey).(*entity.UserAuthToken)
	if !ok {
		return nil
	}

	return userAuthTokenValue
}
