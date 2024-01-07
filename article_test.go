package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
  alist := getAllArticles()

  // Check that the length of the list of articles returned is the
  // same as the length of the global variable holding the list
  if len(alist) != len(articleList) {
    t.Fail()
  }

  // Check that each member is identical
  for i, v := range alist {
    if v.Content != articleList[i].Content ||
      v.ID != articleList[i].ID ||
      v.Title != articleList[i].Title {

      t.Fail()
      break
    }
  }
}

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
  r := getRouter(true)

  r.GET("/", showIndexPage)

  // Create a request to send to the above route
  req, _ := http.NewRequest("GET", "/", nil)

  testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
    // Test that the http status code is 200
    statusOK := w.Code == http.StatusOK

    // Test that the page title is "Home Page"
    // You can carry out a lot more detailed tests using libraries that can
    // parse and process HTML pages
    p, err := ioutil.ReadAll(w.Body)
    pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

    return statusOK && pageOK
  })
}