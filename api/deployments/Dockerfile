FROM golang:1.20 AS build-stage

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o api .

FROM alpine

ARG PORT
ARG DB_URL

ENV IS_PROD=true
ENV PORT=${PORT}
ENV DATABASE_URL=${DB_URL}

COPY --from=build-stage /app/api api
COPY --from=build-stage /app/scripts/migrations /migrations

ENTRYPOINT ./api
