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
	"net/http"

	"github.com/brunotm/f3api/server"
)

// AddAllRoutes api routes to the given server
func AddAllRoutes(server *server.Server) {
	server.AddStoreHandler(http.MethodPost, `/api/v1/payment`, AddPayment)
	server.AddStoreHandler(http.MethodGet, `/api/v1/payment/:paymentId`, GetPaymentByID)
	server.AddStoreHandler(http.MethodPut, `/api/v1/payment/:paymentId`, UpdatePayment)
	server.AddStoreHandler(http.MethodDelete, `/api/v1/payment/:paymentId`, DeletePayment)
	server.AddStoreHandler(http.MethodGet, `/api/v1/payments`, GetAllPayments)
}
