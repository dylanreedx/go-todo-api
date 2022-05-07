FROM golang:1.16-buster AS build
WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /api

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /api /api

COPY --from=build /app/.env .

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT [ "/api" ]
