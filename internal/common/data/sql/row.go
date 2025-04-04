package sql

type RowScanner interface {
	Scan(dest ...any) error
}
