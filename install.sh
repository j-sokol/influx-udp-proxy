#!/bin/bash

SERVICE_NAME=influx-udp-proxy
INSTALL_PATH=/usr/local/bin
SERVICE_PATH=/lib/systemd/system/

set -eu

go build

cp ./$SERVICE_NAME /usr/local/bin

chmod a+x ${INSTALL_PATH}/$SERVICE_NAME

cp ${SERVICE_NAME}.service ${SERVICE_PATH}/

sudo systemctl enable ${SERVICE_NAME}.service
sudo systemctl start ${SERVICE_NAME}.service