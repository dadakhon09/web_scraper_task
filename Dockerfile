# workspace (GOPATH) configured at /go
FROM golang:1.16 as builder


RUN mkdir -p $GOPATH/src/web_scraper_task
WORKDIR $GOPATH/src/web_scraper_task

COPY . ./

RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/web_scraper_task /


FROM alpine

EXPOSE 8082

COPY --from=builder web_scraper_task .

ENTRYPOINT ["/web_scraper_task"]