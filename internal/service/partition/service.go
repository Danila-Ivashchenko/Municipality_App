package partition

import (
	"context"
	"errors"
	"fmt"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
	"sort"
)

func uppOrders(orderMap map[uint]int64, order uint) {
	for i := len(orderMap); i >= int(order); i-- {
		currentOrder := uint(i)
		chapterID := orderMap[currentOrder]

		orderMap[currentOrder+1] = chapterID
		delete(orderMap, currentOrder)
	}
}

func makeNewOrder(orderMap map[uint]int64, baseOrder, newOrder uint) {
	item := orderMap[baseOrder]

	for i := baseOrder; i > newOrder; i-- {
		chapterID := orderMap[i-1]

		orderMap[i] = chapterID
	}

	for i := baseOrder; i < newOrder; i++ {
		chapterID := orderMap[i+1]

		orderMap[i] = chapterID
	}

	orderMap[newOrder] = item
}

func validateOrders(orders []uint) error {
	var (
		currentOrder uint
	)

	for i := uint(0); i < uint(len(orders)); i++ {
		currentOrder = orders[i]

		if currentOrder != i+1 {
			return fmt.Errorf("invalid order number %d", currentOrder)
		}
	}

	return nil
}

func getMaxOrder(orders map[uint]int64) uint {
	maxOrder := uint(0)

	for order := range orders {
		if order > maxOrder {
			maxOrder = order
		}
	}

	return maxOrder
}

func (svc *partitionService) clearOrder(ctx context.Context, order uint, passportID int64) error {
	var (
		orderToChapterID = make(map[uint]int64)
		chapterByID      = make(map[int64]entity.Partition)
		allOrders        []uint
	)

	allChapters, err := svc.PartitionRepository.GetByChapterID(ctx, passportID)
	if err != nil {
		return err
	}

	for _, chapter := range allChapters {
		orderToChapterID[chapter.OrderNumber] = chapter.ID
		chapterByID[chapter.ID] = chapter
	}

	for o := range orderToChapterID {
		allOrders = append(allOrders, o)
	}

	uppOrders(orderToChapterID, order)

	orderToChapterID[order] = 0

	sort.Slice(allOrders, func(i, j int) bool {
		return allOrders[i] < allOrders[j]
	})

	err = validateOrders(allOrders)
	if err != nil {
		return err
	}

	currentOrder := getMaxOrder(orderToChapterID)

	for currentOrder > 0 {
		loopOrder := currentOrder
		currentOrder--

		chapterID, ok := orderToChapterID[loopOrder]
		if !ok || chapterID == 0 {
			continue
		}

		err = svc.PartitionRepository.ChangeOrder(ctx, chapterID, loopOrder)
	}

	return nil
}

