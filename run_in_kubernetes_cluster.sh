
# cleanup any previous mess
minikube delete

# start minikube
minikube start

# create docker image within scope of minikube
eval $(minikube docker-env)
docker image rm -f chessnet-image
docker build . -t chessnet-image

# deploy kubernetes bits
kubectl create -f kubernetes/namespace.yaml
kubectl create -f kubernetes/deployment.yaml
kubectl create -f kubernetes/service.yaml

# get endpoint to hit
minikube service chessnet-service --url -n chessnet-namespace
