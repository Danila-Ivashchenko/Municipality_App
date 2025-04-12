package entity_to_partition

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *entityToPartitionService) ActualizeToPartition(ctx context.Context, partitionID int64, entityIDs []int64) ([]entity.EntityToPartition, error) {
	var (
		entityExists   = make(map[int64]struct{})
		entityProvided = make(map[int64]struct{})

		result []entity.EntityToPartition
	)

	allEntitys, err := svc.GetToPartition(ctx, partitionID)
	if err != nil {
		return nil, err
	}

	for _, entityID := range entityIDs {
		entityProvided[entityID] = struct{}{}
	}

	for _, entity := range allEntitys {
		entityExists[entity.EntityID] = struct{}{}
	}

	for entityID := range entityExists {
		if _, ok := entityProvided[entityID]; !ok {
			entityToPartition := &entity.EntityToPartition{
				EntityID:    entityID,
				PartitionID: partitionID,
			}

			err = svc.EntityToPartitionRepository.Delete(ctx, entityToPartition)
			if err != nil {
				return nil, err
			}
		}
	}

	for entityID := range entityProvided {
		if _, ok := entityExists[entityID]; !ok {
			entityToPartition := &entity.EntityToPartition{
				EntityID:    entityID,
				PartitionID: partitionID,
			}

			err = svc.EntityToPartitionRepository.Create(ctx, entityToPartition)

			result = append(result, *entityToPartition)
		}
	}

	return result, nil
}
