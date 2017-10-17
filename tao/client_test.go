package tao_test

import (
	"os"

	"github.com/mattias/TAOWorkoutDownloader/tao"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var client tao.Client

var _ = BeforeSuite(func() {
	client.Init()
})

var _ = Describe("Client", func() {
	It("can read the configuration file", func() {
		Expect(client.Config.Oauth2.ClientID).NotTo(BeEmpty())
		Expect(client.Config.Oauth2.Endpoint.AuthURL).To(Equal("https://beta.trainasone.com/oauth/authorise"))
		Expect(client.Config.Workout.FileType).To(Equal("fit"))
		Expect(client.Config.Workout.TargetType).To(Equal("heart_rate"))
		Expect(client.Config.DevicePath).To(Equal("../"))
	})

	It("can download a workout from the web service", func() {
		var path string = "../"
		file_name, err := client.SaveNextWorkoutTo(path)

		defer os.Remove(path + file_name)
		Expect(err).NotTo(HaveOccurred())

		Expect(file_name).To(Equal("workout.fit"))

		Expect(path).Should(BeAnExistingFile())
	})

})
