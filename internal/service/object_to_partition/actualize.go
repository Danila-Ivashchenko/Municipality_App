package object_to_partition

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *objectToPartitionService) ActualizeToPartition(ctx context.Context, partitionID int64, objectIDs []int64) ([]entity.ObjectToPartition, error) {
	var (
		objectExists   = make(map[int64]struct{})
		objectProvided = make(map[int64]struct{})

		result []entity.ObjectToPartition
	)

	allObjects, err := svc.GetToPartition(ctx, partitionID)
	if err != nil {
		return nil, err
	}

	for _, objectID := range objectIDs {
		objectProvided[objectID] = struct{}{}
	}

	for _, object := range allObjects {
		objectExists[object.ObjectID] = struct{}{}
	}

	for objectID := range objectExists {
		if _, ok := objectProvided[objectID]; !ok {
			objectToPartition := &entity.ObjectToPartition{
				ObjectID:    objectID,
				PartitionID: partitionID,
			}

			err = svc.ObjectToPartitionRepository.Delete(ctx, objectToPartition)
			if err != nil {
				return nil, err
			}
		}
	}

	for objectID := range objectProvided {
		if _, ok := objectExists[objectID]; !ok {
			objectToPartition := &entity.ObjectToPartition{
				ObjectID:    objectID,
				PartitionID: partitionID,
			}

			err = svc.ObjectToPartitionRepository.Create(ctx, objectToPartition)

			result = append(result, *objectToPartition)
		}
	}

	return result, nil
}
