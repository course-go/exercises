# Generics & Testing

The goal of this exercise is to practice generics and testing package in Go.

## Steps

One note, you can do these steps in reverse, i.e. writting the tests first and implementing the data structure afterwards. You may know this technique as test-driven development.

### Data Structure

The Go standard libary implements some data structures in the [container](https://pkg.go.dev/container) package (linked-list, heap, ring buffer to be precise). Unfortunately, these data structures do not use generics and use the `any` interface.

1. Choose the data structure you would like to implement.
    - linked-list
    - queue
    - heap
    - ring buffer
    - or any other
2. Create a new module and a new Go file.
    - The names are completely up to you, e.g. `<data-structure>.go`.
3. Implement the data structure.
    - No need to implement all of the methods you can think of, just some will do so the data structure is usable.
    - Do not create the `main()` function. We are writting a library not an executable.

### Testing

Having an implementation is nice, but we can't be sure it works right without properly testing it.

1. Create a new Go file named `<data-structure>_test.go` in the same package.
    - Go uses the `_test.go` suffix for all of the test files. 
    - If it lacks this suffix, the `go test` command will never run the tests in it.
2. Implement a couple of tests for your data structure.
    - All the test functions are prefixed with `Test` in their names and receive `t *testing.T` as argument.
    - If any function lacks this prefix, the `go test` command will never run it.
    - Test functions with the `*t.Testing` parameter missing will not even compile.
    - One example of such test function:

```
func TestQueueEmpty(t *testing.T) {
	q := New[rune]()

	actual := q.Values()

	if len(actual) != 0 {
		t.Errorf("%d != %d", 0, actual)
	}
}
```

It worth pointing out that comparing expected and actual state just using the if statements and the `*t.Testing` parameter is quite tidious. Luckily, assertion libraries exist. One such library is called [testify](github.com/stretchr/testify).

3. To import testify execute: `go get github.com/stretchr/testify`.
4. Use it to test you implementation.
    - The previous test could be rewriten like so:

```
func TestQueueEmpty(t *testing.T) {
	q := New[rune]()

	actual := q.Values()

	assert.Equal(t, 0, len(actual))
}
```

4. Executing the tests is as simple as running: `go test ./...`
    - You can also add the `-v` for verbosity.
    - You can also specify which specific tests to run with `-run X` flag. This command runs test with names that match the `X` regex.
    - As always, `go help test` is your friend.

### Coverage

1. To generate coverage of your tests add the `-cover` flag to your `go test` like so: `go test -cover -v ./...`.
2. To generate a coverage profile execute: `go test -coverprofile=c.out ./...`.
3. You can inspect the code coverage using the cover tool: `go tool cover -html="c.out"`.
    - This can help you to check out code your tests do not cover.
