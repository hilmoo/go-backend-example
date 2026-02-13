FROM gcr.io/distroless/static-debian13

ARG TARGETPLATFORM

WORKDIR /app

COPY $TARGETPLATFORM/go-backend-example /usr/bin/
COPY migration /app/migration

USER 1000:1000
VOLUME ["/app/data"]

ENTRYPOINT ["/usr/bin/go-backend-example"]