package chapter

import (
	"context"
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

func (svc *chapterService) simpleCreate(ctx context.Context, data *service.CreateChaptersData) ([]entity.Chapter, error) {
	var (
		names     []string
		repoDatas []repository.CreateChapterData
	)

	for _, chapterData := range data.ChaptersData {
		names = append(names, chapterData.Name)
	}

	chapterExists, err := svc.ChapterRepository.GetByNamesAndPassportID(ctx, names, data.PassportID)
	if err != nil {
		return nil, err
	}

	if len(chapterExists) > 0 {
		return nil, fmt.Errorf("chapter already exists")
	}

	allChapters, err := svc.ChapterRepository.GetByPassportID(ctx, data.PassportID)
	if err != nil {
		return nil, err
	}

	currentOrder := uint(len(allChapters)) + 1

	for _, chapterData := range data.ChaptersData {
		repoData := repository.CreateChapterData{
			Name:        chapterData.Name,
			PassportID:  data.PassportID,
			Description: chapterData.Description,
			Text:        chapterData.Text,
			OrderNumber: currentOrder,
		}

		repoDatas = append(repoDatas, repoData)

		currentOrder++
	}

	return svc.ChapterRepository.CreateMultiply(ctx, repoDatas)
}

//func (svc *municipalityService) CreateToPassport(ctx context.Context, data *service.CreateChaptersData) ([]entity.Chapter, error) {
//	var (
//		orderToID = make(map[uint]int64)
//	)
//
//	chaptersExists, err := svc.GetByPassportID(ctx, data.PassportID)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, chapter := range chaptersExists {
//		orderToID[chapter.OrderNumber] = chapter.ID
//	}
//
//	for _, chapterData := range data.ChaptersData {
//		_, ok := orderToID[chapterData.OrderNumber]
//		if ok {
//			uppOrders(orderToID, chapterData.OrderNumber)
//		}
//
//		orderToID[chapterData.OrderNumber]
//	}
//}

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

//func (svc *municipalityService) UpdateToPassport(ctx context.Context, data *service.UpdateChaptersData) ([]entity.Chapter, error) {
//	//TODO implement me
//	panic("implement me")
//}

func getMaxOrder(orders map[uint]int64) uint {
	maxOrder := uint(0)

	for order := range orders {
		if order > maxOrder {
			maxOrder = order
		}
	}

	return maxOrder
}

func (svc *chapterService) clearOrder(ctx context.Context, order uint, passportID int64) error {
	var (
		orderToChapterID = make(map[uint]int64)
		chapterByID      = make(map[int64]entity.Chapter)
		allOrders        []uint
	)

	allChapters, err := svc.ChapterRepository.GetByPassportID(ctx, passportID)
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

		err = svc.ChapterRepository.ChangeOrder(ctx, chapterID, loopOrder)
	}

	return nil
}

func (svc *chapterService) changeOrder(ctx context.Context, order uint, chapterID, passportID int64) error {
	var (
		orderToChapterID = make(map[uint]int64)
		chapterByID      = make(map[int64]entity.Chapter)
		allOrders        []uint
		chapterBaseOrder uint
	)

	allChapters, err := svc.ChapterRepository.GetByPassportID(ctx, passportID)
	if err != nil {
		return err
	}

	for _, chapter := range allChapters {
		orderToChapterID[chapter.OrderNumber] = chapter.ID
		chapterByID[chapter.ID] = chapter

		if chapter.ID == chapterID {
			chapterBaseOrder = chapter.OrderNumber
		}
	}

	for o := range orderToChapterID {
		allOrders = append(allOrders, o)
	}

	makeNewOrder(orderToChapterID, chapterBaseOrder, order)

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

		chapterLoopID, ok := orderToChapterID[loopOrder]
		if !ok {
			continue
		}

		err = svc.ChapterRepository.ChangeOrder(ctx, chapterLoopID, loopOrder)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *chapterService) Create(ctx context.Context, data *service.CreateOneChapterData) (*entity.Chapter, error) {
	err := svc.clearOrder(ctx, data.OrderNumber, data.PassportID)
	if err != nil {
		return nil, err
	}

	repoData := &repository.CreateChapterData{
		Name:        data.Name,
		PassportID:  data.PassportID,
		Description: data.Description,
		Text:        data.Text,
		OrderNumber: data.OrderNumber,
	}

	return svc.ChapterRepository.Create(ctx, repoData)
}

func (svc *chapterService) Update(ctx context.Context, data *service.UpdateChapterData) (*entity.Chapter, error) {
	chapter, err := svc.ChapterRepository.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if data.Name != nil {
		chapter.Name = *data.Name
	}

	if data.Description != nil {
		chapter.Description = *data.Description
	}

	if data.Text != nil {
		chapter.Text = *data.Text
	}

	if data.OrderNumber != nil && *data.OrderNumber != chapter.OrderNumber {
		err = svc.changeOrder(ctx, *data.OrderNumber, data.ID, chapter.PassportID)
		if err != nil {
			return nil, err
		}
	}

	err = svc.ChapterRepository.Update(ctx, chapter)
	if err != nil {
		return nil, err
	}

	return svc.ChapterRepository.GetByID(ctx, chapter.ID)
}

func (svc *chapterService) DeleteToPassport(ctx context.Context, ids []int64, passportID int64) error {
	var (
		err      error
		maxOrder uint
	)

	allChapters, err := svc.ChapterRepository.GetByPassportID(ctx, passportID)
	if err != nil {
		return err
	}

	maxOrder = uint(len(allChapters))

	for _, id := range ids {
		err = svc.changeOrder(ctx, maxOrder, id, id)
		if err != nil {
			return err
		}

		err = svc.ChapterRepository.Delete(ctx, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *chapterService) GetByPassportID(ctx context.Context, passportID int64) ([]entity.Chapter, error) {
	return svc.ChapterRepository.GetByPassportID(ctx, passportID)
}

func (svc *chapterService) GetByIDAndPassportID(ctx context.Context, id, passportID int64) (*entity.Chapter, error) {
	return svc.ChapterRepository.GetByIDAndPassportID(ctx, id, passportID)
}

func (svc *chapterService) GetByIDsAndPassportID(ctx context.Context, ids []int64, passportID int64) ([]entity.Chapter, error) {
	return svc.ChapterRepository.GetByIDsAndPassportID(ctx, ids, passportID)
}
