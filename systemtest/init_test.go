package system_test

import (
	"github.com/mattias/TAOWorkoutDownloader/tao"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var client tao.Client

var _ = BeforeSuite(func() {
	client.Init()
})

func TestSystemtest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Suite")
}
