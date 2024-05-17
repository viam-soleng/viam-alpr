# Docker Build Setup

```
docker build -t viam-alpr .

docker container run -it viam-alpr bash

docker create --name extract viam-alpr

docker cp extract:/openalpr/src/bindings/go/viam-alpr/module/module.tar.gz .

````

