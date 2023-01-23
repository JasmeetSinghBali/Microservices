> File Uploads in Go with Standard libs

**reff: https://pkg.go.dev/net/http#Request**

- The body io.ReadCloser is an effecient way to handle sent payload to server on request
- when request is sent to go server the payload data is not immediately buffered, the data is indeed can be read gradually over time instead of it at one go i.e limit of data can be processed of a the wholesome data as chunks like 4mb data each time of 4gb data

- the io.Reader can be used to control the number of bytes it read at a time is the power that golang exposes to the developers.

- pre-defined/setup FileServer to serve file content over http requests reff: https://pkg.go.dev/net/http#FileServer

- FileServer automatically determines the content length & wraps the content of file in response stream

- FileServer expects a http.Dir , Dir gives the location of file on file system i.e uri path ---> file dir system
