> exploring the internals of http

- http package exposes HandleFunc that register's a function to a path(pattern) on Default Server Mux ref: https://pkg.go.dev/net/http@go1.19.4

- Default Server Mux is http handler & everything related to web-server is related to the http handler

- the handler is effective interface with ResponseWriter & \*Request as signature that implements the handler interface. ref: https://pkg.go.dev/net/http@go1.19.4#Handler

- response-writer interface, to spit back response to user
  https://pkg.go.dev/net/http@go1.19.4#ResponseWriter
  https://pkg.go.dev/net/http@go1.19.4#Response

- better way to handle errors ref: https://pkg.go.dev/net/http@go1.19.4#Error
  NOTE- http.Error does not terminate flow of the request
