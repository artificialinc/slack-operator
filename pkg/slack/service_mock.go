package slack

import (
	"github.com/artificialinc/slack-operator/pkg/slack/mock"
	"github.com/go-logr/logr"
	"github.com/slack-go/slack"
)

var mockSlackService *SlackService

// NewMockService creates a mock service with SlackTestServer
func NewMockService(log logr.Logger) *SlackService {

	if mockSlackService == nil {
		testServer := mock.InitSlackTestServer()
		go testServer.Start()

		log.Info("Starting Test Server", "url", testServer.GetAPIURL())

		opts := slack.OptionAPIURL(testServer.GetAPIURL())

		mockSlackService = &SlackService{
			api: slack.New("apitoken", opts),
			log: log.WithName("SlackService"),
		}
	}

	return mockSlackService
}
