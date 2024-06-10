package graceful

import (
	"context"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
	"time"

	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"
)

const (
	_defaultTimeout = 30 * time.Second
)

var (
	//nolint: gochecknoglobals // default registry
	DefaultRegistry = NewRegistry(Options{})
)

func Register(closer Closer) error {
	return DefaultRegistry.Register(closer)
}

func MustRegister(closer Closer) {
	DefaultRegistry.MustRegister(closer)
}

func Wait() error {
	return DefaultRegistry.Wait()
}

type Options struct {
	Timeout time.Duration
}

// Registry ensures that when the program receives a signal to stop, all registered resources or services are properly closed within a time limit before the program terminates.
type Registry struct {
	Options Options
	closers map[Closer]struct{}
	closed  *atomic.Bool
	lock    sync.Mutex
}

func NewRegistry(options Options) *Registry {
	options = setDefaults(options)
	return &Registry{
		Options: options,
		closers: map[Closer]struct{}{},
		closed:  atomic.NewBool(false),
	}
}

func setDefaults(options Options) Options {
	if options.Timeout == 0 {
		options.Timeout = _defaultTimeout
	}
	return options
}

func (r *Registry) Register(closer Closer) error {
	return r.register(closer)
}

func (r *Registry) MustRegister(closer Closer) {
	if err := r.register(closer); err != nil {
		panic(err)
	}
}

func (r *Registry) register(closer Closer) error {
	if !reflect.TypeOf(closer).Comparable() {
		return ErrNotComparable
	}

	if r.closed.Load() {
		return ErrRegistrationsClosed
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	if _, exist := r.closers[closer]; exist {
		return ErrAlreadyRegistered
	}

	r.closers[closer] = struct{}{}
	return nil
}

func (r *Registry) Wait() error {
	r.closed.Store(true)

	osShutdownSignal := make(chan os.Signal, 1)
	signal.Notify(
		osShutdownSignal,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	<-osShutdownSignal

	ctx, cancel := context.WithTimeout(context.Background(), r.Options.Timeout)
	defer cancel()

	eg := errgroup.Group{}
	for closer := range r.closers {
		eg.Go(closer.Close)
	}

	done := make(chan error, 1)
	go func() {
		done <- eg.Wait()
		close(done)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ErrTimeout
	}
}
