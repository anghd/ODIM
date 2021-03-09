//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

package model

// Zone is the redfish Zone model according to the 2020.3 release
type Zone struct {
	ODataContext          string       `json:"@odata.context,omitempty"`
	ODataEtag             string       `json:"@odata.etag,omitempty"`
	ODataID               string       `json:"@odata.id"`
	ODataType             string       `json:"@odata.type"`
	Actions               *OemActions  `json:"Actions,omitempty"`
	Description           string       `json:"Description,omitempty"`
	ID                    string       `json:"Id"`
	Links                 *Links       `json:"Links,omitempty"`
	Name                  string       `json:"Name"`
	Oem                   interface{}  `json:"Oem,omitempty"`
	Status                *Status      `json:"Status,omitempty"`
	DefaultRoutingEnabled bool         `json:"DefaultRoutingEnabled,omitempty"`
	ExternalAccessibility string       `json:"ExternalAccessibility,omitempty"`
	Identifiers           []Identifier `json:"Identifiers,omitempty"`
	ZoneType              string
}

// Identifiers is a redfish model under Zone
type Identifiers struct {
	DurableName       string `json:"DurableName,omitempty"`
	DurableNameFormat string `json:"DurableNameFormat,omitempty"`
}
