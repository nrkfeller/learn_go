#!/bin/bash
# Docker setup

echo "Setting up Docker for aws ec2!"

sudo apt-get upgrade
sudo apt-get update

sudo apt-get install docker.io

sudo groupadd docker
sudo usermod -aG docker ubuntu

docker run hello-world
