FROM golang:1.23-alpine AS go-builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -o app container.go

FROM alpine:latest

WORKDIR /app

COPY --from=go-builder /build/app .

RUN addgroup -S ffmpeg && adduser -S -G ffmpeg ffmpeg

RUN apk add --no-cache curl tar xz

RUN curl -L -o /tmp/ffmpeg.tar.xz https://johnvansickle.com/ffmpeg/builds/ffmpeg-git-amd64-static.tar.xz && \
    mkdir -p /opt/ffmpeg && \
    tar -xf /tmp/ffmpeg.tar.xz -C /opt/ffmpeg --strip-components=1 && \
    ln -s /opt/ffmpeg/ffmpeg /usr/local/bin/ffmpeg && \
    ln -s /opt/ffmpeg/ffprobe /usr/local/bin/ffprobe && \
    rm -rf /tmp/ffmpeg.tar.xz

RUN chown -R ffmpeg:ffmpeg /opt/ffmpeg /app && \
    chmod +x /opt/ffmpeg/ffmpeg && \
    chmod +x /opt/ffmpeg/ffprobe

RUN ffmpeg -version && \
    ffprobe -version

USER ffmpeg

EXPOSE 8080

CMD ["/app/app"]
