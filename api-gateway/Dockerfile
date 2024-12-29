FROM envoyproxy/envoy:v1.31.0

RUN apt-get update && apt-get -q install -y \
    curl
COPY ./api-gateway.yaml /etc/api-gateway.yaml
RUN chmod go+r /etc/api-gateway.yaml
CMD ["/usr/local/bin/envoy", "-c", "/etc/api-gateway.yaml", "--service-cluster", "api-gateway"]
