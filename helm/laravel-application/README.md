# Projet T-CLO-902

Toutes les commandes s’exécutent depuis la racine du projet
### A propos

Cette chart helm à pour fonction d'installer l'application laravel et ses dépendances:
Elle deploie:
* Deux pods contenant le projet laravel
* Un service établissant un cluster ip sur le port 8000
* Une configmap pour mysql et rabbitmq
* Un secret pour mysql et rabbitmq

### Installation
```
helm repo add laravel-application https://tomibarreche.github.io/laravel-chart/
helm install <chart-name> laravel-application/laravel-chart -n=devops
```

### Test
L'api peut être contacté via:
```
localhost:80/api
```

### Désinstallation
```
helm uninstall <chart-name> -n=devops
```