# Oolio-Orderahead


### Pre-requisites

If you want to avail coupon applying feature on order placement, we need to populate coupons to redis cache.
<br />
Upload your coupon files in root `coupons` folder.
You can download the files from here
[file1](https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase1.gz)
[file2](https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase2.gz)
[file3](https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase3.gz)

Run `data/data_populator.go` to add coupons to redis (default its assumed it's running on `localhost:6379`)
This processing and populating cache will take some time.

Create a `.env` file, that contains your redis host and port details, and the API-key.

Example:
```
API_KEY=api-secret
REDIS_HOST=127.0.0.1
REDIS_PORT=6379

```

<hr/>

### Commands

To run application : `go run main.go`

To run handler level test cases : `go test ./handlers`