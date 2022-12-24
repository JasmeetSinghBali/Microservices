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
