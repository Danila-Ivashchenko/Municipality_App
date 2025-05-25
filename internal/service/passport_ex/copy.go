package passport_ex

import (
	"context"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *passportExService) Copy(ctx context.Context, data *service.CopyData) (*entity.Passport, error) {
	var (
		newPassport *entity.Passport
		err         error
	)

	passportExExists, err := svc.GetByIDAndMunicipalityID(ctx, data.SrcID, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if passportExExists == nil {
		return nil, nil
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		createPassportData := &service.CreatePassportData{
			Name:           data.NewName,
			MunicipalityID: passportExExists.MunicipalityID,
			Description:    passportExExists.Description,
			Year:           data.NewYear,
			IsMain:         data.IsMain,
		}
		newPassport, err = svc.PassportService.Create(tx, createPassportData)
		if err != nil {
			return err
		}

		for _, chapter := range passportExExists.Chapters {
			chapterCreateData := &service.CreateOneChapterData{
				Name:        chapter.Name,
				PassportID:  newPassport.ID,
				Description: chapter.Description,
				Text:        chapter.Text,
				OrderNumber: chapter.OrderNumber,
			}

			newChapter, err := svc.ChapterService.Create(tx, chapterCreateData)
			if err != nil {
				return err
			}

			for _, partition := range chapter.Partitions {
				objectIDs := make([]int64, 0)
				entityIDs := make([]int64, 0)

				for _, objectTemplate := range partition.Objects {
					for _, object := range objectTemplate.Objects {
						objectIDs = append(objectIDs, object.ID)
					}
				}

				for _, entityTemplate := range partition.Entities {
					for _, en := range entityTemplate.Entities {
						entityIDs = append(entityIDs, en.ID)
					}
				}

				createPartitionData := &service.CreateOnePartitionData{
					Name:        partition.Name,
					ChapterID:   newChapter.ID,
					Description: partition.Description,
					Text:        partition.Text,
					OrderNumber: partition.OrderNumber,
					ObjectIDs:   objectIDs,
					EntityIDs:   entityIDs,
				}

				newPartition, err := svc.PartitionService.Create(tx, createPartitionData)
				if err != nil {
					return err
				}

				for _, route := range partition.Routes {
					setObjectsData := make([]service.SetObjectToRoute, 0)

					for _, obj := range route.RouteObjects {
						var (
							locationDataPtr *service.CreateObjectLocationData
						)
						if obj.LocationID != nil {
							locationVal := service.CreateObjectLocationData{}

							if obj.LocationID.Address != "" {
								locationVal.Address = &obj.LocationID.Address
							}

							if obj.LocationID.Latitude != 0 {
								locationVal.Latitude = &obj.LocationID.Latitude
							}

							if obj.LocationID.Longitude != 0 {
								locationVal.Longitude = &obj.LocationID.Longitude
							}

							locationDataPtr = &locationVal
						}
						setObjectData := service.SetObjectToRoute{
							Name:         obj.Name,
							OrderNumber:  obj.OrderNumber,
							LocationData: locationDataPtr,
						}

						if obj.SourceObject != nil {
							setObjectData.SourceObjectID = &obj.SourceObject.ID
						}

						setObjectsData = append(setObjectsData, setObjectData)
					}

					createRouteData := service.CreateRouteData{
						Name:              route.Name,
						Length:            route.Length,
						Duration:          route.Duration,
						Level:             route.Level,
						MovementWay:       route.MovementWay,
						Seasonality:       route.Seasonality,
						PersonalEquipment: route.PersonalEquipment,
						Dangerous:         route.Dangerous,
						Rules:             route.Rules,
						RouteEquipment:    route.RouteEquipment,
						Geometry:          route.Geometry,
						Objects:           nil,
					}

					if len(setObjectsData) > 0 {
						createRouteData.Objects = &setObjectsData
					}
					createRouteToPartitionData := &service.CreateRouteToPartitionData{
						PartitionID: newPartition.ID,
						Route:       createRouteData,
					}
					_, err = svc.RouteService.CreateToPartition(tx, createRouteToPartitionData)
					if err != nil {
						return err
					}
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return newPassport, nil
}
