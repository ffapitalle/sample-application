#
# Build stage
#
FROM golang:1.14-alpine as compiler
ARG github_user
ARG github_token
ENV GOPRIVATE "github.com/pedidosya"
WORKDIR /go/src/github.com/pedidosya/@project_name@
COPY . .
RUN apk add make git bash \
  && echo "machine github.com login ${github_user} password ${github_token}" > /root/.netrc \
  && make build
#
# Run stage
#
FROM alpine
COPY --from=compiler /go/src /go/src
RUN apk add --no-cache tzdata
WORKDIR /go/src/github.com/pedidosya/@project_name@
CMD ./service
EXPOSE @service-port@
