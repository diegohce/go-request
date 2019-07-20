// Copyright (c) 2019, Diego Cena.
// Use of this source code is governed by an AGPL-3.0
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/diegohce/go-request

/*
Package go-request is a simple implementation of the builder pattern
to create a new *http.Request object.

Example:

	import (
		"log"
		"net/http"

		"github.com/diegohce/go-request"
	)

	func main() {
		rb := &request.RequestBuilder{}

		req, err := rb.Host("localhost:8080").URL("/some/path").
			SetValue("name", "diego").
			Build()

		if err != nil {
			log.Fatal(err)
		}

		c := &http.Client{}

		res, err := c.Do(req)

		...
	}
*/
package request
