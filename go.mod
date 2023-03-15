module github.com/artificialinc/slack-operator

go 1.16

require (
	github.com/go-logr/logr v0.3.0
	github.com/hashicorp/go-retryablehttp v0.7.2
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/slack-go/slack v0.7.2
	github.com/stakater/operator-utils v0.1.13
	github.com/stretchr/testify v1.6.1
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.3
)
