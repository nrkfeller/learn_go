# Spin up Ubuntu Instance

## 1.0 Spin-up Amazon EC2 Instance
* Create amazon AWS account
* Navigate to console home page
* Navigate AWS Services>Compute - click on EC2
* Click on Launch Instance, under Create Instance
* Select Ubuntu (64-bit), Choose tier (free tier works)
* Create a key pair - call it whatever you want
* Download the .pem file and place it in your local ~/.ssh/ directory
* Click Launch Instance
* Navigate to your local ~/.ssh/ directory
* Navigate to "Instances/Instances" in your AWS console
* Right click in instance and choose "Connect"
* Copy and Run "ssh -i ..." command. (if you are denied access try to "chmod 600 yourkeyfilename.pem")

*You should be connected to your ec2 instance*

# Install Docker
Install Docker
```
sudo apt-get install docker.io
```
Add docker user group with root access, if it doesn't yet exist. Then add your current user to that group. (you can find your user name by typing ```whoami``` in the command line)
```
sudo groupadd docker
sudo usermod -aG docker $YOURUSERNAME
```
Test your docker installation
```
docker run hello-world
```
Should give you a message that everything is well setup

### Install go
1. Install the latest version of golang :D
```
$ sudo apt-add-repository ppa:ubuntu-lxc/lxd-stable
$ sudo apt-get update
$ sudo apt-get install golang
```
2. setup environment variables
```
$ sudo vi ~/.profile
```
Add these lines to the bottom of your bash profile file
```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/work
export GOBIN=$HOME/work/bin
```
For these changes to take place...
```
source ~/.profile
```
3. Test your go installation. it is convention to put all your go related stuff in work/src/github.com/.
```
$ mkdir -p work/src/github.com/user/hello
$ cd work/src/github.com/user/hello
$ touch main.go
$ vi main.go
```
Write this in main.go
```
package main
import "fmt"
func main() {
    fmt.Printf("hello, world\n")
}
```
and try your go install!
```
$ go run main.go
```

# Get Google Cloud SDK
http://kubernetes.io/docs/hellonode/
* install python2.7 or more
* go to https://console.cloud.google.com and create a google cloud account
* create a compute engine project
* enable billing
1. navigate to https://cloud.google.com/sdk/ and get the newest version (in this case 122.0.0)
```
wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-122.0.0-linux-x86.tar.gz
tar xvzf google-cloud-sdk-122.0.0-linux-x86.tar.gz
```
2. Install gcloud platform!
```
./google-cloud-sdk/install.sh
```
3. Initialize a project ```gcloud init```. This will prompt some selections: choose project, choose region, and use default for the rest.
4. install Kubernetes! command line tool
```
gcloud components install kubectl
```
5. Test installation with an nginx!
```
kubectl run nginx --image=nginx:1.10.0
```
