package citizen

//Database infrastructure will implement this interface
type ProfileRepo interface {
	Find() []byte
}
