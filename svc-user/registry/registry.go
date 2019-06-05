package registry

type Registry interface {
	Registry()
	Close()
	Reconnect()
}
