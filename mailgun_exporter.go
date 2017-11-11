// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// A simple example exposing fictional RPC latencies with different types of
// random distributions (uniform, normal, and exponential) as Prometheus
// metrics.
package main

import (
	"flag"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
    "gopkg.in/mailgun/mailgun-go.v1"
	"fmt"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)


func init() {
	// Register the summary and the histogram with Prometheus's default registry.
}

func GetDomains(mg mailgun.Mailgun) ([]mailgun.Domain) {
	_ , domains , err := mg.GetDomains(-1, -1)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return domains
}

func main() {
	flag.Parse()

	mg, error := mailgun.NewMailgunFromEnv()

	if error != nil {
		log.Fatal(error)
	}

	domains := GetDomains(mg)

	if domains == nil {
		log.Error("Error trying to get domains from mailgun")
	}

	log.Info("Number of domains for this account", len(domains))

	//for index,element := range domains {
	//
	//	mg.GetStats()
	//
	//}


	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
