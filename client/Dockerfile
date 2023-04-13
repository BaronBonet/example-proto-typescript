FROM node:19

WORKDIR /app

COPY --from=bufbuild/buf:1.10.0 /usr/local/bin/buf /usr/local/bin/buf

COPY client /app/client/
COPY Makefile /app/
COPY buf.work.yaml /app/
COPY ./proto/ /app/proto

RUN cd client && npm ci

RUN npm install @bufbuild/buf @bufbuild/protoc-gen-es @bufbuild/protobuf @bufbuild/protoc-gen-connect-es @bufbuild/connect
RUN make generate-ts

CMD ["npx","tsx", "client/client.ts"]