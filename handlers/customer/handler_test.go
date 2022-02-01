package customer

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/prashant/layeredArc/models"
)

func TestHandler_Create(t *testing.T) {
	tests := []struct {
		reqBody []byte
		resBody []byte
	}{
		{[]byte(`{"name":"ankit","phone_no":123,"address":"jh"}`), []byte(`{"Id":1,"Name":"ankit","Address":"jh",123}`)},
		{[]byte(`{"name":"ankit","phone_no":"456","address":"jh"}`), []byte(`invalid body`)},
		{[]byte(`{"name":"ravi","phone_no":789,"address":"wb"}`), []byte(`not create customer`)},
	}

	for i, v := range tests {
		req, _ := http.NewRequest("GET", "/customer", bytes.NewReader(v.reqBody))
		w := httptest.NewRecorder()

		a := New(mockDatastore{})

		a.Create(w, req)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.resBody)) {
			t.Errorf("[TEST%d]Faild. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.resBody))
		}

	}

}

type mockDatastore struct{}

func (m mockDatastore) Create(c models.Customer) (models.Customer, error) {
	if c.PhoneNo == 789 {
		return models.Customer{}, errors.New("db error")
	}
	return models.Customer{ID: 1, Name: "ankit", Address: "jh", PhoneNo: 123}, nil
}
