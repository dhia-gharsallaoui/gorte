# go-rte

An RTE API client enabling Go programs to interact with RTE API in a simple and uniform way.

[![GitHub license](https://img.shields.io/github/license/dhia-gharsallaoui/go-rte.svg?)](https://github.com/dhia-gharsallaoui/go-rte/blob/main/LICENSE)
[![GoDoc](https://godoc.org/github.com/dhia-gharsallaoui/go-rte?status.svg)](https://pkg.go.dev/github.com/dhia-gharsallaoui/go-rte?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/dhia-gharsallaoui/go-odoo)](https://goreportcard.com/report/github.com/dhia-gharsallaoui/go-rte)
[![GitHub issues](https://img.shields.io/github/issues/dhia-gharsallaoui/go-rte.svg)](https://github.com/dhia-gharsallaoui/go-rte/issues)

## Overview
RTE, the french electricity transmission system operator, provides
access to various data through an API on its [data
portal](https://data.rte-france.com/home). You can retrieve those data
with `go-rte`.


## Installation

```bash
$ go get -u github.com/dhia-gharsallaoui/go-rte
```
## Quick Start

Add this import line to the file you're working in:
```Go
import "github.com/dhia-gharsallaoui/go-rte
```


### Client Authentication

To access the API, you need to [create an account](https://data.rte-france.com/create_account), or login if you have one.
Once logged, you can subscribe (create an application) to the desired API (each APIs must be subscribed individualy), you’ll obtain a base64 encoded key.
Then to set up a client use the command 
```Go
key:= "YmFiMWY3NjMtODhjZC00LWE5ZTgtOTRmMDc1ODcyYmNjOmU5YTIxNDVjLTBkOGUZi04YWI2LWRlNjRmODExM2M"
client, err := gorte.NewClient(gorte.ClientConfig{Token: key})
```
> NB: the token generated with your key is only valid for two hours. So you need to recreated the client each two hours using the same key.

### Usage
The package came with a useful struct called `Period` for calling API's with `start date` and `end date` configuration. this struct ensure the coding of the time in the suitable format.

```Go
layout := "2006-01-02 15:04"
st, err := time.Parse(layout, "2022-03-01 23:00")
if err != nil {
  fmt.Println(err)
  }
et, err := time.Parse(layout, "2022-03-09 13:00")
if err != nil {
  fmt.Println(err)
  }
opt := Period{Time(st), Time(et)}

```
After preparing our configuration we can call an API from the available categories.

For exemple to get the "Peak Period" signals from `2022-03-01 23:00` to `2022-03-09 13:00`.
```Go
signals, _, err := Client.Market.GetSignals(opt)
```

### Available APIs

#### Consumption
- [ ] [Consolidated Consumption](https://data.rte-france.com/catalog/-/api/consumption/Consolidated-Consumption/v1.0)
- [x] [Consumption](https://data.rte-france.com/catalog/-/api/consumption/Consumption/v1.2)
```Go
client.Consumption.GetAnnualForecasts(opt)
client.Consumption.GetShortTerm(opt)
client.Consumption.GetWeeklyForecasts(opt)
```
- [x] [Demand Response](https://data.rte-france.com/catalog/-/api/consumption/Demand-Response/v1.0)
```Go
client.Consumption.GetOperators(opt)
client.Consumption.GetVolumes(opt)
```
- [x] [Electricity Quality](https://data.rte-france.com/catalog/-/api/consumption/Electricity-Quality/v1.0)
```Go
client.Consumption.GetQualityData(gorte.GetQualityDataOptions{
                ID:        0,
                StartDate: opt.StartDate,
                EndDate:   opt.EndDate,
                Type:      "RMS,RMS_VOLTAGE",
        })
```
- [x] [Ecowatt](https://data.rte-france.com/catalog/-/api/consumption/Ecowatt/v3.0)
```Go
client.Consumption.GetSignalEcowatt(opt)
```
- [x] [Tempo Like Supply Contract](https://data.rte-france.com/catalog/-/api/consumption/Tempo-Like-Supply-Contract/v1.1)
```Go 
client.Consumption.GetTempoLikeCalendars(opt)
```
#### Exchanges
> Not implemented yet.

#### Generation
- [ ] [Actual Generation](https://data.rte-france.com/catalog/-/api/generation/Actual-Generation/v1.1)
- [x] [Generation Forecast](https://data.rte-france.com/catalog/-/api/generation/Generation-Forecast/v2.0)
- [ ] [Generation Installed Capacities](https://data.rte-france.com/catalog/-/api/generation/Generation-Installed-Capacities/v1.0)
- [ ] [Unavailability Additional Information](https://data.rte-france.com/catalog/-/api/generation/Unavailability-Additional-Information/v3.0)

#### Market
- [x] [Balancing Capacity](https://data.rte-france.com/catalog/-/api/market/Balancing-Capacity/v4.1) 
```Go
client.Market.GetAcceptedOffers(opt)
client.Market.GetAggregatedoffersAFRREnergybids(opt)
client.Market.GetAggregatedoffersEnergybids(opt)
client.Market.GetDailyProcuredReserves(opt)
client.Market.GetImbalance(opt)
client.Market.GetIndividualoffersEnergybids(opt)
client.Market.GetInsufficientsOffers(opt)
client.Market.GetMarginsData(opt)
client.Market.GetNeeds(opt)
client.Market.GetPeakDailyMargins(opt)
client.Market.GetProcuredReservesResp(opt)
```
- [ ] [Balancing Energy](https://data.rte-france.com/catalog/-/api/market/Balancing-Energy/v1.2)
- [ ] [Bre Referential](https://data.rte-france.com/catalog/-/api/market/Bre-Referential/v1.0)
- [x] [Demand Response Signal](https://data.rte-france.com/catalog/-/api/market/Demand-Response-Signal/v1.0)
```Go                           
client.Market.GetOperators(opt)
client.Market.GetVolumes(opt) 
```
- [ ] [Wholesale Market](https://data.rte-france.com/catalog/-/api/market/Wholesale-Market/v2.0)

#### Partners
> Not implemented yet.

### Enjoy Coding
```Go
package main

import (
        "fmt"
        "time"

        gorte "github.com/dhia-gharsallaoui/go-rte"
)

func main() {
        key := "YmFiMWY3NjMtODhjZC00LWE5ZTgtOTRmMDc1ODcyYmNjOmU5YTIxNDVjLTBkOGUZi04YWI2LWRlNjRmODExM2M"
        client, err := gorte.NewClient(gorte.ClientConfig{Token: key})
        layout := "2006-01-02 15:04"
        st, err := time.Parse(layout, "2022-03-01 23:00")
        if err != nil {
                fmt.Println(err)
        }
        et, err := time.Parse(layout, "2022-03-09 13:00")
        if err != nil {
                fmt.Println(err)
        }
        opt := gorte.Period{gorte.Time(st), gorte.Time(et)}
        signals, _, err := client.Market.GetSignals(opt)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(signals)
}

```


## Todo

- Tests
- Add other RTE API

## Issues

- If you have an issue, please report it on the [issue tracker](https://github.com/dhia-gharsallaoui/go-rte/issues)

## Sponsored by 
  [Skilld](https://www.skilld.cloud/)

