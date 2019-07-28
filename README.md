# UDP to influxdb proxy

Little proxy to resend data from UDP to influxdb.

## Requirements

Golang needs to be installed.


# Configuration

- Set up influxdb endpoint and other stuff in `config.go`


# Instalation

Run `install.sh` under root user, this will crete systemd service called `influx-udp-proxy`.