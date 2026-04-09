# Kubernetes & Operator Assessment - Merkl

Ce dépôt contient les livrables pour le test technique. L'objectif est de démontrer la maîtrise des concepts fondamentaux de Kubernetes via YAML et l'extension des capacités du cluster via le développement d'un Operator en Go.

## 🛠 Environnement Local
- **OS**: macOS (MacBook Pro)
- **Cluster**: `kind` (Kubernetes in Docker)
- **Version Go**: 1.26+
- **Outils**: `kubectl`, `operator-sdk`, `docker`

---

## 🗺 Roadmap de l'exercice

### Partie 1 : Concepts Core Kubernetes (YAML)
- [x] **Task 1.1 — Namespace & Workload** : Création du namespace `intern-assessment`, d'un déploiement Nginx (2 réplicas) et d'un Service NodePort.
- [x] **Task 1.2 — ConfigMap & Injection** : Création d'un ConfigMap pour gérer l'environnement (`APP_ENV: staging`) et injection dans les pods.
- [x] **Task 1.3 — Resource Management** : Configuration des Requests/Limits CPU et Mémoire (50m/64Mi - 100m/128Mi).
- [x] **Task 1.4 — Health Checks** : Implémentation des Liveness et Readiness probes.

### Partie 2 : Développement d'un Opérateur (Go)
- [ ] **Task 2.1** : Scaffolding du projet avec Operator SDK.
- [ ] **Task 2.2** : Définition de la CRD `HelloApp`.
- [ ] **Task 2.3** : Implémentation de la boucle de réconciliation (Reconcile Loop).
- [ ] **Task 2.4** : Déploiement et validation du contrôleur.

---

## 🚀 Partie 1 : Implémentation & Reproduction

### 1. Application des manifestes
Les fichiers ont été créés et appliqués dans l'ordre chronologique suivant :

```bash
kubectl apply -f k8s-manifests/task1-workload.yaml
kubectl apply -f k8s-manifests/task2-configmap.yaml
kubectl apply -f k8s-manifests/task3-resources.yaml
kubectl apply -f k8s-manifests/task4-probes.yaml