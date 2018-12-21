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

type Charge struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// IsValid returns true if payment charge is valid
func (c *Charge) IsValid() (valid bool) {
	switch {
	case c == nil:
		return false
	case c.Amount == "":
		return false
	case c.Currency == "":
		return false
	default:
		return true
	}
}

type ChargesInformation struct {
	BearerCode              string    `json:"bearer_code"`
	ReceiverChargesAmount   string    `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string    `json:"receiver_charges_currency"`
	SenderCharges           []*Charge `json:"sender_charges"`
}

// IsValid returns true if charges information is valid
func (c *ChargesInformation) IsValid() (valid bool) {
	switch {
	case c == nil:
		return false
	case c.BearerCode == "":
		return false
	case c.ReceiverChargesAmount == "":
		return false
	case c.ReceiverChargesCurrency == "":
		return false
	case len(c.SenderCharges) == 0:
		return false
	}

	for i := range c.SenderCharges {
		if !c.SenderCharges[i].IsValid() {
			return false
		}
	}

	return true
}
