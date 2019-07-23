FROM golang:1.12-stretch as builder
LABEL maintainer "Helder Farias <helderfarias@gmail.com>"

RUN apt-get update && apt-get install -y xz-utils && rm -rf /var/lib/apt/lists/*
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

WORKDIR /root
ADD target/gcs_linux gcs_linux
ADD scripts/docker/entrypoint.sh entrypoint.sh
RUN strip --strip-unneeded /root/gcs_linux
RUN upx /root/gcs_linux

FROM helderfarias/alpine:3.7
COPY --from=builder /root/gcs_linux /bin/gcs
COPY --from=builder /root/entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/gcs
RUN chmod +x /bin/entrypoint.sh
EXPOSE 3001

CMD ["/bin/entrypoint.sh"]
