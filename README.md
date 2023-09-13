# YAddress Web API Go Client
Makes calls into YAddress Web API for postal address correction, validation,
standardization and geocoding.

Find more about YAddress Web API at http://www.yaddress.net/WebApi

## Usage

### Geocode

```go
func main() {
    yd, err := yaddress.NewClient("")
	
	if err != nil {
		panic(err)
	}
	result, err := yd.ProcessAddress("42370 Bob Hope Dr, Rancho Mirage, CA")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Yaddress Result %v", result)
}
```

## Tests

You can run tests:
```
go test -v -cover
```