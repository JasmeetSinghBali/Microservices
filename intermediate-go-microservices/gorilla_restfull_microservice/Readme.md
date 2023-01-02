> ## Gorilla web toolkit [restfull microservices]

reff: https://www.gorillatoolkit.org/

> Subrouter reff: https://pkg.go.dev/github.com/gorilla/mux#section-readme

- subrouter is effective when need is to direct a particular type of http request to particular handler via [ Methods method ]

- Note: function/method name starting with capital letter can only be permitted to be accessed from another package

> Middlewares reff: https://github.com/gorilla/mux#middleware

- the aim is to replace repeatable code snippet like marshal/unmarshal or validation with reusable function middlewares

            # to add a middleware to router
            Router.Use()

            # example
            r := mux.NewRouter()
            r.HandleFunc("/",handler)
            r.Use(loggingMiddleware)

- http context (key:value) pairs can be attached to request object that is common & has common reff throughout the request-response cycle. reff: middlewares/validationProduct.go
