# bbx

## Description
The `bbx` provides some commonly used tools, the mini version of busybox, so it is called bbx

## Functions
1. echo: Generate an echo serve
   1. tcp: A TCP server
   2. udp: A UDP server
   3. http: A HTTP server
2. mfa: A google authenticator cmd client
3. verify: Verify something
   1. cert: Verify cert valid
4. merge: Merge something
   1. kubeconfig: Merge many kubeconfig to a single file
5. godaddy: A godaddy cmd client
   1. add: Add a record
   2. del: Delete a record
   3. get: Get a record information
   4. upd: Update a record
   5. get-all: Get all record of a domain
6. alert: Alert something to somewhere
   1. feishu: Alert something to feishu group
7. get: Get system information
   1. net: get net info
      1. ipv4: get ipv4 address
      2. ipv6: get ipv6 address


## Build
```shell
make
```
  
## Install
```shell 
make install
```
  
## Generate releases
```shell
make gen-releases

# Generate with TAG
make gen-releases TAG=v1.0.1
```

## Build docker image
```shell
make build-image
```

## Usage
```shell
./bbx help

./bbx help echo
./bbx echo --help
```

## Other
- If you feel that the current bbx still has any problems or the functions are not rich enough or any other problems, please submit an Issue to me.


## Thinks
- cobra: Thanks to them for providing an excellent modern command line tool library.
- resty: Thanks to them for providing an out-of-the-box http client library.
- k8s: Thanks to them for providing such an easy-to-operate API library for kubeconfig.
- Thanks to all third-party library communities and authors who are indirectly dependent on.