package webutils

import (
	"errors"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
)

type TestData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestRespondWithJSONEncoding(t *testing.T) {
	RegisterTestingT(t)

	respRec := httptest.NewRecorder()
	RespondWithJSON(respRec, TestData{1, "test1"}, nil)
	Expect(respRec.Code).To(Equal(200))
	Expect(respRec.Body.String()).To(MatchJSON(`{"id":1,"name":"test1"}`))
	Expect(respRec.HeaderMap["Content-Type"][0]).To(Equal("application/json; charset=UTF-8"))
}

func TestRespondWithJSONErrorHandling(t *testing.T) {
	RegisterTestingT(t)

	respRec := httptest.NewRecorder()
	RespondWithJSON(respRec, TestData{1, "test1"}, errors.New("some error"))
	Expect(respRec.Code).To(Equal(500))
	Expect(respRec.Body.String()).To(BeEmpty())
}

func TestRespondWithJSONNoDataHandling(t *testing.T) {
	RegisterTestingT(t)

	respRec := httptest.NewRecorder()
	RespondWithJSON(respRec, nil, nil)
	Expect(respRec.Code).To(Equal(404))
	Expect(respRec.Body.String()).To(MatchJSON(`{}`))
}
