package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

var rootUrl string = "https://YummyDualOpen64.dylanwong8.repl.co";

type ItemBody struct {
  Count int `json:"count"`
  Title string `json:"title"`
  Description string `json:"description"`
}

func TestListItems(t *testing.T){
  response, err := http.Get(rootUrl + "/item");

  if err != nil {
      t.Errorf(err.Error())
  }

  _, err = ioutil.ReadAll(response.Body)
  if err != nil {
      t.Errorf(err.Error())
  }

  if response.StatusCode != 200 {
      t.Errorf("got %q Response, wanted 200 Response", response.StatusCode)
  }
}

func TestCreateItem(t *testing.T){  
  reqbody := &ItemBody{
    Count:    1,
    Title: "xyzasdf",
    Description: "what",
  }

  payloadBuf := new(bytes.Buffer)
  json.NewEncoder(payloadBuf).Encode(reqbody)
  
  response, err := http.Post(rootUrl + "/item", "application/json", payloadBuf);

  if err != nil {
    t.Errorf(err.Error())
  }

  _, err = ioutil.ReadAll(response.Body)
  if err != nil {
    t.Errorf(err.Error())
  }

  if response.StatusCode != 200 {
      t.Errorf("got %q Response, wanted 200 Response", response.StatusCode)
  }
}

func TestBadItem(t *testing.T){  
  reqbody := &ItemBody{
    Count:    -1,
    Title: "x",
    Description: "what",
  }

  payloadBuf := new(bytes.Buffer)
  json.NewEncoder(payloadBuf).Encode(reqbody)
  
  response, err := http.Post(rootUrl + "/item", "application/json", payloadBuf);

  if err != nil {
    t.Errorf(err.Error())
  }

  _, err = ioutil.ReadAll(response.Body)
  if err != nil {
    t.Errorf(err.Error())
  }

  if response.StatusCode != 400 {
      t.Errorf("got %q Response, wanted 400 Response", response.StatusCode)
  }
}
