package api

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"github.com/brunotm/f3api/models"
	"github.com/brunotm/f3api/server"
)

var (
	uuidSpace uuid.UUID
)

func init() {
	var err error
	uuidSpace, err = uuid.Parse("39f41424-331d-502b-a690-e07de62c75f4")
	if err != nil {
		panic(err)
	}
}

// AddPayment handler
func AddPayment(store server.Store) (handle server.Handle) {
	return func(w http.ResponseWriter, r *http.Request, p server.Params) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		var err error
		var buf []byte

		if buf, err = ioutil.ReadAll(r.Body); err != nil || len(buf) == 0 {
			http.Error(w, `{"reason": "bad request"}`, http.StatusBadRequest)
			return
		}

		payment := &models.Payment{}
		if err = json.Unmarshal(buf, payment); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !payment.IsValid() {
			http.Error(w, `{"reason": "invalid payment"}`, http.StatusBadRequest)
			return
		}

		id := uuid.NewSHA1(uuidSpace, buf).String()
		payment.ID = id

		var found bool
		store.Get([]byte(id),
			func(value []byte, err error) {
				if err != nil {
					found = false
					return
				}
				if value == nil {
					found = false
					return
				}
				found = true
			},
		)

		if found {
			http.Error(w,
				fmt.Sprintf(`{"reason":"payment already exists", "id":"%s", "links":{"self":"%s%s/%s"}}`,
					id, r.Host, r.URL.Path, id), http.StatusBadRequest)
			return
		}

		buf, err = json.Marshal(payment)
		if err != nil {
			http.Error(w, `{"reason": "could not serialize payment"}`, http.StatusInternalServerError)
			return
		}

		if err = store.Set([]byte(id), buf); err != nil {
			http.Error(w, `{"reason":"error storing payment"}`, http.StatusInternalServerError)
			return
		}

		http.Error(w,
			fmt.Sprintf(`{"reason":"payment added", "id":"%s", "links":{"self":"%s%s/%s"}}`,
				id, r.Host, r.URL.Path, id), http.StatusOK)
	}
}

// DeletePayment handler
func DeletePayment(store server.Store) (handle server.Handle) {
	return func(w http.ResponseWriter, r *http.Request, p server.Params) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		id := p.ByName("paymentId")
		key := []byte(id)

		store.Get(key,
			func(value []byte, err error) {
				if err != nil {
					http.Error(w, `{"reason":"error fetching payment"}`, http.StatusInternalServerError)
					return
				}

				if value == nil {
					http.Error(w, `{"reason":"payment not found"}`, http.StatusNotFound)
					return
				}

				if err = store.Delete(key); err != nil {
					http.Error(w, `{"reason":"error deleting payment"}`, http.StatusInternalServerError)
					return
				}

				http.Error(w, `{"reason":"payment deleted"}`, http.StatusOK)
				return
			},
		)
	}
}

// GetPaymentByID handler
func GetPaymentByID(store server.Store) (handle server.Handle) {
	return func(w http.ResponseWriter, r *http.Request, p server.Params) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		id := p.ByName("paymentId")
		key := []byte(id)

		store.Get(key,
			func(value []byte, err error) {
				if err != nil {
					http.Error(w, `{"reason":"error fetching payment"}`, http.StatusInternalServerError)
					return
				}

				if value == nil {
					http.Error(w, `{"reason":"payment not found"}`, http.StatusNotFound)
					return
				}

				w.WriteHeader(http.StatusOK)
				w.Write(value)
				return
			},
		)
	}

}

// UpdatePayment handler
func UpdatePayment(store server.Store) (handle server.Handle) {
	return func(w http.ResponseWriter, r *http.Request, p server.Params) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		var err error
		var buf []byte

		if buf, err = ioutil.ReadAll(r.Body); err != nil || len(buf) == 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		payment := &models.Payment{}
		if err = json.Unmarshal(buf, payment); err != nil {
			http.Error(w, `{"reason":"invalid payment"}`, http.StatusBadRequest)
			return
		}

		if !payment.IsValid() {
			http.Error(w, `{"reason":"invalid payment"}`, http.StatusBadRequest)
			return
		}

		buf, err = json.Marshal(payment)
		if !payment.IsValid() {
			http.Error(w, `{"reason":"could not serialize payment"}`, http.StatusInternalServerError)
			return
		}

		id := p.ByName("paymentId")
		key := []byte(id)
		found := false

		store.Get(key,
			func(value []byte, err error) {
				if err != nil {
					http.Error(w, `{"reason":"error fetching payment"}`, http.StatusInternalServerError)
					found = false
					return
				}
				if value == nil {
					http.Error(w, `{"reason":"payment not found"}`, http.StatusNotFound)
					found = false
					return
				}

				found = true
			},
		)

		if !found {
			return
		}

		if err = store.Set(key, buf); err != nil {
			http.Error(w, `{"reason":"error storing payment"}`, http.StatusInternalServerError)
			return
		}

		http.Error(w,
			fmt.Sprintf(`{"reason":"payment updated", "id":"%s", "links":{"self":"%s%s"}}`,
				id, r.Host, r.URL.Path), http.StatusOK)

	}
}
