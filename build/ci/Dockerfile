FROM ubuntu:22.04 as ci

ARG REF_ARG="master"
ARG ANSIBLE_INVENTORY_FOLDER_ARG="/srv/runner/inventory/$REF_ARG"
ARG ANSIBLE_INVENTORY_FILE_ARG="inventory.ec2"

ENV REF=$REF_ARG
ENV ANSIBLE_INVENTORY="$ANSIBLE_INVENTORY_FOLDER_ARG/$ANSIBLE_INVENTORY_FILE_ARG"
ENV ANSIBLE_INVENTORY_FOLDER=$ANSIBLE_INVENTORY_FOLDER_ARG

VOLUME /srv/runner/inventory/

RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt install -y make ansible wget unzip git gcc curl

RUN wget "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip"  && unzip awscli-exe-linux-x86_64.zip && ./aws/install

RUN wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

RUN mkdir -p /srv/newrelic

RUN git clone https://github.com/newrelic/infrastructure-agent.git /srv/newrelic/infrastructure-agent
WORKDIR /srv/newrelic/infrastructure-agent

RUN mkdir ~/.ssh/

COPY ./entrypoint.sh /srv/newrelic/infrastructure-agent/entrypoint.sh

RUN chmod +x /srv/newrelic/infrastructure-agent/entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
