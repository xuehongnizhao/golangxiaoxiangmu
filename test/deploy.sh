#!/bin/sh

# deploy project script



ssh -tt 192.168.10.114 << remote

cd /home/yaml/test/canteen

kubectl delete rc canteen-v1 -n test

sleep 1

kubectl create -f canteen-rc.yaml -n test

exit

remote