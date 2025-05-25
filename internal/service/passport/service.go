package passport

import (
	"context"
	"errors"
	"github.com/thanhpk/randstr"
	"municipality_app/internal/domain/core_errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

const (
	generateRevisionCodeTries = 4
)

func generateRevisionCode() string {
	return randstr.String(11)
}

func (svc *passportService) Create(ctx context.Context, data *service.CreatePassportData) (*entity.Passport, error) {
	var (
		result *entity.Passport
	)

	municipalityExists, err := svc.MunicipalityRepository.GetById(ctx, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if municipalityExists == nil {
		return nil, core_errors.MunicipalityNotFound
	}

	passportExists, err := svc.PassportRepository.GetByNameAndMunicipalityID(ctx, data.Name, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if passportExists != nil {
		return nil, core_errors.PassportNameAlreadyUsed
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		revisionCode, err := svc.getNewRevisionCode(tx)
		if err != nil {
			return err
		}

		if data.IsMain {
			mainPassport, err := svc.PassportRepository.GetMainByMunicipalityID(tx, data.MunicipalityID)
			if err != nil {
				return err
			}

			if mainPassport != nil {
				err = svc.PassportRepository.ChangeIsMainByID(tx, mainPassport.ID, false)
				if err != nil {
					return err
				}
			}
		}

		repoData := &repository.CreatePassportData{
			Name:           data.Name,
			MunicipalityID: data.MunicipalityID,
			Description:    data.Description,
			Year:           data.Year,
			IsMain:         data.IsMain,
			RevisionCode:   revisionCode,
		}

		result, err = svc.PassportRepository.Create(tx, repoData)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc *passportService) getNewRevisionCode(ctx context.Context) (string, error) {
	var (
		triesUsed = 0
	)

	revisionCode := generateRevisionCode()

	isUsed, err := svc.checkRevisionCodeIsUsed(ctx, revisionCode)
	if err != nil {
		return "", err
	}

	for isUsed && triesUsed < generateRevisionCodeTries {
		isUsed, err = svc.checkRevisionCodeIsUsed(ctx, revisionCode)
		if err != nil {
			return "", err
		}
		triesUsed++
	}

	if triesUsed == generateRevisionCodeTries {
		return "", errors.New("revision code generation has been reached")
	}

	return revisionCode, nil
}

func (svc *passportService) checkRevisionCodeIsUsed(ctx context.Context, code string) (bool, error) {
	passport, err := svc.GetByRevisionCode(ctx, code)
	if err != nil {
		return false, err
	}

	if passport == nil {
		return false, nil
	}

	return true, nil
}

func isUpdate(updated bool, firstParam, secondParam any) bool {
	return updated || firstParam != secondParam
}

func (svc *passportService) Update(ctx context.Context, data *service.UpdatePassportData) (*entity.Passport, error) {
	var (
		updated bool
	)

	passport, err := svc.GetByIDAndMunicipalityID(ctx, data.ID, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if passport == nil {
		return nil, errors.New("passport not found")
	}

	if data.Name != nil {
		updated = isUpdate(updated, passport.Name, *data.Name)
		passport.Name = *data.Name
	}

	if data.Description != nil {
		updated = isUpdate(updated, passport.Description, *data.Description)
		passport.Description = *data.Description
	}

	if data.Year != nil {
		updated = isUpdate(updated, passport.Year, *data.Year)
		passport.Year = *data.Year
	}

	if data.IsHidden != nil {
		updated = isUpdate(updated, passport.IsHidden, *data.IsHidden)
		passport.IsHidden = *data.IsHidden
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		if data.IsMain != nil && *data.IsMain {
			err = svc.MakeMainPassportToMunicipality(tx, data.ID, data.MunicipalityID)
			if err != nil {
				return err
			}
		}

		if !updated {
			return nil
		}

		err = svc.PassportRepository.Update(tx, passport)
		if err != nil {
			return err
		}

		return nil
	})

	return passport, nil
}

func (svc *passportService) UpdatedAt(ctx context.Context, passportID int64) error {
	return svc.PassportRepository.UpdateUpdatedAt(ctx, passportID)
}

func (svc *passportService) Delete(ctx context.Context, id, municipalityID int64) error {
	passport, err := svc.GetByIDAndMunicipalityID(ctx, id, municipalityID)
	if err != nil {
		return err
	}

	if passport == nil {
		return errors.New("passport not found")
	}

	if passport.IsMain {
		return errors.New("passport is main")
	}

	return svc.PassportRepository.Delete(ctx, id)
}

func (svc *passportService) MakeMainPassportToMunicipality(ctx context.Context, id, municipalityID int64) error {
	err := svc.Transactor.Execute(ctx, func(tx context.Context) error {
		mainPassport, err := svc.PassportRepository.GetMainByMunicipalityID(ctx, municipalityID)
		if err != nil {
			return err
		}

		if mainPassport != nil {
			err = svc.PassportRepository.ChangeIsMainByID(ctx, mainPassport.ID, false)
			if err != nil {
				return err
			}
		}

		passport, err := svc.GetByIDAndMunicipalityID(ctx, id, municipalityID)
		if err != nil {
			return err
		}

		if passport == nil {
			return errors.New("passport not found")
		}

		return svc.PassportRepository.ChangeIsMainByID(ctx, passport.ID, true)
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc *passportService) GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.Passport, error) {
	return svc.PassportRepository.GetByIDAndMunicipalityID(ctx, id, municipalityID)
}

func (svc *passportService) GetByIDsAndMunicipalityID(ctx context.Context, ids []int64, municipalityID int64) ([]entity.Passport, error) {
	return svc.PassportRepository.GetByIDsAndMunicipalityID(ctx, ids, municipalityID)
}

func (svc *passportService) GetByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.Passport, error) {
	return svc.PassportRepository.GetByMunicipalityID(ctx, municipalityID)
}

func (svc *passportService) GetMainByMunicipalityID(ctx context.Context, municipalityID int64) (*entity.Passport, error) {
	return svc.PassportRepository.GetMainByMunicipalityID(ctx, municipalityID)
}

func (svc *passportService) GetByRevisionCode(ctx context.Context, revisionCode string) (*entity.Passport, error) {
	return svc.PassportRepository.GetByRevisionCode(ctx, revisionCode)
}
