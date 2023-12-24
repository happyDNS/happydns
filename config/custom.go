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

package config // import "git.happydns.org/happyDomain/config"

import (
	"encoding/base64"
	"net/url"
)

type JWTSecretKey []byte

func (i *JWTSecretKey) String() string {
	return base64.StdEncoding.EncodeToString(*i)
}

func (i *JWTSecretKey) Set(value string) error {
	z, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return err
	}

	*i = z
	return nil
}

type URL struct {
	URL *url.URL
}

func (i *URL) String() string {
	if i.URL != nil {
		return i.URL.String()
	} else {
		return ""
	}
}

func (i *URL) Set(value string) error {
	u, err := url.Parse(value)
	if err != nil {
		return err
	}

	i.URL = u
	return nil
}
