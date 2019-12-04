package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"strconv"
	"time"
	"math/rand"
)

var powerConsumption = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "power_Consumption", Help: "Displays the power consumption in Watts of an specific VM"})

func getPower(){
	cmd := exec.Command("/bin/sh","-c","sudo ipmi-sensors -h localhost --no-sensor-type-output --no-header-output --comma-separated-output --sensor-types Current")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// Output from the previous ipmi sensor: 52,PSU_Input_Power,204.00,W,'OK'
	// We are only interested in the watts, so we only store that part. 
	
	watts,err := strconv.ParseFloat(strings.Split(string(out), ",")[2],64)
	powerConsumption.Set(watts)
}

func init() {
	prometheus.MustRegister(powerConsumption)
	getPower()
}

func self_update(){
	rand.Seed(time.Now().Unix())
	for {
		getPower()
		time.Sleep(time.Second)
	}
}
