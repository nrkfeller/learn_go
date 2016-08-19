#!/bin/bash
# Docker setup

echo "Setting up Docker for aws ec2!"

sudo apt-get upgrade
sudo apt-get update

sudo apt-get install docker.io

sudo groupadd docker
sudo usermod -aG docker ubuntu # replace ubuntu by whatever your username is

# logout and log back in
echo "NOW LOG OUT OF YOUR SHELL, LOG BACK IN AND TYPE - $ docker run hello-world"
