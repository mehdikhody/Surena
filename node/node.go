package main

import (
	"fmt"
	"surena/node/xray"
)

func main() {
	stats, err := xray.Init().API.StatsService.QueryStats(false)
	if err != nil {
		return
	}

	for _, stat := range stats.GetStat() {
		fmt.Println(stat)
	}
}
