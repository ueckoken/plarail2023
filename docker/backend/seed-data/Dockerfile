FROM --platform=$BUILDPLATFORM golang:1.21-bookworm AS builder
WORKDIR /seed-data
ARG TARGETOS
ARG TARGETARCH
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o seed-data backend/onetime/seed-data/main.go

FROM --platform=$TARGETPLATFORM gcr.io/distroless/base-debian12:nonroot AS runner
COPY --from=builder /seed-data/seed-data /seed-data
# デフォルトのデータファイルをコピー（オプション）
COPY backend/onetime/seed-data/data /data
ENTRYPOINT ["/seed-data"]
