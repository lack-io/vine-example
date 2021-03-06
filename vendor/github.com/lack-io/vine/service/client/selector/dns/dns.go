// Copyright 2020 lack
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package dns provides a dns SRV selector
package dns

import (
	"fmt"
	"net"
	"strconv"

	regpb "github.com/lack-io/vine/proto/registry"
	"github.com/lack-io/vine/service/client/selector"
)

type dnsSelector struct {
	options selector.Options
	domain  string
}

var (
	DefaultDomain = "local"
)

func (d *dnsSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&d.options)
	}
	return nil
}

func (d *dnsSelector) Options() selector.Options {
	return d.options
}

func (d *dnsSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	var srv []*net.SRV

	// check if its host:port
	host, port, err := net.SplitHostPort(service)
	// not host:port
	if err != nil {
		// lookup the SRV record
		_, srvs, err := net.LookupSRV(service, "tcp", d.domain)
		if err != nil {
			return nil, err
		}
		// set SRV records
		srv = srvs
		// got host:port
	} else {
		p, _ := strconv.Atoi(port)

		// lookup the A record
		ips, err := net.LookupHost(host)
		if err != nil {
			return nil, err
		}

		// create SRV records
		for _, ip := range ips {
			srv = append(srv, &net.SRV{
				Target: ip,
				Port:   uint16(p),
			})
		}
	}

	nodes := make([]*regpb.Node, 0, len(srv))
	for _, node := range srv {
		nodes = append(nodes, &regpb.Node{
			Id:      node.Target,
			Address: fmt.Sprintf("%s:%d", node.Target, node.Port),
		})
	}

	services := []*regpb.Service{
		{
			Name:  service,
			Nodes: nodes,
		},
	}

	sopts := selector.SelectOptions{
		Strategy: d.options.Strategy,
	}

	for _, opt := range opts {
		opt(&sopts)
	}

	// apply the filters
	for _, filter := range sopts.Filters {
		services = filter(services)
	}

	// if there's nothing left, return
	if len(services) == 0 {
		return nil, selector.ErrNoneAvailable
	}

	return sopts.Strategy(services), nil
}

func (d *dnsSelector) Mark(service string, node *regpb.Node, err error) {}

func (d *dnsSelector) Reset(service string) {}

func (d *dnsSelector) Close() error {
	return nil
}

func (d *dnsSelector) String() string {
	return "dns"
}

func NewSelector(opts ...selector.Option) selector.Selector {
	options := selector.Options{
		Strategy: selector.Random,
	}

	for _, o := range opts {
		o(&options)
	}

	return &dnsSelector{options: options, domain: DefaultDomain}
}
