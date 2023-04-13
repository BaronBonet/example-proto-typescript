FROM golang:1.20

WORKDIR /app

COPY --from=bufbuild/buf:1.10.0 /usr/local/bin/buf /usr/local/bin/buf

COPY dummy_backend/ /app/dummy_backend/
COPY Makefile /app/
COPY buf.work.yaml /app/
COPY ./proto/ /app/proto/

RUN make generate-go

RUN cd dummy_backend && go mod download && go build -o dummy_backend

CMD ["/app/dummy_backend/dummy_backend"]