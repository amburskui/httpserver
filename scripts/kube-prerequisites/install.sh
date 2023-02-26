#!/bin/bash
set -e

# settings global enviroments
alias kubectl="minikube kubectl --"

# kubernetes environments
NAMESPACE=m
HOMEWORK_HOST=arch.homework

# installing helm
# See https://helm.sh/docs/intro/install/.
helm=$(which helm) || echo "'helm' not found"
[ -z $helm ] && exit 1

# installing nginx ingress controller
kubectl create namespace ${NAMESPACE}
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/
helm repo update
helm install nginx ingress-nginx/ingress-nginx --namespace ${NAMESPACE} -f nginx_ingress-25239-20146a.yaml 

# change current context to ${NAMESPACE}
kubectl config set-context --current --namespace=${NAMESPACE}

# add resolve
cat<< EOF >>hosts
${HOMEWORK_HOST} `minikube ip`
EOF

# append hosts to /etc/hosts
# cat hosts >> /etc/hosts