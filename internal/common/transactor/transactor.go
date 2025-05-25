package transactor

import "context"

type Transactor interface {
	Execute(ctx context.Context, fn func(tx context.Context) error) error
}
