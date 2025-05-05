package partition

import (
	"errors"
	"municipality_app/internal/domain/service"
)

type reqPartitionCreateData struct {
	Name        *string `json:"name"`
	Description string  `json:"description"`
	Text        string  `json:"text"`
	OrderNumber *uint   `json:"order_number"`
	ObjectIDs   []int64 `json:"objects"`
	EntityIDs   []int64 `json:"entities"`
}

func (req *reqPartitionCreateData) Validate() error {
	if req.Name == nil {
		return errors.New("name is required")
	}

	if req.OrderNumber == nil {
		return errors.New("order_number is required")
	}

	return nil
}

func (req *reqPartitionCreateData) Convert(chapterID int64) *service.CreateOnePartitionData {
	result := &service.CreateOnePartitionData{
		Name:        *req.Name,
		ChapterID:   chapterID,
		Description: req.Description,
		Text:        req.Text,
		OrderNumber: *req.OrderNumber,
		ObjectIDs:   req.ObjectIDs,
		EntityIDs:   req.EntityIDs,
	}

	return result
}

type reqPartitionUpdateData struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Text        *string  `json:"text"`
	OrderNumber *uint    `json:"order_number"`
	ObjectIDs   *[]int64 `json:"objects"`
	EntityIDs   *[]int64 `json:"entities"`
}

func (req *reqPartitionUpdateData) Convert(id int64) *service.UpdatePartitionData {
	result := &service.UpdatePartitionData{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Text:        req.Text,
		OrderNumber: req.OrderNumber,
		ObjectIDs:   req.ObjectIDs,
		EntityIDs:   req.EntityIDs,
	}

	return result
}
