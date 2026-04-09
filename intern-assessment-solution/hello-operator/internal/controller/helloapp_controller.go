package controller

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1alpha1 "github.com/hamzaelbellaj/hello-operator/api/v1alpha1"
)

// HelloAppReconciler reconciles a HelloApp object
type HelloAppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Permissions nécessaires (RBAC)
// +kubebuilder:rbac:groups=apps.intern.dev,resources=helloapps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.intern.dev,resources=helloapps/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

func (r *HelloAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	// 1. Récupérer l'objet HelloApp
	helloApp := &appsv1alpha1.HelloApp{}
	err := r.Get(ctx, req.NamespacedName, helloApp)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// 2. Définir le Deployment correspondant
	deploymentName := helloApp.Name + "-deployment"

	// On définit ce qu'on VEUT voir sur le cluster
	desiredDep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName,
			Namespace: helloApp.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &helloApp.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": helloApp.Name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": helloApp.Name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:    "busybox",
						Image:   "busybox:1.36",
						Command: []string{"sh", "-c", fmt.Sprintf("echo %s && sleep 3600", helloApp.Spec.Message)},
					}},
				},
			},
		},
	}

	// 3. Lier le Deployment à la HelloApp (Garbage Collection)
	if err := ctrl.SetControllerReference(helloApp, desiredDep, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// 4. Vérifier si le Deployment existe déjà
	existingDep := &appsv1.Deployment{}
	err = r.Get(ctx, client.ObjectKey{Name: deploymentName, Namespace: helloApp.Namespace}, existingDep)

	if err != nil && errors.IsNotFound(err) {
		// Le Deployment n'existe pas -> On le crée
		l.Info("Creating Deployment", "Name", deploymentName)
		if err := r.Create(ctx, desiredDep); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}

	// 5. Mettre à jour le Status avec les réplicas disponibles
	if helloApp.Status.AvailableReplicas != existingDep.Status.AvailableReplicas {
		helloApp.Status.AvailableReplicas = existingDep.Status.AvailableReplicas
		l.Info("Updating status", "AvailableReplicas", existingDep.Status.AvailableReplicas)
		if err := r.Status().Update(ctx, helloApp); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *HelloAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.HelloApp{}).
		Owns(&appsv1.Deployment{}). // Surveille les modifs sur les Deployments enfants
		Complete(r)
}
