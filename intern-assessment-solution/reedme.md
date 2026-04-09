# Kubernetes & Operator Assessment - Merkl

Ce dépôt contient les livrables pour le test technique d'entretien.

## 🛠 Environnement Local
- **OS**: macOS
- **Cluster**: `kind` (Kubernetes in Docker)
- **Version Go**: 1.21+
- **Outils**: `kubectl`, `operator-sdk`, `docker`

---

## 🗺 Roadmap de l'exercice

### Partie 1 : Concepts Core Kubernetes (YAML)
- [x] **Task 1.1 — Namespace & Workload** : Création du namespace `intern-assessment`, d'un déploiement Nginx (2 réplicas) et d'un Service NodePort.
- [x] **Task 1.2 — ConfigMap & Injection** : Création d'un ConfigMap pour gérer l'environnement (`APP_ENV: staging`) et injection dans les pods.
- [ ] **Task 1.3 — Resource Management** : Configuration des Requests/Limits CPU et Mémoire pour la stabilité du cluster.
- [ ] **Task 1.4 — Health Checks** : Implémentation des Liveness et Readiness probes.

### Partie 2 : Développement d'un Opérateur (Go)
- [ ] **Task 2.1** : Scaffolding du projet avec Operator SDK.
- [ ] **Task 2.2** : Définition de la CRD `HelloApp`.
- [ ] **Task 2.3** : Implémentation de la boucle de réconciliation (Reconcile Loop).
- [ ] **Task 2.4** : Déploiement et validation du contrôleur.

---

## 🚀 Comment reproduire (Partie 1)

1. **Appliquer les manifestes dans l'ordre :**
   ```bash
   kubectl apply -f k8s-manifests/task1-workload.yaml
   kubectl apply -f k8s-manifests/task2-configmap.yaml