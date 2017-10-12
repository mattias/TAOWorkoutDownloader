package system_test

import (
	"github.com/mattias/TAOWorkoutDownloader/tao"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var configuration tao.Configuration

var _ = BeforeSuite(func() {
	configuration.Load()
})

func TestSystemtest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Suite")
}
