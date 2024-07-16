# Make my own Terraform with Golang

This project is a Go implementation that mimics Terraform to create AWS EC2 instances using AWS CLI and YAML configuration files.

# Why I Created This Project
I created this project while studying for the 'AWS DOP-C02' certification to better understand AWS and Infrastructure as Code (IaC).
This project allowed me to create my own Terraform-like tool to easily handle AWS instances and configurations using YAML files.

## Prerequisites

1. **AWS CLI**: You need to have AWS CLI installed and configured. Follow the instructions here to install AWS CLI: [AWS CLI Installation Guide](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)

2. **Go**: Ensure you have Go installed on your machine. Download and install Go from here: [Go Installation Guide](https://golang.org/doc/install)

## Version 2 Features
1. **Simplified Build Process**
You no longer need to manually specify the operating system and architecture when building the project. Just run sh build.sh and the script will detect your OS and create the appropriate executable file.

2. **Enhanced State Management**
A tfstate file is now created for each resource based on its name. This allows you to track and manage the state of your resources more effectively.

3. **Multiple YAML File Handling**
The program now supports handling multiple YAML configuration files. You can specify multiple config files in a single command, and the program will process each one sequentially.

4. **Simplified AWS CLI Integration**
While you can still use shell scripts to run AWS CLI commands, this program allows you to define your infrastructure in YAML files and execute them directly, making the process much simpler and more intuitive.

## Setup

### AWS CLI Installation

1. **macOS**:
   ```sh
   brew install awscli

2. **Windows**:
Download and run the MSI installer from the AWS CLI official site.

3. **Linux**:

```sh
sudo apt-get update
sudo apt-get install awscli
```

After installing, configure AWS CLI with your credentials:

```
aws configure
```

This will prompt you to enter your AWS Access Key ID, Secret Access Key, default region name, and default output format. Example:

```mathematica
AWS Access Key ID [None]: YOUR_ACCESS_KEY
AWS Secret Access Key [None]: YOUR_SECRET_KEY
Default region name [None]: eu-central-1
Default output format [None]: json
```

### Go Installation
1. **macOS**:

```
brew install go
```


2. **Windows**:
Download and run the MSI installer from the Go official site.

3. **Linux**:
```
sudo apt-get update
sudo apt-get install golang
```

## Project Setup
Clone the repository:

```
git clone https://github.com/sejoonkimmm/My-Terraform-with-Golang.git
cd My-Terraform
```

### Build the project
1. **Apple Silicon (M1/M2)**

```
GOARCH=arm64 GOOS=darwin go build -o myterraform-darwin-arm64 main.go
```

2. **Intel macOS**
```
GOARCH=amd64 GOOS=darwin go build -o myterraform-darwin-amd64 main.go
```

3. **Windows**
```
GOARCH=amd64 GOOS=windows go build -o myterraform-windows-amd64.exe main.go
```

4. **Linux**
```
GOARCH=amd64 GOOS=linux go build -o myterraform-linux-amd64 main.go
```

### Usage
Create a configuration file config.yaml with the following content:

``` yaml
provider:
  region: eu-central-1

resources:
  - type: aws_instance
    name: example
    ami: ami-0346fd83e3383dcb4
    instance_type: t2.micro
Run the following command to apply the configuration and create the AWS EC2 instance:

```

1. **macOS** (Apple Silicon)
``` sh
./myterraform-darwin-arm64 apply config.yaml
```

2. **macOS** (Intel)
``` sh
./myterraform-darwin-amd64 apply config.yaml
```

3. **Windows**
``` sh
.\myterraform-windows-amd64.exe apply config.yaml
```

4. **Linux**
```
./myterraform-linux-amd64 apply config.yaml
```

To plan the infrastructure without actually creating it, run:

macOS (Apple Silicon)
```sh
./myterraform-darwin-arm64 plan config.yaml
```

macOS (Intel)
```sh
./myterraform-darwin-amd64 plan config.yaml
```

Windows
```sh
./myterraform-windows-amd64.exe plan config.yaml
```
Linux
```
./myterraform-linux-amd64 plan config.yaml
```


## Notes
AWS CLI Installation: This project assumes that AWS CLI is installed and configured. If you do not wish to install AWS CLI, you can modify the code to use the aws-sdk-go directly.

Platform-specific Binaries: The provided instructions and commands are specific to different platforms (Apple Silicon, Intel macOS, Windows, and Linux). Make sure to use the correct binary for your platform.


## Future Enhancements
**Support for more AWS resources** : Extend the tool to manage other AWS resources such as S3 buckets, RDS instances, and VPCs.

**Configuration Validation** : Add validation for the configuration files to ensure they meet the necessary criteria before applying or planning.

**User Interface** : Develop a simple UI for managing configurations and executing commands more easily.

**Show the AWS AMI list** : added 'amilist' command, show the AMI catalog list.


