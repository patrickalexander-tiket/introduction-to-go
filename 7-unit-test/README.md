# 7-unit-test
## Test Directory
### Under the same folder separated with `_test.go`
https://github.com/prometheus/prometheus/tree/main/scrape

### Special `test` folder to store all of unit test file
https://github.com/kubernetes/kubernetes/tree/master/test

## run test 
### All packages
```> $ go test -v ./...```

### With Race Detector 
```> $ go test -v -race ./...```

### Specific test case
```> $ go test -v resource/car/car_test.go resource/car/car.go -run=TestGetCarMultipleCase```

### Get File Coverage
``` > $ go test -cover ./... ```

### Generate HTML Coverage
``` 
> $ go test -coverprofile cover.out ./...
> $ go tool cover -html=cover.out -o cover.html
```
