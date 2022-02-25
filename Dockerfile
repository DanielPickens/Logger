FROM FROM golang:1.7.1

ARG CI


ADD . /build/docker-logger
WORKDIR /build/docker-logger

RUN cd app && go test -v -mod=vendor -covermode=count -coverprofile=/profile.cov ./...

RUN golangci-lint run --out-format=tab --tests=false ./...

RUN \
    revison=$(/script/git-rev.sh) && \
    echo "revision=${revison}" && \
    go build -mod=vendor -o docker-logger -ldflags "-X main.revision=$revison -s -w" ./app

# submit coverage to coverals if COVERALLS_TOKEN in env
RUN if [ -z "$COVERALLS_TOKEN" ] ; then echo "coverall not enabled" ; \
    else goveralls -coverprofile=/profile.cov -service=travis-ci -repotoken $COVERALLS_TOKEN || echo "coverall failed!"; fi


FROM alpine:3.9

RUN apk add --update --no-cache tzdata

COPY --from=build /build/docker-logger /srv/


WORKDIR /srv

VOLUME ["/srv/logs"]
CMD ["/srv/docker-logger"]
