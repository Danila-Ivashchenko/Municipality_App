package entity

import "context"

func (svc *entityService) DeleteMultiple(ctx context.Context, ids []int64, templateID int64) error {
	entities, err := svc.EntityRepository.GetByIDsAndTemplateID(ctx, ids, templateID)
	if err != nil {
		return err
	}

	if len(entities) == 0 {
		return nil
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		for _, entity := range entities {
			err = svc.EntityRepository.Delete(ctx, entity.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
