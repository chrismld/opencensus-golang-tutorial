# OpenCensus Getting Started with Golang

This is a simple example of an app written in Go that uses OpenCensus to collect metrics and traces that you can later export to other tools like Prometheus.

## Install Dependencies

In order to run the app locally, you need to install the OpenCensus dependencies first. Just run this command:

```
go get -u -v go.opencensus.io/...
```

## Start a Prometheus container

To start collecting metrics, you need to have Prometheus running. The easiest way to do it is by using Docker and running the following command in the root folder of this repository (so it uses the prometheus.yaml config file). Replace '192.168.1.60' with your local private IP:

```
docker run --name prometheus --add-host="localhost:192.168.1.60" -d -p 9090:9090 -v ${PWD}/prometheus.yaml:/etc/prometheus/prometheus.yml prom/prometheus
```

Open your browser and go to [http://localhost:9090/](https://localhost:9090/)

## Start a Zipkin container

To collect the traces of the applicacion, you need to have Zipkin running. Again, the easiest way to do it is by using Docker and running the following command:

```
docker run -d -p 9411:9411 openzipkin/zipkin
```

Open your browser and go to [http://localhost:9411/](https://localhost:9411/)

## Start the Go API

Let's start the API that will generate metrics and traces using OpenCensus. Run the following command:

```
go run src/main.go
```

Open your browser and go to [http://localhost:8080/list](https://localhost:8080/list)

## Generate Load

Let's use the [ApacheBench (ab)](https://httpd.apache.org/docs/2.4/programs/ab.html) tool to generate some load to the Go API. Again, the easiest way to do it is by using Docker. Replace '192.168.1.60' with your local private IP, and run the following command:

```
docker run --rm --add-host="localhost:192.168.1.60" jordi/ab -k -c 10 -n 15000 http://localhost:8080/list 
```