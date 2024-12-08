# Prometheus
## Metric type 
- Counter 
- Gauge 
- Histogram 
- Summary
## PromQL (Prometheus Query Language)
- https://prometheus.io/docs/prometheus/latest/querying/basics/
- https://promlabs.com/promql-cheat-sheet/

# Grafana
## Grafana Dashboard
### Add data source from Prometheus
1. Click menu "Connections" > "Add new connection"
2. Select "Prometheus" data source
3. Click button "Add new data source"
4. Set Name 
5. Set Prometheus server URL
6. Click "Save & test"

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
minikube start --cpus 2 --memory 4000 --vm-driver docker
```