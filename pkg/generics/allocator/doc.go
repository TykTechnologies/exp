// The allocator package handles two concerns related to allocations:
//
// 1. It uses sync.Pool to manage an in-memory cache of reusable types.
// 2. Provides a generic interface to take advantage of type safety.
//
// Extensions may focus on measuring allocation pressure.
//
// To use the allocator, typed code must provide a constructor.
// If this was typed code, the following function is expected:
//
// ~~~
// func NewDocument() (*Document, error)
// ~~~
//
// With Go generics, expected usage is:
//
// ~~~
// repo := allocator.New[*Document](NewDocument)
// value := repo.Get()
// ...
// repo.Put(value)
// ~~~
//
// The type must implement a Reset() function.
package allocator
