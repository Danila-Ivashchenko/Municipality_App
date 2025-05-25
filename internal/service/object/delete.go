package object

import "context"

func (svc *objectService) DeleteMultiple(ctx context.Context, ids []int64, templateID int64) error {
	objects, err := svc.ObjectRepository.GetByIDsAndTemplateID(ctx, ids, templateID)
	if err != nil {
		return err
	}

	if len(objects) == 0 {
		return nil
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		for _, object := range objects {
			err = svc.ObjectRepository.Delete(ctx, object.ID)
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
