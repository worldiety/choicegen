# choicegen
This Go module provides a boilerplate generator for modelling compiler checked tagged union / sum types / choice types.

In a limited way, this can be used like the proposed sum types on general interfaces:

* https://github.com/golang/go/issues/57644

A minor selection of related issues:
* https://github.com/golang/go/issues/19412
* https://github.com/golang/go/issues/19814
* https://github.com/golang/go/issues/52096
* ...

As proposed in ticket #57644 we use the regular generic syntax, however you must write it in an ignored source code file.
If invoked, the choicgen tool parses all files with the following build tags:

```go
//go:build choice

package mypkg

// MyChoiceType documentation is copied over.
type MyChoiceType interface {
	// optionally error is supported which creates a simple boilerplate (Error string) method.
    Inc | Dec | None
}
```

Then, it generates a `choicetypes.gen.go` files, which contains all the boilerplate.

## Example

See the `example` package.

## What about a tagged union?

E.g. Rust implements its enum type using a low level tagged union, but questions about how the precise GC should work in Go, arise. 
One could still use an interface, but that double boxing does not make much sense and breaks the Go interface polymorphism entirely.
We experimented with pre-generated generic (tuple like) union types, which results in a lot of mapping boilerplate, which naturally disappears when using plain interfaces.

## Limitations

Most notably are the following limitations of the current approach:
* a real closed sum type cannot be expressed in Go and using type embedding, there may be awkward situations
* because the generator creates a bunch of marker methods, this kind of modelling is only applicable for named package local types
* a sum type can be always nil, even though its members never can't. 
* everything escapes to the heap (as always with regular interfaces). 