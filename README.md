# go-config-server

go-config-server pulls configuration for remote clients from various sources

## Refes

- https://cloud.spring.io/spring-cloud-config/2.1.x/single/spring-cloud-config.html#_security
- https://github.com/pavel-v-chernykh/keystore-go
- https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/

## Run

```bash
go run cmd/main.go -config ./configs/application-file.yml
```

## Templates

```yml
# git - application.yml

server:
 port: 8001

security:
 apikey:
  enabled: false
  keylookup: "query:apikey"
  token: 9a0697eb595309177

encrypt:
 key: s3Cr3tk3y00

logging:
 level:
  root: INFO

spring:
 profiles:
  active: native, git, vault
 cloud:
  config:
   server:
    native:
     searchLocations: ./configs
    git:
     uri: https://github.com/helderfarias/go-config-server-examples
     clone_dir: ./tmp
     force_pull: true
    vault:
     uri: http://localhost:8200
 nats:
  servers: nats://localhost:4222
  subject: springCloudBus
  auth:
   type: token
   token: S3cretT0ken
```
