.PHONY: generate-go
generate-go:
	@rm -rf dummy_backend/generated/go
	@buf generate --template proto/buf.gen.go.yaml

.PHONY: generate-ts
generate-ts:
	@rm -rf client/generated/ts
	@npx buf generate --template proto/buf.gen.ts.yaml
