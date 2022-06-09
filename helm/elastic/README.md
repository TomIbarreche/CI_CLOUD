# Projet T-CLO-902

Toutes les commandes s’exécutent depuis la racine du projet
### A propos

Cette chart helm à pour fonction d'installer le serveur elasticsearch
Elle deploie:
* Un pod contenant le serveur elasticsearch
* Un service de type cluster IP
* Une configmap

### Installation
```
helm repo add elasticsearch https://tomibarreche.github.io/elasticsearch-chart/
helm install <chart-name> elasticsearch/elasticsearch-chart -n=devops
```

### Désinstallation
```
helm uninstall <chart-name> -n=devops
```