# build backend
FROM golang:1.20-alpine AS backend
RUN apk --no-cache --no-progress add --virtual \
  build-deps \
  build-base \
  git

WORKDIR /navis
COPY . .
ENV GOPROXY=https://goproxy.cn
RUN make build TAGS='go_json'

FROM alpine:3.17
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app/navis
COPY --from=backend /navis/release/navis .
COPY --from=backend /navis/app.sample.yaml.sample app.yaml

VOLUME ["/app/navis/custom"]
EXPOSE 8012
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD ps -ef | grep navis || exit 1
ENTRYPOINT ["/app/navis/navis", "serv"]
