FROM node:22-alpine AS tailwind_builder
WORKDIR /app
COPY . .
RUN npm install
RUN npx tailwindcss -i /app/static/css/style.css -o /app/static/css/tailwind.css -m

FROM golang:1.24-alpine AS go_modules
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM ghcr.io/a-h/templ:latest AS templ_builder
WORKDIR /app
COPY --chown=65532:65532 . .
RUN ["templ", "generate", "-path", "components/"]

FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY --from=go_modules /app/go.mod /app/go.sum ./
COPY --from=templ_builder /app /app
RUN CGO_ENABLED=0 GOOS=linux go build -o blog

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/blog ./blog
COPY --from=builder /app/posts ./posts
COPY --from=tailwind_builder /app/static ./static
CMD ["/app/blog", "serve"]
