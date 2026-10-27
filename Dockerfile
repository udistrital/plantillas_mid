FROM gcr.io/distroless/base-debian13

WORKDIR /

COPY main main
COPY conf/app.conf conf/app.conf

ENTRYPOINT ["/main"]
