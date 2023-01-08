package plugin

import (
	"errors"
	"fmt"
	"os"
	"plugin"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

const pluginFile = "plugin.so"

// Connection is a connection to a plugin.
type Connection struct {
	mtx   sync.Mutex
	cache map[any]bool
}

// NewConnection creates a new Connection.
func NewConnection() *Connection {
	return &Connection{cache: make(map[any]bool)}
}

// Connect connects a method with its implementation.
func Connect[T any](c *Connection, target *T) {
	// Check if we already connected this method
	if _, ok := c.cache[target]; ok {
		return
	}

	// Avoid concurrently connecting
	c.mtx.Lock()
	defer c.mtx.Unlock()

	// Synchronously check again if we already connected this method
	if _, ok := c.cache[target]; ok {
		return
	}

	// Get name of the calling function (clean up the name a bit)
	pc, _, _, ok := runtime.Caller(1)
	_ = ok
	fullName := runtime.FuncForPC(pc).Name()
	trimmed := strings.TrimSuffix(fullName, "[...]")
	split := strings.Split(trimmed, ".")
	name := split[len(split)-1]

	// Connect by convention
	ConnectSymbol(name, target)
	c.cache[target] = true
}

func ConnectSymbol[T any](name string, target *T) {
	if _, err := os.Stat(pluginFile); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "could not find plugin %q\n", pluginFile)
		os.Exit(1)
	}
	p, err := plugin.Open(pluginFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening plugin %q\n", pluginFile)
		os.Exit(1)
	}
	sym, err := p.Lookup(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting symbol %q\n", name)
		os.Exit(1)
	}
	// Names in the plugin are associated with pointers to functions.
	// Thus we cannot: *target = sym(T)
	*target = reflect.ValueOf(sym). // *func(...) as reflect.Value
					Elem().         // dereferences to func(...)
					Interface().(T) // any.(func(...))
}
