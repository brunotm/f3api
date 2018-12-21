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

type Attributes struct {
	Amount               string              `json:"amount"`
	Currency             string              `json:"currency"`
	EndToEndReference    string              `json:"end_to_end_reference"`
	NumericReference     string              `json:"numeric_reference"`
	PaymentID            string              `json:"payment_id"`
	PaymentPurpose       string              `json:"payment_purpose"`
	PaymentScheme        string              `json:"payment_scheme"`
	PaymentType          string              `json:"payment_type"`
	ProcessingDate       string              `json:"processing_date"`
	Reference            string              `json:"reference"`
	SchemePaymentSubType string              `json:"scheme_payment_sub_type"`
	SchemePaymentType    string              `json:"scheme_payment_type"`
	FX                   *FX                 `json:"fx"`
	SponsorParty         *Party              `json:"sponsor_party"`
	BeneficiaryParty     *Party              `json:"beneficiary_party"`
	DebtorParty          *Party              `json:"debtor_party"`
	ChargesInformation   *ChargesInformation `json:"charges_information"`
}

// IsValid returns true if payment attributes is valid
func (a *Attributes) IsValid() (valid bool) {
	switch {
	case a == nil:
		return false
	case a.Amount == "":
		return false
	case a.Currency == "":
		return false
	case a.EndToEndReference == "":
		return false
	case a.NumericReference == "":
		return false
	case a.PaymentID == "":
		return false
	case a.PaymentPurpose == "":
		return false
	case a.PaymentScheme == "":
		return false
	case a.PaymentType == "":
		return false
	case a.ProcessingDate == "":
		return false
	case a.Reference == "":
		return false
	case a.Reference == "":
		return false
	case a.SchemePaymentType == "":
		return false
	case a.SchemePaymentSubType == "":
		return false
	}

	if !a.FX.IsValid() {
		return false
	}

	if !a.SponsorParty.IsValid() {
		return false
	}

	if !a.BeneficiaryParty.IsValid() {
		return false
	}

	if !a.DebtorParty.IsValid() {
		return false
	}

	if !a.ChargesInformation.IsValid() {
		return false
	}

	return true
}
