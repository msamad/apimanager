# Overview

**Under Construction**

API Manager in go

Project layout is based on https://github.com/golang-standards/project-layout

To run
```shell script
go run ./cmd/apimanager/main.go
```

To test
```shell script
curl http://localhost:8080/v1/ping
curl http://localhost:8080/v2/ping
```

Will use `~/go/bin` dir for binaries add that to PATH
```shell script
echo 'export PATH=$PATH:~/go/bin' >> ~/.profile
source ~/.profile
```

To generate code from OpenAPI specs
``` shell script
# install oapi-codegen
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
go build -o ~/go/bin/ github.com/deepmap/oapi-codegen/cmd/oapi-codegen

# generate code
./scripts/generate_from_spec.sh
```

ElasticSearch as data storage
```shell script
docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.2
```

# Libraries

|                                          |                                  | 
|------------------------------------------|----------------------------------|
| https://github.com/deepmap/oapi-codegen  |  Generate code from OpenAPI spec | 
| https://github.com/labstack/echo         | Router                           |
| https://github.com/rakyll/statik         | Package swagger ui in binary     |

# TODO

- Create binary with swagger ui (Use statik)