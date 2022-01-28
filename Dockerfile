FROM docker.io/library/golang:1.17.6 AS builder
WORKDIR /go/src/github.com/tnozicka/cassandra-stresser
COPY . .
RUN make build --warn-undefined-variables

FROM docker.io/scylladb/scylla:4.5.3
COPY --from=builder /go/src/github.com/tnozicka/cassandra-stresser/cassandra-stresser /usr/bin/
ENTRYPOINT ["/usr/bin/cassandra-stresser"]
