FROM golang:1.23rc2 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 GOOS=linux go build -o go-server


FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/go-server /app/app.db /

COPY ./static/ ./static/

ENV PORT=8000

EXPOSE 8000


ENTRYPOINT ["/go-server"]