FROM --platform=$BUILDPLATFORM golang:1.21-bookworm AS builder
WORKDIR /state-manager
ARG TARGETOS
ARG TARGETARCH
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o state-manager backend/state-manager/cmd/main.go

# FROM --platform=$TARGETPLATFORM gcr.io/distroless/base-debian12:nonroot AS runner
FROM --platform=$TARGETPLATFORM debian:bullseye-slim AS runner
COPY --from=builder /state-manager/state-manager /state-manager
COPY backend/state-manager/cmd/emqxsl-ca.pem /emqxsl-ca.pem
RUN apt update
RUN apt install -y curl
ENTRYPOINT ["/state-manager"]