/*
Copyright 2018 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"golang.org/x/oauth2/google"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	pksv1alpha1 "github.com/crossplaneio/crossplane/pkg/apis/pks/v1alpha1"
)

const (
	controllerName = "pks.provider"
	// errorRetrievingSecret   = "Failed retrieving provider secret"
	// errorInvalidCredentials = "Invalid provider credentials"
)

// Add creates a new Provider Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// Reconciler reconciles a Provider object
type Reconciler struct {
	client.Client
	scheme     *runtime.Scheme
	kubeclient kubernetes.Interface
	recorder   record.EventRecorder

	validate func(*google.Credentials, []string) error
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	r := &Reconciler{
		Client:     mgr.GetClient(),
		scheme:     mgr.GetScheme(),
		kubeclient: kubernetes.NewForConfigOrDie(mgr.GetConfig()),
		recorder:   mgr.GetRecorder(controllerName),
	}
	r.validate = r._validate
	return r
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New(controllerName, mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Provider
	err = c.Watch(&source.Kind{Type: &pksv1alpha1.Provider{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

func (r *Reconciler) _validate(creds *google.Credentials, permissions []string) error {
	// if len(permissions) == 0 {
	// 	return nil
	// }
	//
	// return pks.TestPermissions(creds, permissions)

	return nil
}

// Reconcile reads that state of the cluster for a Provider object and makes changes based on the state read
// and what is in the Provider.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pks.crossplane.io,resources=provider,verbs=get;list;watch;create;update;patch;delete
func (r *Reconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// log.V(logging.Debug).Info("reconciling", "kind", pksv1alpha1.ProviderKindAPIVersion, "request", request)
	// // Fetch the Provider instance
	// instance := &pksv1alpha1.Provider{}
	// err := r.Get(ctx, request.NamespacedName, instance)
	// if err != nil {
	// 	if errors.IsNotFound(err) {
	// 		// Object not found, return.  Created objects are automatically garbage collected.
	// 		// For additional cleanup logic use finalizers.
	// 		return result, nil
	// 	}
	// 	// Error reading the object - requeue the request.
	// 	return result, err
	// }
	//
	// creds, err := pks.ProviderCredentials(r.kubeclient, instance)
	// if err != nil {
	// 	return r.fail(instance, errorRetrievingSecret, err.Error())
	// }
	//
	// err = r.validate(creds, instance.Spec.RequiredPermissions)
	// if err != nil {
	// 	return r.fail(instance, errorInvalidCredentials, err.Error())
	// }
	//
	// if instance.Status.IsReady() {
	// 	return result, nil
	// }
	//
	// // Update status condition
	// instance.Status.UnsetAllConditions()
	// instance.Status.SetReady()
	//
	// return result, r.Update(ctx, instance)
	return reconcile.Result{}, nil
}
