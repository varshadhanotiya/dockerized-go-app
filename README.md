# Sample dev project

Simple project that takes in a stock ticker and date ranges and returns highest and lowest prices. Includes a Go server, a Python frontend, and a PostgreSQL database. Containerized these components, deploy them using Docker Compose and Kubernetes (Minikube), and ensured they can communicate with each other.

## Docker setup
### Building Docker Images:
Built Docker images for Go server and Python frontend applications using docker build commands.
```
docker build -t go_server .
docker run -p 8080:8080 go_server
docker ps
```

### Docker Compose:
Used a Docker Compose file to define services for Go server, Python frontend, and PostgreSQL database. Defined a network for them to communicate.

### Docker Network:
Created a Docker network for services to communicate within.

## Kubernetes (Minikube) Setup
### Minikube Installation:
Installed Minikube to create a local Kubernetes cluster.

```
minikube start
minikube image load image_name      #If image is not remote and only availabel locally
minikube dashboard
```
### Deployment Manifests:
Created Kubernetes deployment manifest files (deployment.yaml) for Go server and Python frontend.

### Service Manifests:
Created Kubernetes service manifest files (service.yaml) for Go server and Python frontend.

### Applying Manifests:
Used kubectl apply to deploy Kubernetes resources.
```
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

### Communication Between Services:
After deploying to Kubernetes, frontend should send requests to the service names of backend components.
```kubectl get pods
kubectl describe pod pod_name
kubectl logs pod_name
kubectl delete deployment deployment_name
kubectl delete service service_name
```

## Errors and Troubleshooting
### ImagePullBackoff and ErrImagePull:
Faced issues with images not being pulled or running in pods.

Solution: Ensure the images are accessible to Kubernetes cluster. Use minikube image load to load local images.
### Pending Status of Pods:
Pods were in a "Pending" state.

Solution: Check Minikube status, allocate more resources if needed, and check for any issues with resources or configuration.
### Connection Errors:
You encountered connection errors between components.

Solution: Ensure services are using correct names for communication. Update connection strings in  server code to use service names.
### Further Steps
Debugging Logs:
Check pod logs using kubectl logs to diagnose issues within pods.

### Minikube Dashboard:
Use minikube dashboard to access the Kubernetes dashboard for monitoring and debugging.


## Commands

Build command for creating docker image
docker build -t  varsha/docker_training:server_v1 -f .\docker_files\Dockerfile.server .
docker build -t  varsha/docker_training:app_v1 -f .\docker_files\Dockerfile.app .

Command to push the image to the docker hub
docker push varsha/docker_training:server_v1
docker push varsha/docker_training:app_v1

command to deploy kubernetes resoucres
kubectl apply -f server.yaml
kubectl apply -f server_service.yaml
kubectl apply -f app.yaml
kubectl apply -f app_service.yaml

minikube start

Run below command and let it be running in a terminal
minikube dashboard



To expose a Kubernetes service running within a Minikube cluster and open a web browser to view that service
Leave this service running in a terminal
minikube service go-server
minikube service pytthon-frontend

command to see the logs of the pods
kubectl logs -f -l app=go-server --all-containers=true
kubectl logs -f -l app=python-frontend --all-containers=true
