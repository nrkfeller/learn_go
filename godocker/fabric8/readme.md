# Fabric8 on AWS

## 1.0 Spin-up Amazon EC2 Instance
* Create amazon account
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

## 2.0 Setup AWSCLI
*
