package webutils

import (
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
	RespondWithJSON(respRec, TestData{1, "test1"}, false)
	Expect(respRec.Code).To(Equal(200))
	Expect(respRec.Body.String()).To(MatchJSON(`{"id":1,"name":"test1"}`))
	Expect(respRec.HeaderMap["Content-Type"][0]).To(Equal("application/json; charset=UTF-8"))
}

func TestRespondWithJSONPointerEncoding(t *testing.T) {
	RegisterTestingT(t)

	respRec := httptest.NewRecorder()
	RespondWithJSON(respRec, &TestData{1, "test1"}, false)
	Expect(respRec.Code).To(Equal(200))
	Expect(respRec.Body.String()).To(MatchJSON(`{"id":1,"name":"test1"}`))
	Expect(respRec.HeaderMap["Content-Type"][0]).To(Equal("application/json; charset=UTF-8"))
}

func TestRespondWithJSONSliceEncoding(t *testing.T) {
	RegisterTestingT(t)

	respRec := httptest.NewRecorder()
	RespondWithJSON(respRec, []TestData{TestData{1, "test1"}}, false)
	Expect(respRec.Code).To(Equal(200))
	Expect(respRec.Body.String()).To(MatchJSON(`[{"id":1,"name":"test1"}]`))
	Expect(respRec.HeaderMap["Content-Type"][0]).To(Equal("application/json; charset=UTF-8"))
}

func TestRespondWithJSONErrorHandling(t *testing.T) {
	RegisterTestingT(t)

	respRec := httptest.NewRecorder()
	RespondWithJSON(respRec, TestData{1, "test1"}, true)
	Expect(respRec.Code).To(Equal(500))
	Expect(respRec.Body.String()).To(BeEmpty())
}

func TestRespondWithJSONNoDataHandling(t *testing.T) {
	RegisterTestingT(t)

	respRec := httptest.NewRecorder()
	RespondWithJSON(respRec, nil, false)
	Expect(respRec.Code).To(Equal(404))
	Expect(respRec.Body.String()).To(MatchJSON(`{}`))
}

func TestRespondWithJSONDataNilPointerHandling(t *testing.T) {
	RegisterTestingT(t)

	var testDataPtr *TestData
	respRec := httptest.NewRecorder()
	RespondWithJSON(respRec, testDataPtr, false)
	Expect(respRec.Code).To(Equal(404))
	Expect(respRec.Body.String()).To(MatchJSON(`{}`))
}
