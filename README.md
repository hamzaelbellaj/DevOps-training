# Kubernetes & Operator Assessment - Merkl

Ce dépôt contient les livrables pour le test technique d'entretien. L'objectif est de démontrer la maîtrise des concepts fondamentaux de Kubernetes via YAML et l'extension des capacités du cluster via le développement d'un Operator en Go.

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
- [x] **Task 2.1 — Scaffolding** : Initialisation du projet avec Operator SDK.
- [x] **Task 2.2 — CRD Definition** : Définition des champs `message` et `replicas` dans l'API.
- [ ] **Task 2.3 — Reconciliation** : Implémentation de la boucle de réconciliation (Logiciel de contrôle).
- [ ] **Task 2.4 — Validation** : Déploiement et tests fonctionnels.

---

## 🚀 Partie 1 : Implémentation & Reproduction

### 1. Application des manifestes
Les fichiers ont été appliqués dans l'ordre chronologique suivant :

```bash
kubectl apply -f k8s-manifests/task1-workload.yaml
kubectl apply -f k8s-manifests/task2-configmap.yaml
kubectl apply -f k8s-manifests/task3-resources.yaml
kubectl apply -f k8s-manifests/task4-probes.yaml
```

### 2. Validation du succès
La configuration a été vérifiée avec kubectl describe deployment nginx-deployment -n intern-assessment.

Ressources : Limits et Requests conformes (50m/64Mi - 100m/128Mi).

Probes : Liveness et Readiness actives (port 80).

Config : Variable APP_ENV injectée dynamiquement depuis le ConfigMap.

🏗 Partie 2 : Operator SDK (Détails techniques)

Task 2.1 & 2.2 — Initialisation et API
Le projet a été isolé dans le dossier /hello-operator pour respecter les contraintes du SDK et éviter les conflits avec les manifestes YAML.

Modifications effectuées dans api/v1alpha1/helloapp_types.go :

Spec : Ajout de Message (string) et Replicas (int32).

Status : Ajout de AvailableReplicas (int32) pour le monitoring.

Bash
# Génération des manifestes techniques (CRD)
cd hello-operator
make generate
make manifests
Note: Ce projet a été réalisé en utilisant l'assistance de l'IA (Gemini 3 Flash) pour la structuration méthodologique et l'explication des concepts théoriques.