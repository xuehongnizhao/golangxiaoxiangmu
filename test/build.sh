#!/bin/sh

# build project script


go build -i .


docker build -t 192.168.10.14:5000/canteen:v1.1.7 .

#docker tag platform-registry:v1 192.168.10.14:5000/platform-registry:v1 

docker push 192.168.10.14:5000/canteen:v1.1.7 