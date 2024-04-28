FROM --platform=$BUILDPLATFORM golang:1.21-bookworm AS builder
WORKDIR /proxy
ARG TARGETOS
ARG TARGETARCH
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o proxy backend/proxy/main.go

FROM --platform=$TARGETPLATFORM gcr.io/distroless/base-debian12:nonroot AS runner
# FROM --platform=$TARGETPLATFORM debian:bullseye-slim AS runner
COPY --from=builder /proxy/proxy /proxy
ENTRYPOINT ["/proxy"]