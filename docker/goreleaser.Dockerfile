FROM alpine:3.16
RUN rm -rf /var/cache/apk/*
WORKDIR /
COPY pprof-web /bin/pprof-web
ENTRYPOINT ["pprof-web"]