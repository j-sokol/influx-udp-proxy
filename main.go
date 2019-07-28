package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func generateMetric(value interface{}, metricname, remote string) []byte {
	return []byte(fmt.Sprintf("%s,location=%s value=%f", metricname, remote, value))
}

func main() {
	conn, err := net.ListenUDP(listenProto, &net.UDPAddr{
		Port: listenPort,
		IP:   net.ParseIP(listenIP),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		var metric map[string]interface{}

		message := make([]byte, inputPacketSize)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}
		// Trim input udp string
		data := strings.TrimSpace(string(message[:rlen]))

		json.Unmarshal([]byte(data), &metric)
		fmt.Printf("Metric: %+v\n", metric)

		remoteStr := fmt.Sprint(remote)
		fmt.Printf("received: %s from %s, sending to %s\n", data, remoteStr, influxAddr)

		// value is interface type
		for key, value := range metric {
			// Print the metric line
			fmt.Printf("%s\n", generateMetric(value, key, remoteStr))
			// send the metric to influx
			resp, err := http.Post(influxAddr, "application/txt", bytes.NewBuffer(generateMetric(value, key, remoteStr)))

			if err != nil {
				panic(err)
			}

			fmt.Printf("Response from influx: %+v\n", resp.StatusCode)

		}

	}
}
