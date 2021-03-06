# Frabric 8 Locally

http://fabric8.io/guide/getStarted/vagrant.html

# Spin up Ubuntu Instance

## 1.0 Spin-up Amazon EC2 Instance
* Create amazon AWS account
* Navigate to console home page
* Navigate AWS Services>Compute - click on EC2
* Click on Launch Instance, under Create Instance
* Select Ubuntu (64-bit), Choose tier (free tier works)
* Create a key pair - call it whatever you want
* Download the .pem file and place it in your local ~/.ssh/ directory
* Click Review Instance Lauch
* Click "Edit security Groups"
* Click "Add Rule"
* Select Type-> Custom TCP Rule || Port Range -> 5901 || Source -> Anywhere
* Click Launch Instance
* Navigate to your local ~/.ssh/ directory
* Navigate to "Instances/Instances" in your AWS console
* Right click in instance and choose "Connect"
* Copy and Run "ssh -i ..." command. (if you are denied access try to "chmod 600 yourkeyfilename.pem")

*You should be connected to your ec2 instance*

# Setup VNCserver to get GUI
* get script "ec2vncsetup.sh"
* send it to the ec2 instance:
```
scp -i ~/.ssh/yourpemfile.pem ec2vncsetup.sh ubuntu@ec2-54-213-234-237.us-west-2.compute.amazonaws.com:~/ec2vncsetup.sh
```
* run it ```./ec2vncsetup.sh```
* You need to accept prompts and enter a password at some point!
* In AWS console, click on the instance and get the public IP address (something like 54.213.234.237)
* Download VNCviewer on you local machine
* When running VNCviewer use 54.213.234.237::5901 - public IP::5901
* this should connect to the instance via gui

# Install Docker - AWS Ubuntu
*This can be done just with the ec2dockersetup.sh script**
* Install Docker
```
sudo apt-get install docker.io
```
* Add docker user group with root access, if it doesn't yet exist. Then add your current user to that group. (you can find your user name by typing ```whoami``` in the command line)
```
sudo groupadd docker
sudo usermod -aG docker $YOURUSERNAME
```
* Log out of shell and log back in. type ```exit``` and ssh back into the instance
* Test your docker installation
```
docker run hello-world
```
Should give you a message that everything is well setup

# Pre-Setup for Kubernetes - AWS Ubuntu
### Install awscli
*Amazon web services command line interface*

https://aws.amazon.com/cli/
1. Install pip ```sudo apt-get install python-pip```
2. Install awscli ```pip install awscli``` if this does not work ```sudo -H pip install awscli --upgrade --ignore-installed six```

test install by typing ```aws help``` and for autocomplete run ```complete -C aws_completer aws```

### Get and setup AWS access keys and credentials
1. Create iam user groups. follow http://docs.aws.amazon.com/IAM/latest/UserGuide/getting-started_create-admin-group.html
2. In your aws IAM console download the csv file with AWSAccessKeyId and AWSSecretKey
3. configure aws
```
$ aws configure
```
```
AWS Access Key ID: $YOURACCESKEY
AWS Secret Access Key: $YOURSECRETKEY
Default region name [us-west-2]: us-west-2 (or whatever your region is)
Default output format [None]: json (or text or wtv you want)
```
4. check if you have successfully configures your instance by typing ```cat ~/.aws/config```

### Setup Kubectl
add k8s variable to bash profile 
```
sudo vi ~/.profile

export KUBERNETES_PROVIDER=aws # use the AWS specific scripts
export KUBE_AWS_ZONE=eu-west-1c # choose your region, the default is us-west-2
export NUM_NODES=3 # choose the number of nodes you want  
export MASTER_SIZE=m3.large
export NODE_SIZE=m3.xlarge
export KUBE_ENABLE_INSECURE_REGISTRY=true # required to set the insecure registry flag on each node so we can push images to the cluster docker registry
```
Get Kubernetes with default AWS
```
export KUBERNETES_PROVIDER=aws; wget -q -O - https://get.k8s.io | bash
```
move the kubectl to user/local/bin
```
cd kubernetes
sudo mv kubectl /usr/local/bin
```
Run setup script
```
bash cluster/kube-up.sh
```

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

### Install Kops (if you want)
It is my opinion that this will be the best way of doing things. But the tool is not yet production ready. (you can skip this step)
https://github.com/kubernetes/kops

# Install Kubernetes and kubectl - AWS Ubuntu
1. Get kubernetes and do a quick setup script
```
export KUBERNETES_PROVIDER=aws; wget -q -O - https://get.k8s.io | bash
```
2. setup you path
```
export PATH=<path/to/kubernetes-directory>/platforms/linux/amd64:$PATH
(for example export PATH=/home/ubuntu/kubernetes/platforms/linux/amd64:$PATH )
```

3. Setup kubeconfig...
