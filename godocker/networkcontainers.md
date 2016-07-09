# Network Containers
You can create docker networks and put multiple containers in these networks. For example a webapp container and a database container. That would make sense!

1. Create/Pull an image, call it "web"
```
docker run -d -P --name web $(some docker image)
```
2. Create a docker (bridge) Network
```
docker network create -d bridge my-new-network
```
3. Create/Pull another image and add it to the network
```
docker run -d --net=my-new-network --name db training/postgres
```
4. Add the first image to the network
```
docker network connect my-new-network web
```
* * *
* at any time, inspect your network!
```
docker network inspect my-new-network
```
* Open a shell and see if your networks are well connected
```
docker exec -it db bash

ping web
or
ping 172.18.0.3 - web's ip address
```
