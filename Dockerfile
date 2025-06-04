FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN go build -o plan-analyzer ./cmd/plan-analyzer

FROM gcr.io/distroless/base
COPY --from=build /app/plan-analyzer /plan-analyzer
ENTRYPOINT ["/plan-analyzer"]
