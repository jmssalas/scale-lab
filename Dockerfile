FROM golang:latest AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG SERVICE
RUN CGO_ENABLED=0 go build -o /out/app ./cmd/${SERVICE}

FROM gcr.io/distroless/static-debian13
COPY --from=build /out/app /app
ENTRYPOINT ["/app"]
