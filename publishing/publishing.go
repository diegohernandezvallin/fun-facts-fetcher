package publishing

type Publisher interface {
	Publish(message string) error
}
