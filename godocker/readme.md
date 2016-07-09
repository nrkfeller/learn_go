# GO Docker
$ sudo service docker start
$ sudo docker run hello-world
$ sudo docker build -t docker-whale .
$ sudo docker images
sudo docker version

### Basic Commands
##### Run
* -d tells docker container to run in background
* -P map ports to host ports -> so we can see apps run locally (5000 is default flask port)

```
docker run -d -P training/webapp python app.py
```
##### Stop
```
docker stop nostalgic_morse
```
##### Start
```
docker start nostalgic_morse
```
##### Restart - turn off than on again
```
docker restart
```
##### Check mapped ports with port commands
```
docker port pedantic_goldwasser
```
##### Check the logs for you webapp
```
docker logs -f pedantic_goldwasser
```
##### Check the processes running
```
docker stop pedantic_goldwasser
```
##### Configs
Check the configuration statuses - returns a JSON of configs
```
sudo docker inspect pedantic_goldwasser
```
for more specific needs
```
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' nostalgic_morse
```
##### Remove all docker images
```
docker rmi $(docker images -q)
```
##### Stop and remove all docker containers
```
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```
### Running containers
* -l flag is for the latest.
* -a see stopped containers
* PORTS are how the docker engine maps the ports

```
docker ps -l
```
### Tag - Push
```
sudo docker images
```
Get the image ID for the image you want.

Tag your image to the repo. Kind of like remote!
```
docker tag f8e64a2c5ee3 nicolasfeller/docker-whale-nick:latest
```
Login to Docker
```
docker login --username=nicolasfeller --email=me@nicolasfeller.com
```
Push!
```
sudo docker push nicolasfeller/docker-whale-nick
```
### Pull
Remove docker images
```
sudo docker rmi -f 7d9495d03763
sudo docker rmi -f docker-whale
```
Pull - Try new pulled image
```
sudo docker run nicolasfeller/docker-whale-nick
```
Check images to see if it worked!

### Commit
Go on the docker command line
```
docker run -t -i training/sinatra /bin/bash
```
Make some changes
```
root@0b2616b0e5a8:/# gem install json
```
Commit these changes
```
docker commit -m "Added json gem" -a "Kate Smith" 0b2616b0e5a8 ouruser/sinatra:v2
```

* * *
### Build your own images
Some repos have multiple variants of and image. for example ubuntu.
To get the specific image:
```
docker run -t -i ubuntu:14.04 /bin/bash
```
If you want to preload an image without running it:
```
docker pull centos
# then run it later
docker run -t -i centos /bin/bash
```
Look up some docker images from command line - repo search
```
docker search flask
```

### Build Image from dockerfile
When you use ```docker run``` you are either taking an image from the hub or using one that is already saved on your host and spinning up a container.
* Manage and wor with locally store images
* Create basic images
* Upload images from the hub

##### Access image
This will get the lastest version of ubuntu and start a shell!
```
docker run -t -i ubuntu /bin/bash
```
