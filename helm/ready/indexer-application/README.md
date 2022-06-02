# Projet T-CLO-902

Toutes les commandes s’exécutent depuis la racine du projet
### A propos

Cette chart helm à pour fonction d'installer l'application d'indexation:
Elle deploie:
* Un pod contenant l'application d'indexation
* Un service établissanty un cluster ip sur le port 3000

### Installation
```
helm repo add indexer-application https://tomibarreche.github.io/indexer-chart/
helm install <chart-name> indexer-application/indexer-chart -n=devops
```

### Désinstallation
```
helm uninstall <chart-name> -n=devops
```