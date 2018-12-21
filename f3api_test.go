package main

/*
Copyright 2018 Bruno Moura <brunotm@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/brunotm/f3api/api.v1"
	"github.com/brunotm/f3api/models"
	"github.com/brunotm/f3api/server"
	"github.com/brunotm/f3api/store/badgerdb"
)

var (
	paymentBytes = []byte(`{
		"type": "Payment",
		"version": 0,
		"organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		"attributes": {
			"amount": "100.21",
			"beneficiary_party": {
				"account_name": "W Owens",
				"account_number": "31926819",
				"account_number_code": "BBAN",
				"account_type": 0,
				"address": "1 The Beneficiary Localtown SE2",
				"bank_id": "403000",
				"bank_id_code": "GBDSC",
				"name": "Wilfred Jeremiah Owens"
			},
			"charges_information": {
				"bearer_code": "SHAR",
				"sender_charges": [
					{
						"amount": "5.00",
						"currency": "GBP"
					},
					{
						"amount": "10.00",
						"currency": "USD"
					}
				],
				"receiver_charges_amount": "1.00",
				"receiver_charges_currency": "USD"
			},
			"currency": "GBP",
			"debtor_party": {
				"account_name": "EJ Brown Black",
				"account_number": "GB29XABC10161234567801",
				"account_number_code": "IBAN",
				"address": "10 Debtor Crescent Sourcetown NE1",
				"bank_id": "203301",
				"bank_id_code": "GBDSC",
				"name": "Emelia Jane Brown"
			},
			"end_to_end_reference": "Wil piano Jan",
			"fx": {
				"contract_reference": "FX123",
				"exchange_rate": "2.00000",
				"original_amount": "200.42",
				"original_currency": "USD"
			},
			"numeric_reference": "1002001",
			"payment_id": "123456789012345678",
			"payment_purpose": "Paying for goods/services",
			"payment_scheme": "FPS",
			"payment_type": "Credit",
			"processing_date": "2017-01-18",
			"reference": "Payment for Em's piano lessons",
			"scheme_payment_sub_type": "InternetBanking",
			"scheme_payment_type": "ImmediatePayment",
			"sponsor_party": {
				"account_number": "56781234",
				"bank_id": "123123",
				"bank_id_code": "GBDSC"
			}
		}
	}`)

	store *badgerdb.Store
)

func buildServer() (srv *server.Server, err error) {
	store, err = badgerdb.Open(*dataPath)
	if err != nil {
		return nil, err
	}

	config := server.Config{}
	config.Addr = "localhost:8080"
	srv = server.New(config, store)
	api.AddAllRoutes(srv)

	return srv, nil

}
func TestPaymentCRUD(t *testing.T) {
	srv, err := buildServer()
	assert.Nil(t, err)

	// Create
	req, err := http.NewRequest(http.MethodPost, "/api/v1/payment", bytes.NewReader(paymentBytes))
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	r := w.Result()
	assert.True(t, r.StatusCode == http.StatusOK)

	buf, err := ioutil.ReadAll(r.Body)
	assert.Nil(t, err)

	resp := &response{}

	err = json.Unmarshal(buf, resp)
	t.Log(string(buf))
	assert.Nil(t, err)

	// READ
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/payment/%s", resp.ID), bytes.NewReader(paymentBytes))
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	r = w.Result()
	assert.True(t, r.StatusCode == http.StatusOK)

	buf, err = ioutil.ReadAll(r.Body)
	assert.Nil(t, err)

	payment := &models.Payment{}

	err = json.Unmarshal(buf, payment)
	assert.Nil(t, err)
	assert.True(t, payment.IsValid())

	// UPDATE
	req, err = http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/payment/%s", resp.ID), bytes.NewReader(paymentBytes))
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	r = w.Result()
	assert.True(t, r.StatusCode == http.StatusOK)

	buf, err = ioutil.ReadAll(r.Body)
	assert.Nil(t, err)

	resp = &response{}

	err = json.Unmarshal(buf, resp)
	t.Log(string(buf))
	assert.Nil(t, err)

	// DELETE
	req, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/payment/%s", resp.ID), bytes.NewReader(paymentBytes))
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	r = w.Result()
	assert.True(t, r.StatusCode == http.StatusOK)

	buf, err = ioutil.ReadAll(r.Body)
	assert.Nil(t, err)

	resp = &response{}

	err = json.Unmarshal(buf, resp)
	t.Log(string(buf))
	assert.Nil(t, err)

	assert.Nil(t, store.Close())

}

func TestPaymentsGET(t *testing.T) {
	srv, err := buildServer()
	assert.Nil(t, err)

	// Create
	req, err := http.NewRequest(http.MethodPost, "/api/v1/payment", bytes.NewReader(paymentBytes))
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	r := w.Result()
	assert.True(t, r.StatusCode == http.StatusOK)

	buf, err := ioutil.ReadAll(r.Body)
	assert.Nil(t, err)

	resp := &response{}

	err = json.Unmarshal(buf, resp)
	t.Log(string(buf))
	assert.Nil(t, err)

	// READ
	req, err = http.NewRequest(http.MethodGet, "/api/v1/payments", bytes.NewReader(paymentBytes))
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	r = w.Result()
	assert.True(t, r.StatusCode == http.StatusOK)

	buf, err = ioutil.ReadAll(r.Body)
	assert.Nil(t, err)

	payments := &api.Payments{}

	err = json.Unmarshal(buf, payments)
	assert.Nil(t, err)
	for _, payment := range payments.Data {
		assert.True(t, payment.IsValid())
	}

	assert.Nil(t, store.Close())
}

type response struct {
	ID     string
	Reason string
	Links  struct {
		Self string
	}
}
