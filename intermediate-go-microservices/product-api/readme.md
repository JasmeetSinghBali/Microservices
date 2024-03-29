> ## Restful-api in go [standard libs] reff: product-api folder

- encoding/decoding json [marshalling & unmarshalling]

ref: https://pkg.go.dev/encoding/json

- struct/go data types ---> JSON[marshalling] or JSON ---> struct/go data types[unmarshalling]

> struct tags

- in order to customize the fields that are returned to the client
- adds annotation to the fields, and then parsers can be used to pick up those annotation

```bash
# now this Field will be named as myName when parsed/marshalled and spit back to user/client
Field int `json:"myName"`

# omit returning this field if its empty
Field int `json:"myName,omitempty"`

# dont return this field
Field int `json: "-"`
```

> new encoder over normal marshalling !

- making use of direct encoder is better, as the standard marshalling is somewhat heavy where the direct encoder is lightweight
  ref: https://pkg.go.dev/encoding/json#NewEncoder
