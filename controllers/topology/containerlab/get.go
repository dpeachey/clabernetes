package containerlab

import (
	"context"

	clabernetesapistopologyv1alpha1 "github.com/srl-labs/clabernetes/apis/topology/v1alpha1"
	apimachinerytypes "k8s.io/apimachinery/pkg/types"
	ctrlruntime "sigs.k8s.io/controller-runtime"
)

// getContainerlabFromReq fetches the reconcile target Containerlab topology from the Request.
func (c *Controller) getContainerlabFromReq(
	ctx context.Context,
	req ctrlruntime.Request,
) (*clabernetesapistopologyv1alpha1.Containerlab, error) {
	containerlab := &clabernetesapistopologyv1alpha1.Containerlab{}

	err := c.BaseController.Client.Get(
		ctx,
		apimachinerytypes.NamespacedName{
			Namespace: req.Namespace,
			Name:      req.Name,
		},
		containerlab,
	)

	return containerlab, err
}
