FROM golang:1.20
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git
RUN git clone https://github.com/MaheshRKumawat/piper
COPY ./bash.sh .
RUN chmod +x bash.sh
CMD ["./bash.sh"]