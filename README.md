## Overview

Go client for [ETH Gas Station](https://docs.ethgasstation.info/). 

### Older Go versions
```sh
go get github.com/osmian/gasstation-go
```

### Setup
The client requires an API Key from [Defi Pulse Data](https://data.defipulse.com/).


## Examples

### Create Client

```go
apiKey := "<API_KEY>"
client := gasstation.NewGasStationClient(apiKey)
```

### Get Gas Price

```go
gasPrice, err := client.GetGasPrice()
if err != nil {
    return err
}
```

### Get PredictionTable

```go
predictions, err := client.GetPredictionTable()
if err != nil {
    return err
}
```