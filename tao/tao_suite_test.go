package tao_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTao(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tao Suite")
}
