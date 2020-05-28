package controller

import (
	"context"

	"fmt"

	_ "github.com/adibrastegarnia/sample-k8s-controller/pkg/apis/v1"
	v1 "github.com/adibrastegarnia/sample-k8s-controller/pkg/apis/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var log = logf.Log.WithName("raft_storage_controller")

type Reconciler struct {
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Cluster object and makes changes based on the state read
// and what is in the Cluster.Spec
func (r *Reconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {

	imageProcessing := &v1.ImageProcessingClass{}

	err := r.client.Get(context.TODO(), request.NamespacedName, imageProcessing)
	if err != nil {
		log.Error(err, "Reconcile Database")
		if k8serrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{Requeue: true}, err
	}
	log.Info("Reconcile is called")
	fmt.Println("image", imageProcessing.Spec.Image)

	return reconcile.Result{}, nil

}
