# Prometheus
## PromQL (Prometheus Query Language) 

# Grafana
## Grafana Dashboard
### Searching dashboard ID
1. Go to https://grafana.com/grafana/dashboards/
2. Input keyword in search box (k6, redis, blackbox)
3. Click button "Copy ID to clipboard"

# Terraform
- Infrastructure as a code
## Install Terraform
1. Download binary from https://developer.hashicorp.com/terraform/install
2. Add Terraform binary to PATH environment variable
``` bash
# Open the .bash_profile file with a text editor
code ~/.bash_profile
```
3. Use the export command to add new environment variables
``` bash
# export [existing_variable_name]=[new_variable_value]:$[existing_variable_name]
export PATH="$PATH:~/Downloads/terraform"
```
4. Execute the new .bash_profile by either restarting the terminal window or using
``` bash
source ~/.bash-profile
```
5. Check a PATH environment variable
``` bash
echo $PATH
```
6. Verify the installation
``` bash
terraform -help
```

# Minikube
## Install Minikube
``` bash
brew install minikube
```

## Start Minikube
``` bash
minikube start
# or
minikube start --cpus 2 --memory 4000 --vm-driver hyperkit
```