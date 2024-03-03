package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ethereumPoolBalanceMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "selaginella_ethereum_pool_balance",
			Help: "The balance of ethereum pool"},
		[]string{"PoolAddress", "tokenAddress"},
	)
	opPoolBalanceMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "selaginella_optimism_pool_balance",
			Help: "The balance of optimism pool"},
		[]string{"PoolAddress", "tokenAddress"},
	)
	polygonZkEvmPoolBalanceMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "selaginella_polygon_zkEvm_pool_balance",
			Help: "The balance of polygonZkEvm pool"},
		[]string{"PoolAddress", "tokenAddress"},
	)
	scrollPoolBalanceMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "selaginella_scroll_pool_balance",
			Help: "The balance of scroll pool"},
		[]string{"PoolAddress", "tokenAddress"},
	)
	basePoolBalanceMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "selaginella_base_pool_balance",
			Help: "The balance of base pool"},
		[]string{"PoolAddress", "tokenAddress"},
	)
	arbPoolBalanceMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "selaginella_arb_pool_balance",
			Help: "The balance of arb pool"},
		[]string{"PoolAddress", "tokenAddress"},
	)
	chainAverageBalanceMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "selaginella_chain_average_balance",
			Help: "The average balance of chain pool"},
		[]string{"symbol"},
	)
)

func init() {
	prometheus.MustRegister(ethereumPoolBalanceMetric)
	prometheus.MustRegister(opPoolBalanceMetric)
	prometheus.MustRegister(polygonZkEvmPoolBalanceMetric)
	prometheus.MustRegister(scrollPoolBalanceMetric)
	prometheus.MustRegister(basePoolBalanceMetric)
	prometheus.MustRegister(arbPoolBalanceMetric)
}
