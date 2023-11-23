FROM --platform=$BUILDPLATFORM golang:1.21-bookworm AS builder
WORKDIR /state-manager
ARG TARGETOS
ARG TARGETARCH
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o state-manager backend/state-manager/cmd/main.go

FROM --platform=$TARGETPLATFORM gcr.io/distroless/base-debian12:nonroot AS runner
COPY --from=builder /state-manager/state-manager /state-manager
ENTRYPOINT ["/state-manager"]