package region

import "municipality_app/internal/domain/entity"

type regionView struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func newRegionView(i *entity.Region) *regionView {
	if i == nil {
		return nil
	}

	return &regionView{
		ID:   i.ID,
		Name: i.Name,
		Code: i.Code,
	}
}

func newRegionsView(regions []entity.Region) []regionView {
	result := make([]regionView, 0, len(regions))

	for _, region := range regions {
		result = append(result, *newRegionView(&region))
	}

	return result
}
