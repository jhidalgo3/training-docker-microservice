# Go microservice to play with ConfigMap and Secret in Kubernetes

By default this microservice load config file `env.yaml` from directories:

* "./configs"
* "$HOME/configs"

And it is using prefix `APP_` to overwriting values from enviroment variables.

Example config file:

```yaml
port: :8080

mongo:
  host: localhost
  port: 27017
  database: test

logger:
  logFile: /tmp/file.log
```

# Docker

## Build Docker image

```
docker build -t jhidalgo3/training-docker-microservice --rm .
```

## Run Docker image

```
docker run --rm  -ti -p 8080:8080 -e VERSION=1.0 -v $PWD/src/github.com/jhidalgo3/training-docker-microservice/configs:/root/configs jhidalgo3/training-docker-microservice
```

# URLS

API
+ http://localhost:8080/api/config
+ http://localhost:8080/api/info

VUE application
+ http://localhost:8080/static/

# Github repository

[https://github.com/jhidalgo3/training-docker-microservice](https://github.com/jhidalgo3/training-docker-microservice)

Author: Jose Maria Hidalgo Garcia
MIT License