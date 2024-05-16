# Options & Katas

Like the exercise from the last lecture, this exercise involves solving a couple of katas. Additionally, it explores the options builder pattern.

## Sections

### Options builder pattern

In this section, you will explore the "options builder pattern" commonly used in idiomatic Go code.

The `internal/pizza` directory contains an imlementation of a traditional builder pattern as you know it. Your goal is to rewrite it using the options pattern. This pattern replaces the traditional method chaining on a builder type with supplying a set of functions to the constructor that then modify the structure.

To give you a hint:

```go
func New(options ...Option) MyType {
    // construct type
}
```

```go
func WithSomeOption(args) Option {
    return func(t *MyType) {
        // modify type
    }
}
```

If you feel lost or do not know how to go about it, Google it or check [this blog](https://golang.cafe/blog/golang-functional-options-pattern.html) directly.

### Katas

The project also includes a couple of katas in `internal/katas` directory. The katas also contain auxilery functions and data in the subdirectories which you are free to explore, but should not edit in any way. Each kata contains a brief description of the problem which should be fixed.
