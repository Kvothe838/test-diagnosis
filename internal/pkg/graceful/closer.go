package graceful

// Closer contains the method that is called once
// the OS sends shutdown signal to the application.
type Closer interface {
	Close() error
}
