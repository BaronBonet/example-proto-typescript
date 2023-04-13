import { createPromiseClient } from "@bufbuild/connect";
import { createGrpcTransport } from "@bufbuild/connect-node";
import { ExampleService } from "./generated/example/v1/example_service_connect";
import { GetExampleMessageRequest } from "./generated/example/v1/example_domain_pb";

const transport = createGrpcTransport({
    httpVersion: "2",
    baseUrl: "http://dummy_backend:8080"
});

async function main() {
    const client = createPromiseClient(ExampleService, transport);
    const request = new GetExampleMessageRequest();
    request.ranNumber = 42;

    const response = await client.getExampleMessage(request);
    console.log(response);
}

void main();
