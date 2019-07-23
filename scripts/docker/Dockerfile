FROM helderfarias/go:1.12-stretch as builder

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
