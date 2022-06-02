# Projet T-CLO-902

Toutes les commandes s’exécutent depuis la racine du projet
### A propos

Cette chart helm à pour fonction d'installer l'application angular et ses dépendances:
Elle deploie:
* Un pod contenant le projet laravel
* Un service établissant un cluster ip sur le port 4200

### Installation
```
helm repo add angular-application https://tomibarreche.github.io/angular-chart/
helm install <chart-name> angular-application/angular-chart -n=devops
```

### Test
L'api peut être contacté via:
```
localhost:80
```

### Désinstallation
```
helm uninstall <chart-name> -n=devops
```