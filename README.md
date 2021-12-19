# AutoScalingK8S

Dockerfile: Contains multi stage build which produces small image. First part just generates the image and in second stage only the image is copied.

1. Need to enable metrics-server addon on minikube:
   a. minikube list addon
   b. minikube addons enable metrics-server
2. Start the deployment using yaml file
3. Start ASG for deployment 'basic-hpaimage': kubectl autoscale deployment basic-hpaimage --cpu-percent=80  --max=4. 
