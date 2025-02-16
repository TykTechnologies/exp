# Package github.com/TykTechnologies/exp/pkg/generics/allocator

The allocator package handles two concerns related to allocations:

1. It uses sync.Pool to manage an in-memory cache of reusable types.
2. Provides a generic interface to take advantage of type safety.

Extensions may focus on measuring allocation pressure.

To use the allocator, typed code must provide a constructor.
If this was typed code, the following function is expected:

~~~
func NewDocument() (*Document, error)
~~~

With Go generics, expected usage is:

~~~
repo := allocator.New[*Document](NewDocument)
value := repo.Get()
...
repo.Put(value)
~~~

The type must implement a Reset() function.

## Types

```go
// Allocator holds a sync.Pool of objects of type T.
type Allocator[T Reseter] struct {
	pool sync.Pool
}
```

```go
// Reseter is the interface that types must implement to be managed by Allocator.
type Reseter interface {
	Reset()
}
```

## Function symbols

- `func New (newFunc func() T) *Allocator[T]`
- `func *Allocator[T].Get () T`
- `func *Allocator[T].Put (t T)`

### New

```go
func New (newFunc func() T) *Allocator[T]
```

New creates an Allocator for type T using the provided constructor.

### Get

```go
func *Allocator[T].Get () T
```

Get retrieves an object from the internal pool.

### Put

```go
func *Allocator[T].Put (t T)
```

Put returns an object to the pool after resetting it.


