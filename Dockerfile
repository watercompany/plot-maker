FROM golang:1.16
WORKDIR /go/src/plot-maker
COPY . .
RUN make build

FROM ubuntu:latest
ENV DEBIAN_FRONTEND="noninteractive"
RUN apt-get update && apt-get install build-essential git cmake -y
WORKDIR /root/
RUN git clone https://github.com/watercompany/chiapos.git
RUN cd chiapos && \
    mkdir build && \
    cd build && \
    cmake ../ && \
    cmake --build . -- -j 6

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/plot-maker/build/plot-maker plot-maker
COPY --from=1 /root/chiapos/build/ProofOfSpace ProofOfSpace
CMD ["/root/plot-maker", "-bin", "/root/ProofOfSpace", "-json", "args.json", "-d", "/root/final_dir"]
