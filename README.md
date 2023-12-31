# YAddress Web API Go Client
Makes calls into YAddress Web API for postal address correction, validation,
standardization and geocoding.

Find more about YAddress Web API at http://www.yaddress.net/WebApi

## Usage

### Yaddress

```go
func main() {
    logger := yaddress.DefaultLogger()
    yd := yaddress.NewClient("123", yaddress.WithLogger(logger))
	
    
    request := yaddress.Request{AddressLine1: "506 Fourth Avenue Unit 1", AddressLine2: "Asbury Prk, NJ"}

    resp, err := yd.ProcessAddress(request)
}
```

You are able to provide:`logger`, `httpClient`

Each one is called with following syntax:

**With logger**
```go
    logger := yaddress.DefaultLogger()
    yd, err := yaddress.NewClient("123", yaddress.WithLogger(logger))
```

**With client**
```go
	client := &http.Client{Timeout: time.Second * 10}
	
	yd, err := yaddress.NewClient("", yaddress.WithClient(client))
```

**With both**
```go
	logger := yaddress.DefaultLogger()
	client := &http.Client{Timeout: time.Second * 10}
	
	yd, err := yaddress.NewClient("", yaddress.WithLogger(logger), yaddress.WithClient(client))
```

> **Warning**
> Gives you first 1,000 transactions for free. You can find later pricing plans [here](https://www.yaddress.net/Pricing)


## Tests
**Tests are written using table _driven technique_**

You can run tests:
```
go test -v -cover
```
