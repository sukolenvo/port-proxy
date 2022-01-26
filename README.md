# Port proxy util

Temporary expose port for remote connections. E.g. database/windows docker containers.

### Install

* Unix
  ```bash
  wget https://github.com/sukolenvo/port-proxy/releases/download/v1.0.0/port-proxy
  chmod +x ./port-proxy
  ```
* Windows
  ```powershell
   curl -Uri "https://github.com/sukolenvo/port-proxy/releases/download/v1.0.0/port-proxy.exe" -OutFile "port-proxy.exe"
  ```
### Usage
* Help
  ```bash
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
* Forward port `7048` to `nav-test`
  ```powershell
  .\port-proxy.exe -port 7048 -connectAddress nav-test
  ```
* Forward port 80 to 8080
  ```bash
  ./port-proxy -port 80 -connectPort 8000
  ```