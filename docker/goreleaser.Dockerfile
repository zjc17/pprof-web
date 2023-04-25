FROM alpine:3.16
RUN apk add --update graphviz
RUN rm -rf /var/cache/apk/*
WORKDIR /
COPY pprof-web /bin/pprof-web
ENTRYPOINT ["pprof-web"]