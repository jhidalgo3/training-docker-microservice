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