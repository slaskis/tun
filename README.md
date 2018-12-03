# tun

A very minimal TLS tunnel.

My own use case is to connect to AWS ElastiCache Redis with Encryption In-Transit enabled.

While the AWS docs recommend setting up a stunnel with this it's simply:

```sh
go get -u github.com/slaskis/tun
tun --connect master.xxxxxxxxxx.xxxx.euc1.cache.amazonaws.com:6379 &
redis-cli -p 6060
```

or if you, like me, proxy the connection through an instance within the VPC simply

```sh
# run proxy on $PORT
tun --connect $PORT --servername master.xxxxxxxxxx.xxxx.euc1.cache.amazonaws.com
redis-cli -p 6060
```

## Usage

```
Usage of tun:
  -connect string
    	connect address
  -insecure
    	skip tls verify
  -listen string
    	listening address (default "127.0.0.1:6060")
  -servername string
    	remote tls servername
```
