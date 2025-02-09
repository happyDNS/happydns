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

package providers // import "git.happydns.org/happyDomain/providers"

import (
	"flag"

	"github.com/StackExchange/dnscontrol/v4/providers"
	_ "github.com/StackExchange/dnscontrol/v4/providers/bind"
)

type BindServer struct {
	Directory  string `json:"directory,omitempty" happydomain:"label=Directory,placeholder=/etc/named/zones/,required,description=Local directory on the same host running happyDomain, containing your zones"`
	Fileformat string `json:"fileformat,omitempty" happydomain:"label=File format,placeholder=%U.zone,description=See format at https://docs.dnscontrol.org/service-providers/providers/bind#filenameformat"`
}

func (s *BindServer) NewDNSServiceProvider() (providers.DNSServiceProvider, error) {
	config := map[string]string{
		"directory": s.Directory,
	}

	if s.Fileformat != "" {
		config["filenameformat"] = s.Fileformat
	}

	return providers.CreateDNSProvider(s.DNSControlName(), config, nil)
}

func (s *BindServer) DNSControlName() string {
	return "BIND"
}

func init() {
	flag.BoolFunc("with-bind-provider", "Enable the BIND provider (not suitable for cloud/shared instance as it'll access the local file system)", func(s string) error {
		RegisterProvider(func() Provider {
			return &BindServer{}
		}, ProviderInfos{
			Name:        "Bind files/RFC 1035",
			Description: "Use zone files saved in the RFC 1035 format.",
		})
		return nil
	})
}
