FROM public.ecr.aws/docker/library/golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod files from both main module and proto submodule
COPY go.mod go.sum ./
COPY proto/go.mod proto/go.sum ./proto/

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/app cmd/main.go

FROM public.ecr.aws/docker/library/golang:1.24-alpine

COPY --from=builder /app/bin ./bin

ENTRYPOINT ["bin/app"]