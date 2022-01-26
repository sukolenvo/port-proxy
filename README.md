# Port proxy util

Temporary expose port for remote connections. E.g. database/windows docker containers.

### Usage
```
./port-proxy -h
Usage of ./port-proxy:
  -address string
        Interface to listen on (default "0.0.0.0")
  -connectAddress string
        Address to connect to (default "127.0.0.1")
  -connectPort string
        Port to connect to (default same as -port)
  -port string
        Interface to listen on (default "8080")

```
