# REST API

The goal of this exercise is to practice implementing a HTTP API using the Go standard library.

## Steps

The assignment consists of implementing a todo app API. (Yes, yes... I know...). You may find it boring, but that's good. We do not want to tackle complex business logic. The sole goal is to practice the [net/http](https://pkg.go.dev/net/http) package's API and convention.

### REST API

The API specification is provided in the form of a OpenAPI Specification. It can be found under the `/docs` directory. Many IDEs allow you to inspect it using a UI. If your IDE or code editor does not support it, you can use the [Swagger Editor](https://editor.swagger.io) or any other similar tool.

Also note, that the initial implementation already contains a in-memory repository implementation. If the implementation of the repository does not fit your needs, feel free to adjust it.

1. Extend the Todo struct with additional fields.
    - The components sections of the OpenAPI spec could help you out.
2. Implement endpoint handlers.
    - Starting with a single endpoint first is adviced.
    - I suggest implementing all of the behaviour in the `main` function first.
    - The endpoints will require new struct definitions for parsing the requests and responses.
3. Create a file for your controllers and move them out of the `main` file.
    - The structure is up to you.
4. Wire all the dependendecies in the `main` function and run the server.
    - This should be the `main` function's only responsibility.

Here is a sample from the [documentation](https://pkg.go.dev/net/http#Get) to get you started:

```
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```
