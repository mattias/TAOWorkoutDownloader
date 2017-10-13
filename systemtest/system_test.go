package system_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the application", func() {
	It("can read the configuration file", func() {
		Expect(client.Config.Oauth2.Endpoint.AuthURL).To(Equal("https://beta.trainasone.com/oauth/authorise"))
		Expect(client.Config.FileType).To(Equal("fit"))
		Expect(client.Config.TargetType).To(Equal("heart_rate"))
	})

	It("saves the upcoming workout from TrainAsONE in specified path", func() {
		var path = "../"

		file_name, err := client.SaveNextWorkoutTo(path)
		defer os.Remove(path + file_name)
		Expect(err).NotTo(HaveOccurred())

		Expect(file_name).To(Equal("workout.fit"))

		Expect(path).Should(BeAnExistingFile())

	})
})
