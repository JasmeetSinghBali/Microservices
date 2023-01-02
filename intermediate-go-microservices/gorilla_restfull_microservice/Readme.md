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

> Validations reff: https://github.com/go-playground/validator

- Essential to prevent top 10 most common OWASP mentioned vulnerability Injections,Broken Auth,sensitive data exposure etc...

- Sanitization inputs from client is always a must

> Test ref: https://stackoverflow.com/questions/28240489/golang-testing-no-test-files

- each file that is suffix with \_test will be considered as test file
- Function that name prefix Test with signature func (t \*testing.T)

            # to run test in vs code
            golang extension shows options to run/debug test
            # or cd to directory containing the _test file &
            go test
