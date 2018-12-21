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

type FX struct {
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}

// IsValid returns true if fx information is valid
func (f *FX) IsValid() (valid bool) {
	switch {
	case f == nil:
		return false
	case f.ContractReference == "":
		return false
	case f.ExchangeRate == "":
		return false
	case f.OriginalAmount == "":
		return false
	case f.OriginalCurrency == "":
		return false
	default:
		return true
	}
}
