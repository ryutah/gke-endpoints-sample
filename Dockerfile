FROM golang:1.10-alpine

COPY ./ /go/src/github.com/ryutah/gke-endpoints-sample

RUN go install github.com/ryutah/gke-endpoints-sample

CMD gke-endpoints-sample -stderrthreshold=INFO
