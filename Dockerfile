FROM golang:1.23rc2 AS build-stage

WORKDIR /app

COPY . ./

RUN go mod download && CGO_ENABLED=1 GOOS=linux go build -o go-server ./cmd/main.go


FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/go-server /app/app.db /

COPY ./static/ ./static/

ENV PORT=8080

EXPOSE 8080

ENTRYPOINT ["/go-server"]