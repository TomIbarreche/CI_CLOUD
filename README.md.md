# Projet T-CLO-902

Toutes les commandes s’exécutent depuis la racine du projet
### Création des namespaces
```bash
kubectl create ns devops
kubectl create ns monitoring
```

### Installation du service ingress:

```bash
kubectl apply -f .\helm\ready\ingress\ingress.yaml 
kubectl apply -f .\helm\ready\ingress\ingress-service.yaml -n=devops
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
helm install front .\helm\ready\front-chart -n=devops
helm install back .\helm\ready\back-chart -n=devops
helm install indexer .\helm\ready\indexer-chart -n=devops
helm install reporting .\helm\ready\reporting-chart -n=devops
``` 

### L’application est accessible à l’adresse [localhost:80](http://localhost:80) (si k8 tourne avec un autre moteur que minikube)

### Installation de prometheus:

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack --set prometheus-node-exporter.hostRootFsMount.enabled=false -n=monitoring
```

### L'application graphna est accessbile de la manière suivante:
```bash
kubectl get pod
#Copier le nom du pod prometheus-graphana
kubectl port-forward <nom du pod prometheus-graphana> 3000
#Se rendre sur localhost:3000
```

### Installation de Loki
```bash
helm repo add loki https://grafana.github.io/loki/charts
helm repo update
helm upgrade --install loki loki/loki-stack -n=monitoring
```