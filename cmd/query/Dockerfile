FROM harbor.shopeemobile.com/shopee/peg-compiler:latest

WORKDIR /workspace

RUN mkdir -p /root/go/src/github.com/jaegertracing/jaeger/

ADD . /root/go/src/github.com/jaegertracing/jaeger/

RUN cd /root/go/src/github.com/jaegertracing/jaeger/ && dep ensure -v -vendor-only && make build-query && mv cmd/query/query- /workspace/jaeger-query

