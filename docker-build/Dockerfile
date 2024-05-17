# syntax=docker/dockerfile:1

FROM ubuntu:22.04
#COPY . /app

RUN apt update

# Set container timezone / local
RUN apt install -y tzdata
RUN ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime
RUN dpkg-reconfigure -f noninteractive tzdata
ENV TZ="America/New_York"

# Install required commands
RUN apt install -y wget git

# Install go
RUN wget https://go.dev/dl/go1.22.1.linux-arm64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.1.linux-arm64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin

# Install openalpr libraries
RUN apt install -y libopencv-dev libtesseract-dev git cmake build-essential libleptonica-dev liblog4cplus-dev libcurl3-dev

# https://github.com/openalpr/openalpr/wiki/Build-instructions-for-Ubuntu-Linux
RUN git clone https://github.com/openalpr/openalpr.git
RUN mkdir -p openalpr/src/build
RUN cd openalpr/src/build/ && cmake -DCMAKE_INSTALL_PREFIX:PATH=/usr -DCMAKE_INSTALL_SYSCONFDIR:PATH=/etc .. && make && make install
# Add and build viam-alpr module
RUN cd openalpr/src/bindings/go/ && git clone https://github.com/viam-soleng/viam-alpr.git && cd viam-alpr && /usr/local/go/bin/go build -o bin/viam-alpr cmd/module/cmd.go