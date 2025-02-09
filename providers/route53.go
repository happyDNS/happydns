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
	_ "github.com/StackExchange/dnscontrol/v4/providers/route53"
)

type Route53API struct {
	DelegationSet string `json:"delegation_set,omitempty" happydomain:"label=Delegation Set ID,placeholder=xxxxxxxx,description=Optional delegation set ID."`
	KeyId         string `json:"key_id,omitempty" happydomain:"label=AWS key,placeholder=xxxxxxxx,required,description=Your AWS key."`
	SecretKey     string `json:"secret_key,omitempty" happydomain:"label=AWS secret key,placeholder=xxxxxxxx,required,description=Your AWS secret key."`
	Token         string `json:"token,omitempty" happydomain:"label=Token,placeholder=xxxxxxxx,description=Optional STS token."`
}

func (s *Route53API) NewDNSServiceProvider() (providers.DNSServiceProvider, error) {
	config := map[string]string{
		"DelegationSet": s.DelegationSet,
		"KeyId":         s.KeyId,
		"SecretKey":     s.SecretKey,
		"Token":         s.Token,
	}
	return providers.CreateDNSProvider(s.DNSControlName(), config, nil)
}

func (s *Route53API) DNSControlName() string {
	return "ROUTE53"
}

func init() {
	RegisterProvider(func() Provider {
		return &Route53API{}
	}, ProviderInfos{
		Name:        "AWS Route 53",
		Description: "Amazon Web Services (AWS) is a subsidiary of Amazon that provides on-demand cloud computing platforms and APIs. Route 53 is their DNS service.",
	})
}
