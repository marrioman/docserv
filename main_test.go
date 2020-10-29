package main

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCRUD(t *testing.T) {
	var newUser = struct {
		username string `json:"username"`
		lastname string `json:"lastname"`
		password string `json:"password"`
	}{
		username: "Yar",
		lastname: "asafa",
		password: "1234",
	}

	client := http.Client{}

	requestByte, _ := json.Marshal(newUser)

	url := "http://localhost:1234/api/user/create"
	response, err := client.Post(url, "application/json", bytes.NewBuffer(requestByte))
	if err != nil {
		t.Errorf("the HTTP request failed with error %s", err)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(errors.Errorf("cannot read response %s", err))
		return
	}

	if response.StatusCode == http.StatusInternalServerError {
		t.Error("status 500")
		return
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("statusCode: %d, Url: http://localhost:1234/api/user/create", response.StatusCode)
	}

	if string(body) != `{"data":"success"}` {
		t.Errorf("not passed")
	}
}
