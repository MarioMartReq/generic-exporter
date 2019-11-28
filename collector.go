package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var powerConsumption = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "power_Consumption", Help: "Displays the power consumption in Watts of an specific VM"})

func init() {
	prometheus.MustRegister(powerConsumption)

	powerConsumption.Set(0)
}