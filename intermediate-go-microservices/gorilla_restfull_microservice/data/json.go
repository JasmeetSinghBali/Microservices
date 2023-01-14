package data

import (
	"encoding/json"
	"io"
)

/*
to decode json ----> data via NewDecoder via io.Reader to the given interface
*/
func FromJSON(i interface{}, r io.Reader) error {

	d := json.NewDecoder(r)
	return d.Decode(i)
}

/*
to encode  data ----> json via NewEncoder via io.Writer to the given interface
*/
func ToJSON(i interface{}, w io.Writer) error {

	e := json.NewEncoder(w)
	return e.Encode(i)
}
