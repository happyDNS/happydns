// This file is part of the happyDomain (R) project.
// Copyright (c) 2020-2024 happyDomain
// Authors: David Dernoncourt, et al.
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

package providers // import "git.happydns.org/happyDomain/providers"

import (
	"github.com/StackExchange/dnscontrol/v4/providers"
	_ "github.com/StackExchange/dnscontrol/v4/providers/autodns"
)

type AutoDNSAPI struct {
	Username string `json:"username,omitempty" happydomain:"label=Username,placeholder=autodns.service-account@example.com,required,description=Your AutoDNS user name."`
	Password string `json:"password,omitempty" happydomain:"label=Password,placeholder=xxxxxxxx,required,description=Your AutoDNS password."`
	Context  string `json:"context,omitempty" happydomain:"label=Context,placeholder=33004,description=Your AutoDNS context."`
}

func (s *AutoDNSAPI) NewDNSServiceProvider() (providers.DNSServiceProvider, error) {
	config := map[string]string{
		"username": s.Username,
		"password": s.Password,
		"context":  s.Context,
	}
	return providers.CreateDNSProvider(s.DNSControlName(), config, nil)
}

func (s *AutoDNSAPI) DNSControlName() string {
	return "AUTODNS"
}

func init() {
	RegisterProvider(func() Provider {
		return &AutoDNSAPI{}
	}, ProviderInfos{
		Name:        "AutoDNS / InterNetX",
		Description: "German hosting provider.",
	})
}
