#/bin/bash
kubectl apply -f k8s/configmap.yaml && \
kubectl apply -f k8s/mongodb-secret.yaml && \
kubectl apply -f k8s/deployment.yaml && \
kubectl apply -f k8s/service.yaml