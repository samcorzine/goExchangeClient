FROM scratch
ADD GoExchangeClient /GoExchangeClient
ENV EXCHANGE_URL "http://172.17.0.2:8080/order"
ENTRYPOINT ["/GoExchangeClient"]
