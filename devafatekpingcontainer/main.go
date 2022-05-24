package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/HsmTeknoloji/ping_lib_go/devhsmtekping"
)

//ResultVal return value for request
type ResultVal struct {
	Host    string `json:"host"`
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

func main() {

	var serverPort int
	serverPort, _ = strconv.Atoi(os.Getenv("PORT"))
	http.HandleFunc("/", Index)
	http.ListenAndServe(":"+strconv.Itoa(serverPort), nil)

}

//Index root request path
func Index(w http.ResponseWriter, r *http.Request) {

	var ipAddress string = "192.168.1.1"
	var pingCount int = 5
	var timeOutSecond int = 5
	var verbose bool = false

	var err error
	var retVal bool = true
	var result string = ""

	ipaddress, ok := r.URL.Query()["ipaddress"]
	count, ok := r.URL.Query()["count"]
	timeout, ok := r.URL.Query()["timeout"]
	v, ok := r.URL.Query()["v"]

	if !ok || len(ipaddress[0]) < 1 {

	} else {
		ipAddress = ipaddress[0]
	}

	if !ok || len(count[0]) < 1 {

	} else {
		pingCount, err = strconv.Atoi(count[0])
		if err != nil {
			result = err.Error()
			retVal = false
		}
	}

	if !ok || len(timeout[0]) < 1 {

	} else {
		timeOutSecond, err = strconv.Atoi(timeout[0])
		if err != nil {
			result = err.Error()
			retVal = false
		}
	}

	if !ok || len(v[0]) < 1 {

	} else {
		verbose = v[0] == "1"
	}

	if retVal {
		retVal, result = pingOp(ipAddress, pingCount, timeOutSecond, verbose)
	}
	w.Header().Set("Content-Type", "application/json")
	resultVal := ResultVal{
		Host:    ipAddress,
		Result:  retVal,
		Message: result,
	}

	json.NewEncoder(w).Encode(resultVal)
}

func pingOp(ipAddress string, pingCount int, timeOutSecond int, verbose bool) (bool, string) {

	var retVal bool = false
	var result string = ""
	pinger, err := devhsmtekping.NewPinger(ipAddress)
	pinger.SetPrivileged(true)
	pinger.Count = pingCount
	pinger.Timeout = time.Duration(time.Second * time.Duration(timeOutSecond))
	if err != nil {
		result = err.Error()
		retVal = false
	}

	pinger.OnRecv = func(pkt *devhsmtekping.Packet) {
		if verbose {
			result += fmt.Sprintf("%d bytes from %s: icmp_seq=%d time=%v\n",
				pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
		}
	}

	pinger.OnFinish = func(stats *devhsmtekping.Statistics) {
		if verbose {
			result += fmt.Sprintf("\n--- %s ping statistics ---\n", stats.Addr)
			result += fmt.Sprintf("%d packets transmitted, %d packets received, %v%% packet loss\n",
				stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			result += fmt.Sprintf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
				stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		}
		if stats.PacketLoss == 100 {
			retVal = false
		} else {
			retVal = true
		}
	}

	err = pinger.Run()
	if err != nil {
		result = err.Error()
		retVal = false
	}
	if !verbose {
		if retVal {
			result = "PING SUCCESS"
		} else {
			result = "PING FAIL"
		}
	}
	return retVal, result
}
