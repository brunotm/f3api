package models

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

type Party struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

// IsValid returns true if payment attributes parties is valid
func (p *Party) IsValid() (valid bool) {
	switch {
	case p == nil:
		return false
	case p.AccountNumber == "":
		return false
	case p.BankID == "":
		return false
	case p.BankIDCode == "":
		return false
	default:
		return true
	}
}
