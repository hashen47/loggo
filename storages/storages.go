package storages

type StorageI interface {
	Name() string
	Notify(text string) error
}
