# Kubernetes & Operator Assessment - Merkl

Ce dépôt contient les livrables pour le test technique.
L'objectif est de démontrer la maîtrise des concepts fondamentaux de Kubernetes via YAML et l'extension des capacités du cluster via le développement d'un Operator en Go.

---

## 🛠 Environnement Local

* **OS** : macOS (MacBook Pro)
* **Cluster** : kind (Kubernetes in Docker)
* **Version Go** : 1.26+
* **Outils** : `kubectl`, `operator-sdk`, `docker`

---

## 🗺 Roadmap de l'exercice

### Partie 1 : Concepts Core Kubernetes (YAML)

* [x] **Task 1.1 — Namespace & Workload**
  Création du namespace `intern-assessment`, d'un déploiement Nginx (2 réplicas) et d'un Service NodePort.

* [x] **Task 1.2 — ConfigMap & Injection**
  Création d'un ConfigMap pour gérer l'environnement (`APP_ENV: staging`) et injection dans les pods.

* [x] **Task 1.3 — Resource Management**
  Configuration des Requests/Limits CPU et Mémoire (`50m/64Mi - 100m/128Mi`).

* [x] **Task 1.4 — Health Checks**
  Implémentation des Liveness et Readiness probes.

---

### Partie 2 : Développement d'un Opérateur (Go)

* [x] **Task 2.1 — Scaffolding**
  Initialisation du projet avec Operator SDK.

* [x] **Task 2.2 — CRD Definition**
  Définition des champs `message` et `replicas` dans l'API.

* [x] **Task 2.3 — Reconciliation**
  Implémentation de la logique de création du Deployment Busybox et synchronisation du status.

* [x] **Task 2.4 — Validation**
  Déploiement et tests fonctionnels réussis.

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

---

### 2. Validation du succès

La configuration a été vérifiée avec :

```bash
kubectl describe deployment nginx-deployment -n intern-assessment
```

**Résultats :**

* **Ressources** : Limits et Requests conformes (`50m/64Mi - 100m/128Mi`)
* **Probes** : Liveness et Readiness actives (port 80)
* **Config** : Variable `APP_ENV` injectée dynamiquement depuis le ConfigMap

---

## 🏗 Partie 2 : Operator SDK (Détails techniques)

### Task 2.1 & 2.2 — Initialisation et API

Le projet est isolé dans le dossier `/hello-operator`.

Les modifications dans `api/v1alpha1/helloapp_types.go` permettent de définir :

* l'état souhaité (**Spec**)
* l'état observé (**Status**)

---

### Task 2.3 — Logique du Contrôleur

Le contrôleur dans `internal/controller/helloapp_controller.go` assure les fonctions suivantes :

* **Watch** : Surveillance des ressources `HelloApp` et des Deployments enfants
* **Reconcile** : Création/Mise à jour d'un deployment Busybox injectant le message via une variable d'environnement (`MY_MESSAGE`)
* **OwnerReference** : Liaison des cycles de vie pour un nettoyage automatique (Garbage Collection)
* **Status Update** : Remontée du nombre de réplicas réels vers l'objet Custom Resource

---

### Task 2.4 — Validation Finale

Le bon fonctionnement a été validé en appliquant une ressource de test :

```bash
# Installation et lancement
make install
make run

# Vérification des logs du pod généré
kubectl logs -n intern-assessment helloapp-sample-deployment-6f8459d9d-kpt7w
```

**Output :**

```text
Bonjour Merkl, l'operateur fonctionne !
```

---

## 📝 Note

Ce projet a été réalisé en utilisant l'assistance de l'IA (**Gemini 3 Flash**) pour :

* la structuration méthodologique
* l'explication des concepts théoriques
