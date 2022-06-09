# Projet T-CLO-902

Toutes les commandes s’exécutent depuis la racine du projet
### A propos

Cette chart helm à pour fonction d'installer l'application de reporting:
Elle deploie:
* Un cron job qui s'exécute tout les soirs à minuit

### Installation
```
helm repo add reporting-application https://tomibarreche.github.io/reporting-chart/
helm install <chart-name> reporting-application/reporting-chart -n=devops
```

### Désinstallation
```
helm uninstall <chart-name> -n=devops
```