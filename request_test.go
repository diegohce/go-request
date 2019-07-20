package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddValue(t *testing.T) {

	rb := &RequestBuilder{}

	rb.AddValue("key", "to the universe")

	qstr := rb.Values()
	if qstr != "key=to+the+universe" {
		t.Fatal("Got:", qstr, "Want: key=to+the+universe")
	}
}

func TestWithBody(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "CENA")
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)

	rb := &RequestBuilder{}

	req, err := rb.Method("POST").
		UserPassword("diego", "dirtysecret").
		Scheme("").
		Host(u.Host).URL("/test").
		SetValue("name", "diego").
		AddValue("pet", "simona").
		AddValue("pet", "lola").
		AddValue("pet2", "frida").
		DelValue("pet2").
		Payload([]byte("DIEGO")).
		Build()

	if err != nil {
		t.Logf("%+v\n", req)
		t.Fatal(err)
	}

	qstring := rb.Values()
	if qstring != "name=diego&pet=simona&pet=lola" {
		t.Fatal("Query string failed. Got:", qstring, "Want: name=diego&pet=simona&pet=lola")
	}

	c := &http.Client{}

	res, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Logf("%+v\n", res)
		t.Fatal("Bad status code. Got:", res.StatusCode, "Want: 200")
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != "CENA" {
		t.Fatal("Body error. Got:", string(b), "Want: CENA")
	}
}
func TestWithoutBody(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "CENA")
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)

	rb := &RequestBuilder{}

	req, err := rb.Method("GET").
		UserPassword("diego", "dirtysecret").
		Scheme("").
		Host(u.Host).URL("/test").
		SetValue("name", "diego").
		AddValue("pet", "simona").
		AddValue("pet", "lola").
		AddValue("pet2", "frida").
		DelValue("pet2").
		Build()

	if err != nil {
		t.Logf("%+v\n", req)
		t.Fatal(err)
	}

	qstring := rb.Values()
	if qstring != "name=diego&pet=simona&pet=lola" {
		t.Fatal("Query string failed. Got:", qstring, "Want: name=diego&pet=simona&pet=lola")
	}

	c := &http.Client{}

	res, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Logf("%+v\n", res)
		t.Fatal("Bad status code. Got:", res.StatusCode, "Want: 200")
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != "CENA" {
		t.Fatal("Body error. Got:", string(b), "Want: CENA")
	}
}

func TestMinimalBuild(t *testing.T) {

	rb := &RequestBuilder{}

	req, err := rb.Host("localhost").Build()
	if err != nil {
		t.Fatal(err)
	}

	if req.URL.String() != "http://localhost" {
		t.Fatal("Got:", req.URL.String(), "Want: http://localhost")
	}
}

func TestHeaders(t *testing.T) {

	rb := &RequestBuilder{}

	req, err := rb.Host("localhost").
		SetHeader("X-My-Header", "HI").
		AddHeader("Accept", "application/json").
		Build()
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("Accept") != "application/json" {
		t.Fatal("Got:", req.Header.Get("Accept"), "Want: application/json")
	}

	if req.Header.Get("X-My-Header") != "HI" {
		t.Fatal("Got:", req.Header.Get("X-My-Header"), "Want: HI")
	}
}

func TestAddHeader(t *testing.T) {

	rb := &RequestBuilder{}

	req, err := rb.AddHeader("Accept", "A").
		AddHeader("Accept", "B").
		Build()
	if err != nil {
		t.Fatal(err)
	}

	if strings.Join(req.Header["Accept"], ",") != "A,B" {
		t.Fatal("Got:", strings.Join(req.Header["Accept"], ","), "Want: A,B")
	}
}
