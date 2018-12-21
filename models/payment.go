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

type Payment struct {
	ID             string      `json:"id"`
	OrganisationID string      `json:"organisation_id"`
	Type           string      `json:"type"`
	Version        int         `json:"version"`
	Attributes     *Attributes `json:"attributes"`
}

// IsValid returns true if payment is valid
func (p *Payment) IsValid() (valid bool) {
	switch {
	case p == nil:
		return false
	case p.OrganisationID == "":
		return false
	case p.Type == "":
		return false
	case !p.Attributes.IsValid():
		return false
	default:
		return true
	}
}
