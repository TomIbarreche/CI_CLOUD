# Projet T-CLO-902

Toutes les commandes s’exécutent depuis la racine du projet

## Installation avec le rôle sysadmin
```bash
kubectl config use-context sysadmin
```

### Création des namespaces
```bash
kubectl create ns devops
kubectl create ns monitoring
```
### Création des roles
```bash
kubectl apply -f .\rbac\role-manage.yaml
kubectl apply -f .\rbac\role-monitoring.yaml
kubectl apply -f .\rbac\rolebinding-manage.yaml
kubectl apply -f .\rbac\rolebinging-monitoring.yaml
```

### Installation de prometheus:

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack --set prometheus-node-exporter.hostRootFsMount.enabled=false -n=monitoring
```

### L'application grafana est accessbile de la manière suivante:
```bash
kubectl get pod
#Copier le nom du pod prometheus-graphana
kubectl port-forward <nom du pod prometheus-graphana> 3000 -n=monitoring
#Se rendre sur localhost:3000
```

### Installation de Loki
```bash
helm repo add loki https://grafana.github.io/loki/charts
helm repo update
helm upgrade --install loki loki/loki-stack -n=monitoring
```
### Installation du service ingress:

```bash
kubectl apply -f .\helm\ready\ingress\ingress.yaml
kubectl apply -f .\helm\ready\ingress\ingress-service.yaml -n=devops
```

## Installation avec le rôle devops
```bash
kubectl config use-context devops
```

### Installation de la chart elasticsearch:

```bash
helm install elasticsearch .\helm\ready\elastic\ -n=devops
```

### Installation de la chart mysql:

```bash
helm install mysql .\helm\ready\mysql\ -n=devops
```

### Installation de la chart rabbitmq

```bash
helm install rabbitmq .\helm\ready\rabbitmq\ -n=devops
```

### Installation des chart front, back, indexer et reporting

```bash
helm repo add angular-application https://tomibarreche.github.io/angular-application/
helm install front angular-application/angular-chart -n=devops

helm repo add laravel-application https://tomibarreche.github.io/laravel-application/
helm install back laravel-application/laravel-chart -n=devops

helm repo add indexer-application https://tomibarreche.github.io/indexer-application/
helm install indexer indexer-application/indexer-chart -n=devops

helm repo add reporting-application https://tomibarreche.github.io/reporting-application/
helm install reporting reporting-application/reporting-chart -n=devops
```

### L’application est accessible à l’adresse [localhost:80](http://localhost:80) (si k8 tourne avec un autre moteur que minikube)
### Si minikube
```bash
minikube tunnel
kubectl get svc -n=ingress-nginx # accessible depuis l'EXTERNAL-IP
```

### Désinstallation des charts:

```bash
kubectl config use-context sysadmin
helm uninstall elasticsearch, mysql, rabbitmq, front, back, indexer, reporting -n=devops
helm uninstall loki, prometheus -n=monitoring
```