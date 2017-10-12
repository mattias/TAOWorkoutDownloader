package system_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the application", func() {
	It("can read the configuration file", func() {
		Expect(configuration.Oauth2.Endpoint.AuthURL).To(Equal("https://beta.trainasone.com/oauth/authorise"))
		Expect(configuration.FileType).To(Equal("fit"))
		Expect(configuration.TargetType).To(Equal("heart_rate"))
	})

	PIt("saves the upcoming workout from TrainAsONE in specified path", func() {

	})
})
