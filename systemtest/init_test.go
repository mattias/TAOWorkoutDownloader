package system_test

import (
	"github.com/mattias/TAOWorkoutDownloader/tao"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

// TODO: Rewrite to use real command instead of the client here. A real system test...
var client tao.Client

var _ = BeforeSuite(func() {
	client.Init()
})

func TestSystemtest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Suite")
}
