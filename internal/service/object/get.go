package object

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *objectService) GetByTemplateID(ctx context.Context, templateID int64) ([]entity.Object, error) {
	return svc.ObjectRepository.GetByTemplateID(ctx, templateID)
}

func (svc *objectService) GetByIDs(ctx context.Context, ids []int64) ([]entity.Object, error) {
	return svc.ObjectRepository.GetByIDs(ctx, ids)
}

func (svc *objectService) GetByID(ctx context.Context, id int64) (*entity.Object, error) {
	return svc.ObjectRepository.GetByID(ctx, id)
}

func (svc *objectService) GetByNamesAndTemplateID(ctx context.Context, names []string, templateID int64) ([]entity.Object, error) {
	return svc.ObjectRepository.GetByTemplateIDAndNames(ctx, names, templateID)
}

func (svc *objectService) GetExByIDs(ctx context.Context, ids []int64) ([]entity.ObjectEx, error) {
	var (
		locationIDByObjectID = make(map[int64]int64)
		locationByID         = make(map[int64]entity.Location)
		objectByID           = make(map[int64]entity.Object)
		locationIDs          []int64
		result               []entity.ObjectEx
	)

	objects, err := svc.ObjectRepository.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	for _, obj := range objects {
		if obj.LocationID != nil {
			locationIDByObjectID[obj.ID] = *obj.LocationID
			locationIDs = append(locationIDs, *obj.LocationID)
		}

		objectByID[obj.ID] = obj
	}

	locations, err := svc.LocationRepository.GetByIDs(ctx, locationIDs)
	if err != nil {
		return nil, err
	}

	for _, location := range locations {
		locationByID[location.ID] = location
	}

	for _, obj := range objects {
		var (
			objectLocation *entity.Location
		)

		if obj.LocationID != nil {
			location, ok := locationByID[*obj.LocationID]
			if ok {
				objectLocation = &location
			}
		}

		attributeValues, err := svc.ObjectAttributeService.GetAttributesExByObjectID(ctx, obj.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, *entity.NewObjectExPtr(&obj, objectLocation, attributeValues))
	}

	return result, nil
}

func (svc *objectService) GetExByTemplateID(ctx context.Context, templateID int64) ([]entity.ObjectEx, error) {
	var (
		objectIDs []int64
	)

	objects, err := svc.ObjectRepository.GetByTemplateID(ctx, templateID)
	if err != nil {
		return nil, err
	}

	for _, obj := range objects {
		objectIDs = append(objectIDs, obj.ID)
	}

	return svc.GetExByIDs(ctx, objectIDs)
}
