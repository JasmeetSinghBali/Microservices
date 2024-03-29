> ## structuring microservices basics [reff: web-server code here is structured as an example]

- seprate out handlers into packages
- have abstracted objects that replacable for future testing
- follow dependency injection patterns & reusable instance/object in struct

- registering the handlers with the ServeMux, multiplexer that handles multiple handlers and registers them with the current http server ref: https://pkg.go.dev/net/http#ServeMux

```bash
    /*
	reff to the greetings handler gh
	injecting the logger into the NewGreetings method implementing the traces/logger instance of Greetings struct for greeting interface
	*/
	gh := handlers.NewGreetings(gl)



	/*
	registering the greetings handler with the servemux with servermux sm instance
	for pattern /
	*/
	sm := http.NewServeMux()
	sm.Handle("/",gh)

```

> Tuning [setting up timeouts, graceful shutdown(to avoid current database transaction or current active client connection to just cut off abruptly)]

ref: https://pkg.go.dev/net/http#Server

- ReadTimeout user-->go server, example for larger files this wud be more for smaller files upload by user this wud be relatively smaller.

- WriteTimeout go server--->user

- IdleTimeout [connection-pooling] i.e connection alive time used to keep the connection active, useful for microservices conected to each other for persistent connection between microservices specically for TLS connection that TCP, ideally this is higher

```bash
/*wait for the request currently been processed, and from this point wont take any more request and after current req are processed it shuts down the server */
	s.Shutdown()
```
