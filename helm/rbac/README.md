# Projet T-CLO-902

Toutes les commandes s’exécutent depuis la racine du projet
### A propos

Cette chart helm à pour fonction d'installer les roles suivants:
* Manage pods
* Moniroting pods

### Installation
```
helm repo add rbac https://tomibarreche.github.io/laravel-chart/
helm install <chart-name> rbac/rbac-chart -n=devops
```

### Désinstallation
```
helm uninstall <chart-name> -n=devops
```