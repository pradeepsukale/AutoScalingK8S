# AutoScalingK8S

Dockerfile: Contains multi stage build which produces small image. First part just generates the image and in second stage only the image is copied.


Start ASG for deployment: kubectl autoscale deployment basic-hpaimage --cpu-percent=80  --max=4
