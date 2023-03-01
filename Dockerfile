FROM golang:1.19-alpine AS stage
WORKDIR /app
COPY [ "go.mod", "go.sum", "./" ]
RUN go mod download 
COPY . .
RUN go build -o dist/httpserver ./cmd/httpserver

FROM alpine:3.16
WORKDIR /app
COPY --from=stage /app/dist/httpserver /app/
COPY configs/httpserver.yml /app/
CMD [ "/app/httpserver", "-c", "/app/httpserver.yml" ]