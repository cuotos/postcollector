package main

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/something/else", strings.NewReader(`{"name":"dan"}`))

	actual := &bytes.Buffer{}

	h := getPostHandler(actual)
	h(rr,req)

	fmt.Println(actual.String())
}