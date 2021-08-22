FROM redis:6.2-buster

RUN mkdir /go_server
WORKDIR /go_server
RUN  apt-get update \
  && apt-get install -y wget \
  && rm -rf /var/lib/apt/lists/*
RUN wget https://github.com/K-ran/diceChessDiceGoServer/releases/download/latest/dcd_server
RUN chmod +x dcd_server
RUN echo "redis-server --daemonize yes" >> start.sh
RUN echo "./dcd_server" >> start.sh
RUN chmod +x ./start.sh
EXPOSE 8081
CMD [ "bash", "./start.sh" ]