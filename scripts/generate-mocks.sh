#!/bin/sh

docker run --rm \
  --entrypoint=/bin/sh \
  -v "$PWD":/src \
  -w /src \
  golang:1.24-alpine \
  -c 'apk add --no-cache git && go install github.com/vektra/mockery/v2@latest && mockery --dir=internal/domain/interfaces --output=mocks/internal_/domain/interfaces --all'