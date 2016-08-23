#!/bin/bash
# Docker setup

echo "Setting up Docker for aws ec2!"

# sudo apt-get upgrade
sudo apt-get update

sudo apt-get install docker.io

sudo groupadd docker
sudo usermod -aG docker ubuntu

docker run hello-world
ubuntu@ip-172-31-24-54:~$ cat ec2vncsetup.sh
#!/bin/bash
# vncserver setup

echo "Setting up vnc for aws ec2!"

# sudo apt-get upgrade
sudo apt-get update

# get desktop setup installations
sudo apt-get install ubuntu-desktop
sudo apt-get install vnc4server
sudo apt-get install gnome-panel

# test vncserver
vncserver
vncserver -kill :1

# uncomment lines
sed -i '4,4 s/^##*//' .vnc/xstartup

# add gnome stuff
sed -i '6i\gnome-session –session=gnome-classic &\' .vnc/xstartup
sed -i '7i\gnome-panel&' .vnc/xstartup

# start server
vncserver
