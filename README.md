# Make my own Terraform with Golang

This project is a Go implementation that mimics Terraform to create AWS EC2 instances using AWS CLI and YAML configuration files.

## Prerequisites

1. **AWS CLI**: You need to have AWS CLI installed and configured. Follow the instructions here to install AWS CLI: [AWS CLI Installation Guide](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)

2. **Go**: Ensure you have Go installed on your machine. Download and install Go from here: [Go Installation Guide](https://golang.org/doc/install)

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

**State Management** : Implement state management to keep track of resources created and manage updates/destroy operations.

**Configuration Validation** : Add validation for the configuration files to ensure they meet the necessary criteria before applying or planning.

**User Interface** : Develop a simple UI for managing configurations and executing commands more easily.


