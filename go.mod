module github.com/artificialinc/slack-operator

go 1.16

require (
	github.com/go-logr/logr v1.2.4
	github.com/hashicorp/go-retryablehttp v0.7.2
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.27.10
	github.com/slack-go/slack v0.7.2
	github.com/stakater/operator-utils v0.1.13
	github.com/stretchr/testify v1.8.2
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.28.1
	k8s.io/client-go v0.28.1
	sigs.k8s.io/controller-runtime v0.16.1
)
