package main

import (
	"ILoveChairs/uy_to_usd/fetcher"
	"ILoveChairs/uy_to_usd/gui"
)

func main() {
	convRate, err := fetcher.FetchConversionRate()
	if err != nil {
		convRate = 42.0
	}
	gui.Get(convRate).Run()
}
