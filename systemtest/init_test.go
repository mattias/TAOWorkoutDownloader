package system_test

import (
	"github.com/mattias/TAOWorkoutDownloader/tao"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/oauth2"

	"testing"
)

var oauth2configuration oauth2.Config
var configuration tao.Configuration

var _ = BeforeSuite(func() {
	err := configuration.Load()
	Expect(err).NotTo(HaveOccurred())
})

func TestSystemtest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Systemtest Suite")
}
