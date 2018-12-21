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
	"net/http"

	"github.com/brunotm/f3api/models"
	"github.com/brunotm/f3api/server"
)

// GetAllPayments handler
func GetAllPayments(store server.Store) (handle server.Handle) {
	return func(w http.ResponseWriter, r *http.Request, p server.Params) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		payments := &Payments{}
		payments.Links.Self = fmt.Sprintf(`%s%s`, r.Host, r.URL.Path)

		err := store.Iter(nil, nil,
			func(key, value []byte) (err error) {
				payment := &models.Payment{}
				if err = json.Unmarshal(value, payment); err != nil || !payment.IsValid() {
					return err
				}
				payments.Data = append(payments.Data, payment)
				return nil
			},
		)

		if err != nil {
			http.Error(w, `{"reason":"error fetching payments"}`, http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(payments)
		if err != nil {
			http.Error(w, `{"reason":"error fetching payments"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)

	}
}

// Payments response
type Payments struct {
	Data  []*models.Payment `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
