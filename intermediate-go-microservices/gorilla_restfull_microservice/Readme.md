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

- Sanitization inputs from client is always a must.

- ✨ custom validation function can also be added for fields that have custom & dynamic possible combination of allowed values ref: Glaze field in data/products.go that has custom validation function validateGlaze reff: https://github.com/go-playground/validator/blob/master/_examples/custom/main.go

> Test ref: https://stackoverflow.com/questions/28240489/golang-testing-no-test-files

- each file that is suffix with \_test will be considered as test file
- Function that name prefix Test with signature func (t \*testing.T)

            # to run test in vs code
            golang extension shows options to run/debug test
            # or cd to directory containing the _test file &
            go test

> Documentation via OPEN API swagger ref: https://goswagger.io/

- first add swagger:meta doc tags as comments as done in handlers/products.go ref:https://goswagger.io/use/spec/meta.html
- second, install swagger cli & then generate swagger doc

                    # make sure golang must be updated >1.17
                    choco upgrade golang -y
                    # make dummy dir
                    git clone https://github.com/go-swagger/go-swagger
                    # cd to go-swagger
                    go install ./cmd/swagger
                    # cd to gorilla_restfull dir
                    swagger generate spec -o ./swagger.yaml --scan-models
                    go run main.go
                    # curl localhost:PORT/swagger.yaml or visit localhost:PORT/docs

- NOTE: each time when the documentation changes just run swagger generate spec -o ./swagger.yaml --scan-models command
