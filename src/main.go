package main

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

func main() {
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "golangsvc",
	})
	if err != nil {
		log.Fatalf("Failed to create Prometheus exporter: %v", err)
	}
	view.RegisterExporter(pe)

	mux := http.NewServeMux()
	mux.HandleFunc("/list", list)

	mux.Handle("/metrics", pe)

	h := &ochttp.Handler{Handler: mux}
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		log.Fatal("Failed to register ochttp.DefaultServerViews")
	}

	log.Printf("Server listening! ...")
	log.Fatal(http.ListenAndServe(":8080", h))
}

func list(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	res := strings.Repeat("o", rand.Intn(99971)+1)
	time.Sleep(time.Duration(rand.Intn(977)+1) * time.Millisecond)
	w.Write([]byte("Hello, w" + res + "rld!"))
}
