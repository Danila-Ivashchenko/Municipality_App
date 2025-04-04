package entity

type Permission uint8

const (
	None Permission = iota
	Read
	Write
	Delete
)

func (p Permission) Have(other Permission) bool {
	return uint8(p) <= uint8(other)
}

func PermissionFromUint8(p uint8) Permission {
	switch p {
	case 1:
		return Read
	case 2:
		return Write
	case 3:
		return Delete
	default:
		return None
	}
}

func (p Permission) ToUint8() uint8 {
	return uint8(p)
}
