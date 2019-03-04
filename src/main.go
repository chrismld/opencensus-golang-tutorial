package main

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/exporter/zipkin"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"

	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
)

func registerPrometheus() *prometheus.Exporter {
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "golangsvc",
	})
	if err != nil {
		log.Fatalf("Failed to create Prometheus exporter: %v", err)
	}
	view.RegisterExporter(pe)
	return pe
}

func registerZipkin(){
	localEndpoint, err := openzipkin.NewEndpoint("golangsvc", "192.168.1.61:8080") 
 	if err != nil { 
		log.Fatalf("Failed to create Zipkin exporter: %v", err)
 	} 
 	reporter := zipkinHTTP.NewReporter("http://localhost:9411/api/v2/spans") 
 	exporter := zipkin.NewExporter(reporter, localEndpoint) 
	trace.RegisterExporter(exporter) 
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

func main() {
	pe := registerPrometheus()
	registerZipkin()

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
	database(r)
	serviceb(r)

	res := strings.Repeat("o", rand.Intn(99971)+1)
	time.Sleep(time.Duration(rand.Intn(977)+1) * time.Millisecond)
	w.Write([]byte("Hello, w" + res + "rld!"))
}

func database(r *http.Request) {
	cache(r)
	_, span := trace.StartSpan(r.Context(), "database")
	defer span.End()
	time.Sleep(time.Duration(rand.Intn(977)+300) * time.Millisecond)	
}

func cache(r *http.Request) {
	_, span := trace.StartSpan(r.Context(), "cache")
	defer span.End()
	time.Sleep(time.Duration(rand.Intn(100)+1) * time.Millisecond)
}

func serviceb(r *http.Request) {
	_, span := trace.StartSpan(r.Context(), "serviceb")
	defer span.End()
	time.Sleep(time.Duration(rand.Intn(800)+200) * time.Millisecond)
	servicec(r)
}

func servicec(r *http.Request) {
	_, span := trace.StartSpan(r.Context(), "servicec")
	defer span.End()
	time.Sleep(time.Duration(rand.Intn(700)+100) * time.Millisecond)
}