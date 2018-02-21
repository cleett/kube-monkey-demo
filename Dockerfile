FROM golang:1.8

COPY . /go/src/github.com/cleett/kube-monkey-demo
RUN go install github.com/cleett/kube-monkey-demo
