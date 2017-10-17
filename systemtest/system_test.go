package system_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the application", func() {
	// TODO: Rewrite to use real command, this is a duplicate of the client unit test
	It("saves the upcoming workout from TrainAsONE in specified path", func() {
		var path string = "../"

		file_name, err := client.SaveNextWorkoutTo(path)
		defer os.Remove(path + file_name)
		Expect(err).NotTo(HaveOccurred())

		Expect(file_name).To(Equal("workout.fit"))

		Expect(path).Should(BeAnExistingFile())

	})
})
