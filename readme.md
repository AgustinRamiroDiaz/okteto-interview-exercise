# Okteto interview exercise

Usage guide:
- Clone the repository
- Install kubectl
- Follow some of the [Okteto's instructions](https://okteto.com/docs/samples/golang/)
  - [Install Okteto CLI](https://okteto.com/docs/getting-started/installation/) 
  - [Configure access to your Okteto Cloud Namespace](https://okteto.com/docs/cloud/credentials/)
  - Apply the kubernetes manifest
  ```kubectl apply -f k8s.yml```
  - Activate your development container: ```okteto up```
  - Inside the container, run the application: ```go run main.go```
- Now you can access the application from your browser at ```http://localhost:8080``` or from Okteto cloud at ```https://hello-world-<YOUR_USER_ID>.cloud.okteto.net```. There are currently 2 endpoints, the root which returns a hello world message and the ```/pods``` endpoint which returns a pretty message of all the pods in the namespace.