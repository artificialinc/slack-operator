package pkgutil

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/artificialinc/slack-operator/pkg/config"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"

	slackv1alpha1 "github.com/artificialinc/slack-operator/api/v1alpha1"
	reconcilerUtil "github.com/stakater/operator-utils/util/reconciler"
)

// MapErrorListToError maps multiple errors into a single error
func MapErrorListToError(errs []error) error {

	if len(errs) == 0 {
		return nil
	}

	errMsg := []string{}
	for _, err := range errs {
		errMsg = append(errMsg, err.Error())
	}

	return fmt.Errorf(strings.Join(errMsg, "\n"))
}

func ManageError(ctx context.Context, client k8sClient.Client, channelInstance *slackv1alpha1.Channel, issue error) (ctrl.Result, error) {

	// Base object for patch, which patches using the merge-patch strategy with the given object as base.
	channelInstancePatchBase := k8sClient.MergeFrom(channelInstance.DeepCopy())

	// Update status
	channelInstance.Status.Conditions = []metav1.Condition{
		{
			Type:               "ReconcileError",
			LastTransitionTime: metav1.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Now().Location()),
			Message:            issue.Error(),
			Reason:             reconcilerUtil.FailedReason,
			Status:             metav1.ConditionTrue,
		},
	}

	// Patch status
	err := client.Status().Patch(ctx, channelInstance, channelInstancePatchBase)
	if err != nil {
		return ctrl.Result{}, err
	}

	return reconcilerUtil.RequeueAfter(config.ErrorRequeueTime)
}
