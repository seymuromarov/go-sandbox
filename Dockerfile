FROM ubuntu:24.04

RUN apt-get update \
    && apt-get install -y python3.10 \
    && apt-get install -y gcc \
    && apt-get install -y g++ \
    && apt-get install -y default-jre \
    && apt-get install -y openjdk-21-jdk \
    && apt-get install -y curl \
    && apt-get clean;

# Install Go 1.18
RUN curl -OL https://golang.org/dl/go1.18.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz \
    && rm go1.18.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH=$PATH:/usr/local/go/bin
# Install Node.js 18
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
    && apt-get install -y nodejs

RUN npm i -g npm 

RUN npm i -g typescript