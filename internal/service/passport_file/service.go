package passport_file

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"municipality_app/internal/domain/entity"
	"time"
)

func (svc *passportFileService) Create(ctx context.Context, municipality *entity.Municipality, passport *entity.PassportEx) (*entity.PassportFile, error) {
	var (
		err error
	)

	fileName := fmt.Sprintf("Паспорт туризма муниципального образования %s %s.pdf", municipality.Name, passport.Name)
	storagePass := "storage"
	uniquePath := uuid.New().String() + ".pdf"
	filePath := fmt.Sprintf("%s/%s", storagePass, uniquePath)

	err = svc.BuildPassportFile(filePath, passport)
	if err != nil {
		return nil, err
	}

	return svc.savePassportFile(ctx, fileName, filePath, passport.ID)
}

func (svc *passportFileService) GetByPassportID(ctx context.Context, passportID int64) ([]entity.PassportFile, error) {
	return svc.PassportFileRepository.GetByPassportID(ctx, passportID)
}

func (svc *passportFileService) GetLastByPassportID(ctx context.Context, passportID int64) (*entity.PassportFile, error) {
	passPortFile, err := svc.PassportFileRepository.GetLastByPassportID(ctx, passportID)
	if err != nil {
		return nil, err
	}

	if passPortFile != nil {
		storageBaseURL := svc.Config.GetFileStorageBaseURL()

		fullPath := fmt.Sprintf("%s/%s", storageBaseURL, passPortFile.Path)
		passPortFile.Path = fullPath
	}

	return passPortFile, nil
}

func (svc *passportFileService) savePassportFile(ctx context.Context, fileName, path string, passportID int64) (*entity.PassportFile, error) {
	newPassportFile := &entity.PassportFile{
		Path:       path,
		PassportID: passportID,
		FileName:   fileName,

		CreateAt: time.Now().UTC(),
	}

	passportFile, err := svc.PassportFileRepository.Create(ctx, newPassportFile)
	if err != nil {
		return nil, err
	}

	if passportFile != nil {
		storageBaseURL := svc.Config.GetFileStorageBaseURL()

		fullPath := fmt.Sprintf("%s/%s", storageBaseURL, passportFile.Path)
		passportFile.Path = fullPath
	}

	return passportFile, nil
}
