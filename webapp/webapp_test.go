package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/riccardomc/k8sbhw/webapp/datastore"
)

func TestStoreHandler(t *testing.T) {

	t.Run("Test GET", func(*testing.T) {
		dataStore := datastore.NewSliceDataStore()
		dataStore.Add(datastore.Record{"a", "b"})
		dataStore.Add(datastore.Record{"c", "1"})

		req, err := http.NewRequest("GET", "/store", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(getStoreHandler(dataStore))

		handler.ServeHTTP(rr, req)

		expectedStatus := http.StatusOK
		if actualStatus := rr.Code; actualStatus != expectedStatus {
			t.Errorf("Unexpected HTTP code: %d != %d", actualStatus, expectedStatus)
		}

		expectedBody := `[{"Key":"a","Value":"b"},{"Key":"c","Value":"1"}]`
		actualBody := strings.Trim(rr.Body.String(), "\n")
		if actualBody != expectedBody {
			t.Errorf("Unexpected body: '%s' != '%s'", actualBody, expectedBody)
		}
	})

	t.Run("Test PUT", func(*testing.T) {
		dataStore := datastore.NewSliceDataStore()

		record := `{"Key": "bla", "Value": "lots"}`
		req, err := http.NewRequest("PUT", "/store", strings.NewReader(record))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(getStoreHandler(dataStore))

		handler.ServeHTTP(rr, req)

		expectedStatus := http.StatusCreated
		if actualStatus := rr.Code; actualStatus != expectedStatus {
			t.Errorf("Unexpected HTTP code: %d != %d", actualStatus, expectedStatus)
		}

		expectedBody := `OK`
		actualBody := strings.Trim(rr.Body.String(), "\n")
		if actualBody != expectedBody {
			t.Errorf("Unexpected body: '%s' != '%s'", actualBody, expectedBody)
		}
	})

	t.Run("Test DELETE", func(*testing.T) {
		dataStore := datastore.NewSliceDataStore()
		dataStore.Add(datastore.Record{"bla", "somevalue"})
		dataStore.Add(datastore.Record{"blu", "someothervalue"})

		record := `{"Key": "bla", "Value": ""}`
		req, err := http.NewRequest("DELETE", "/store", strings.NewReader(record))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(getStoreHandler(dataStore))

		handler.ServeHTTP(rr, req)

		expectedStatus := http.StatusOK
		if actualStatus := rr.Code; actualStatus != expectedStatus {
			t.Errorf("Unexpected HTTP code: %d != %d", actualStatus, expectedStatus)
		}

		expectedBody := `OK`
		actualBody := strings.Trim(rr.Body.String(), "\n")
		if actualBody != expectedBody {
			t.Errorf("Unexpected body: '%s' != '%s'", actualBody, expectedBody)
		}

		expectedSize := 1
		if actualSize := dataStore.Size(); actualSize != expectedSize {
			t.Errorf("Unexpected size: '%d' != '%d'", actualBody, expectedBody)
		}
	})
}
