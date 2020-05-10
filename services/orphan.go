// Copyright or © or Copr. happyDNS (2020)
//
// contact@happydns.org
//
// This software is a computer program whose purpose is to provide a modern
// interface to interact with DNS systems.
//
// This software is governed by the CeCILL license under French law and abiding
// by the rules of distribution of free software.  You can use, modify and/or
// redistribute the software under the terms of the CeCILL license as
// circulated by CEA, CNRS and INRIA at the following URL
// "http://www.cecill.info".
//
// As a counterpart to the access to the source code and rights to copy, modify
// and redistribute granted by the license, users are provided only with a
// limited warranty and the software's author, the holder of the economic
// rights, and the successive licensors have only limited liability.
//
// In this respect, the user's attention is drawn to the risks associated with
// loading, using, modifying and/or developing or reproducing the software by
// the user in light of its specific status of free software, that may mean
// that it is complicated to manipulate, and that also therefore means that it
// is reserved for developers and experienced professionals having in-depth
// computer knowledge. Users are therefore encouraged to load and test the
// software's suitability as regards their requirements in conditions enabling
// the security of their systems and/or data to be ensured and, more generally,
// to use and operate it in the same conditions as regards security.
//
// The fact that you are presently reading this means that you have had
// knowledge of the CeCILL license and that you accept its terms.

package svcs

import (
	"fmt"
	"strings"

	"github.com/miekg/dns"

	"git.happydns.org/happydns/model"
)

type Orphan struct {
	RR dns.RR
}

func (s *Orphan) GetNbResources() int {
	return 1
}

func (s *Orphan) GenComment(origin string) string {
	return fmt.Sprintf("%s", s.RR.String()[strings.LastIndex(s.RR.Header().String(), "\tIN\t")+4:])
}

func (s *Orphan) GenRRs(domain string, ttl uint32) (rrs []dns.RR) {
	rrs = append(rrs, s.RR)

	return
}

func init() {
	RegisterService(
		"git.happydns.org/happydns/services/Orphan",
		func() happydns.Service {
			return &Orphan{}
		},
		nil,
		ServiceInfos{
			Name:        "Orphan Record",
			Description: "",
			Categories:  []string{},
		},
		99999999,
	)
}
