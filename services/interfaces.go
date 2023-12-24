// This file is part of the happyDomain (R) project.
// Copyright (c) 2020-2024 happyDomain
// Authors: Pierre-Olivier Mercier, et al.
//
// This program is offered under a commercial and under the AGPL license.
// For commercial licensing, contact us at <contact@happydomain.org>.
//
// For AGPL licensing:
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package svcs

import (
	"encoding/json"
	"fmt"

	"git.happydns.org/happyDomain/model"
)

type ServiceRestrictions struct {
	// Alone restricts the service to be the only one for a given subdomain.
	Alone bool `json:"alone,omitempty"`

	// ExclusiveRR restricts the service to be present along with others given types.
	ExclusiveRR []string `json:"exclusive,omitempty"`

	// GLUE allows a service to be present under Leaf, as GLUE record.
	GLUE bool `json:"glue,omitempty"`

	// Leaf restricts the creation of subdomains under this kind of service (blocks NearAlone).
	Leaf bool `json:"leaf,omitempty"`

	// NearAlone allows a service to be present along with Alone restricted services (eg. services that will create sub-subdomain from their given subdomain).
	NearAlone bool `json:"nearAlone,omitempty"`

	// NeedTypes restricts the service to sources that are compatibles with ALL the given types.
	NeedTypes []uint16 `json:"needTypes,omitempty"`

	// RootOnly restricts the service to be present at the root of the domain only.
	RootOnly bool `json:"rootOnly,omitempty"`

	// Single restricts the service to be present only once per subdomain.
	Single bool `json:"single,omitempty"`
}

type ServiceInfos struct {
	Name         string              `json:"name"`
	Type         string              `json:"_svctype"`
	Icon         string              `json:"_svcicon,omitempty"`
	Description  string              `json:"description"`
	Family       string              `json:"family"`
	Categories   []string            `json:"categories"`
	RecordTypes  []uint16            `json:"record_types"`
	Tabs         bool                `json:"tabs,omitempty"`
	Restrictions ServiceRestrictions `json:"restrictions,omitempty"`
}

type serviceCombined struct {
	Service happydns.Service
}

type ServiceNotFoundError struct {
	name string
}

func (err ServiceNotFoundError) Error() string {
	return fmt.Sprintf("Unable to find corresponding service for `%s`.", err.name)
}

// UnmarshalServiceJSON implements the UnmarshalJSON function for the
// encoding/json module.
func UnmarshalServiceJSON(svc *happydns.ServiceCombined, b []byte) (err error) {
	var svcType happydns.ServiceMeta
	err = json.Unmarshal(b, &svcType)
	if err != nil {
		return
	}

	var tsvc happydns.Service
	tsvc, err = FindService(svcType.Type)
	if err != nil {
		return
	}

	mySvc := &serviceCombined{
		tsvc,
	}

	err = json.Unmarshal(b, mySvc)

	svc.Service = tsvc
	svc.ServiceMeta = svcType

	return
}

func init() {
	// Register the UnmarshalServiceJSON variable thats points to the
	// Service's UnmarshalJSON implementation that can't be made in model
	// module due to cyclic dependancy.
	happydns.UnmarshalServiceJSON = UnmarshalServiceJSON
}
