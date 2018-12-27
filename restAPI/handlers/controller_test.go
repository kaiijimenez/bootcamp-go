package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCart(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/cart/create", nil)
	if err != nil {
		t.Fatal("Error in the request", err)
	}
	w := httptest.NewRecorder()
	CreateCart(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	fmt.Println(body)

}
