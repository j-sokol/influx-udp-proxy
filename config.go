package main

// Influxdb database specifics
var influxAddr = "http://raspberrypi.local:8086/write?db=test"
var influxProto = "tcp"
var influxDatabase = "test"

// Listen protocol specifics
var listenIP = "0.0.0.0"
var listenPort = 1337
var listenProto = "udp"
var inputPacketSize = 128