func (svc *partitionService) changeOrder(ctx context.Context, order uint, partitionID, chapterID int64) error {
	var (
		orderToPartitionID = make(map[uint]int64)
		partitionByID      = make(map[int64]entity.Partition)
		allOrders          []uint
		partitionBaseOrder uint
	)

	allPartitions, err := svc.PartitionRepository.GetByChapterID(ctx, chapterID)
	if err != nil {
		return err
	}

	for _, partition := range allPartitions {
		orderToPartitionID[partition.OrderNumber] = partition.ID
		partitionByID[partition.ID] = partition

		if partition.ID == partitionID {
			partitionBaseOrder = partition.OrderNumber
		}
	}

	for o := range orderToPartitionID {
		allOrders = append(allOrders, o)
	}

	makeNewOrder(orderToPartitionID, partitionBaseOrder, order)

	sort.Slice(allOrders, func(i, j int) bool {
		return allOrders[i] < allOrders[j]
	})

	err = validateOrders(allOrders)
	if err != nil {
		return err
	}

	currentOrder := getMaxOrder(orderToPartitionID)

	for currentOrder > 0 {
		loopOrder := currentOrder
		currentOrder--

		chapterLoopID, ok := orderToPartitionID[loopOrder]
		if !ok {
			continue
		}

		err = svc.PartitionRepository.ChangeOrder(ctx, chapterLoopID, loopOrder)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *partitionService) Create(ctx context.Context, data *service.CreateOnePartitionData) (*entity.PartitionEx, error) {
	var (
		objectIDs            []int64
		objectTemplateMapIDs = make(map[int64]struct{})
		objectTemplateIDs    []int64

		entityIDs            []int64
		entityTemplateMapIDs = make(map[int64]struct{})
		entityTemplateIDs    []int64
	)

	allChapters, err := svc.PartitionRepository.GetByChapterID(ctx, data.ChapterID)
	if err != nil {
		return nil, err
	}

	if int(data.OrderNumber) > len(allChapters)+1 {
		data.OrderNumber = uint(len(allChapters) + 1)
	}

	err = svc.clearOrder(ctx, data.OrderNumber, data.ChapterID)
	if err != nil {
		return nil, err
	}

	repoData := &repository.CreatePartitionData{
		Name:        data.Name,
		ChapterID:   data.ChapterID,
		Description: data.Description,
		Text:        data.Text,
		OrderNumber: data.OrderNumber,
	}

	partition, err := svc.PartitionRepository.Create(ctx, repoData)
	if err != nil {
		return nil, err
	}

	objectsToPartition, err := svc.ObjectToPartitionService.ActualizeToPartition(ctx, partition.ID, data.ObjectIDs)
	if err != nil {
		return nil, err
	}

	for _, o := range objectsToPartition {
		objectIDs = append(objectIDs, o.ObjectID)
	}

	objects, err := svc.ObjectService.GetByIDs(ctx, objectIDs)
	if err != nil {
		return nil, err
	}

	for _, o := range objects {
		objectTemplateMapIDs[o.ObjectTemplateID] = struct{}{}
	}

	for templateID := range objectTemplateMapIDs {
		objectTemplateIDs = append(objectTemplateIDs, templateID)
	}

	templatesEx, err := svc.ObjectTemplateService.GetExByIDs(ctx, objectTemplateIDs)
	if err != nil {
		return nil, err
	}

	entitiesToPartition, err := svc.EntityToPartitionService.ActualizeToPartition(ctx, partition.ID, data.EntityIDs)
	if err != nil {
		return nil, err
	}

	for _, e := range entitiesToPartition {
		entityIDs = append(entityIDs, e.EntityID)
	}

	entities, err := svc.EntityService.GetByIDs(ctx, entityIDs)
	if err != nil {
		return nil, err
	}

	for _, e := range entities {
		entityTemplateMapIDs[e.EntityTemplateID] = struct{}{}
	}

	for templateID := range entityTemplateMapIDs {
		entityTemplateIDs = append(entityTemplateIDs, templateID)
	}

	entitiesTemplatesEx, err := svc.EntityTemplateService.GetExByIDs(ctx, entityTemplateIDs)
	if err != nil {
		return nil, err
	}

	result := entity.NewPartitionEx(*partition, templatesEx, entitiesTemplatesEx)

	return &result, err
}

func (svc *partitionService) getObjectsToPartition(ctx context.Context, partitionID int64) ([]entity.ObjectTemplateEx, error) {
	var (
		objectIDs            []int64
		objectTemplateMapIDs = make(map[int64]struct{})
		objectTemplateIDs    []int64
	)

	objectsToPartition, err := svc.ObjectToPartitionService.GetToPartition(ctx, partitionID)
	if err != nil {
		return nil, err
	}

	for _, o := range objectsToPartition {
		objectIDs = append(objectIDs, o.ObjectID)
	}

	objects, err := svc.ObjectService.GetByIDs(ctx, objectIDs)
	if err != nil {
		return nil, err
	}

	for _, o := range objects {
		objectTemplateMapIDs[o.ObjectTemplateID] = struct{}{}
	}

	for templateID := range objectTemplateMapIDs {
		objectTemplateIDs = append(objectTemplateIDs, templateID)
	}

	return svc.ObjectTemplateService.GetExByIDs(ctx, objectTemplateIDs)
}

func (svc *partitionService) getEntitiesToPartition(ctx context.Context, partitionID int64) ([]entity.EntityTemplateEx, error) {
	var (
		entityIDs            []int64
		entityTemplateMapIDs = make(map[int64]struct{})
		entityTemplateIDs    []int64
	)

	entitiesToPartition, err := svc.EntityToPartitionService.GetToPartition(ctx, partitionID)
	if err != nil {
		return nil, err
	}

	for _, e := range entitiesToPartition {
		entityIDs = append(entityIDs, e.EntityID)
	}

	entities, err := svc.EntityService.GetByIDs(ctx, entityIDs)
	if err != nil {
		return nil, err
	}

	for _, e := range entities {
		entityTemplateMapIDs[e.EntityTemplateID] = struct{}{}
	}

	for templateID := range entityTemplateMapIDs {
		entityTemplateIDs = append(entityTemplateIDs, templateID)
	}

	return svc.EntityTemplateService.GetExByIDs(ctx, entityTemplateIDs)
}

func (svc *partitionService) Update(ctx context.Context, data *service.UpdatePartitionData) (*entity.PartitionEx, error) {
	partition, err := svc.PartitionRepository.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if partition == nil {
		return nil, errors.New("invalid partition id")
	}

	if data.Name != nil {
		partition.Name = *data.Name
	}

	if data.Description != nil {
		partition.Description = *data.Description
	}

	if data.Text != nil {
		partition.Text = *data.Text
	}

	if data.OrderNumber != nil && *data.OrderNumber != partition.OrderNumber {
		err = svc.changeOrder(ctx, *data.OrderNumber, data.ID, partition.ChapterID)
		if err != nil {
			return nil, err
		}
	}

	err = svc.PartitionRepository.Update(ctx, partition)
	if err != nil {
		return nil, err
	}

	if data.ObjectIDs != nil {
		_, err = svc.ObjectToPartitionService.ActualizeToPartition(ctx, partition.ID, *data.ObjectIDs)
		if err != nil {
			return nil, err
		}
	}

	if data.EntityIDs != nil {
		_, err = svc.EntityToPartitionService.ActualizeToPartition(ctx, partition.ID, *data.EntityIDs)
		if err != nil {
			return nil, err
		}
	}

	objectsEx, err := svc.getObjectsToPartition(ctx, partition.ID)
	if err != nil {
		return nil, err
	}

	entitiesEx, err := svc.getEntitiesToPartition(ctx, partition.ID)
	if err != nil {
		return nil, err
	}

	partitionNew, err := svc.PartitionRepository.GetByID(ctx, partition.ID)
	if err != nil {
		return nil, err
	}

	result := entity.NewPartitionEx(*partitionNew, objectsEx, entitiesEx)

	return &result, nil
}

func (svc *partitionService) DeleteToChapter(ctx context.Context, ids []int64, chapterID int64) error {
	var (
		err      error
		maxOrder uint
	)

	allPartitions, err := svc.PartitionRepository.GetByChapterID(ctx, chapterID)
	if err != nil {
		return err
	}

	maxOrder = uint(len(allPartitions))

	for _, id := range ids {
		err = svc.changeOrder(ctx, maxOrder, id, chapterID)
		if err != nil {
			return err
		}

		err = svc.PartitionRepository.Delete(ctx, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *partitionService) GetByChapterID(ctx context.Context, chapterID int64) ([]entity.Partition, error) {
	return svc.PartitionRepository.GetByChapterID(ctx, chapterID)
}

func (svc *partitionService) GetByIDAndChapterID(ctx context.Context, id, chapterID int64) (*entity.Partition, error) {
	return svc.PartitionRepository.GetByIDAndChapterID(ctx, id, chapterID)
}

func (svc *partitionService) GetByIDsAndChapterID(ctx context.Context, ids []int64, chapterID int64) ([]entity.Partition, error) {
	return svc.PartitionRepository.GetByIDsAndChapterID(ctx, ids, chapterID)
}

func (svc *partitionService) GetByIDs(ctx context.Context, ids []int64) ([]entity.Partition, error) {
	return svc.PartitionRepository.GetByIDs(ctx, ids)
}

func (svc *partitionService) GetExByID(ctx context.Context, id int64) (*entity.PartitionEx, error) {
	partition, err := svc.PartitionRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if partition == nil {
		return nil, errors.New("invalid partition id")
	}

	objectsEx, err := svc.getObjectsToPartition(ctx, partition.ID)
	if err != nil {
		return nil, err
	}

	entitiesEx, err := svc.getEntitiesToPartition(ctx, partition.ID)
	if err != nil {
		return nil, err
	}

	result := entity.NewPartitionEx(*partition, objectsEx, entitiesEx)

	return &result, nil
}
