FROM golang:latest as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN  CGO_ENABLED=0 go build -o app router.go
# 运行阶段指定scratch作为基础镜像
FROM alpine as runner

WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]