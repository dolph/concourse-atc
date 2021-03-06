package dbng_test

import (
	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourceType", func() {
	var pipeline dbng.Pipeline

	BeforeEach(func() {
		var (
			created bool
			err     error
		)

		pipeline, created, err = defaultTeam.SavePipeline(
			"pipeline-with-types",
			atc.Config{
				ResourceTypes: atc.ResourceTypes{
					{
						Name:   "some-type",
						Type:   "docker-image",
						Source: atc.Source{"some": "repository"},
					},
					{
						Name:   "some-other-type",
						Type:   "docker-image-ng",
						Source: atc.Source{"some": "other-repository"},
					},
				},
			},
			0,
			dbng.PipelineUnpaused,
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(created).To(BeTrue())
	})

	Describe("(Pipeline).ResourceTypes", func() {
		var resourceTypes []dbng.ResourceType

		JustBeforeEach(func() {
			var err error
			resourceTypes, err = pipeline.ResourceTypes()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns the resource types", func() {
			Expect(resourceTypes).To(HaveLen(2))

			ids := map[int]struct{}{}

			for _, t := range resourceTypes {
				ids[t.ID()] = struct{}{}

				switch t.Name() {
				case "some-type":
					Expect(t.Name()).To(Equal("some-type"))
					Expect(t.Type()).To(Equal("docker-image"))
					Expect(t.Source()).To(Equal(atc.Source{"some": "repository"}))
					Expect(t.Version()).To(BeNil())
				case "some-other-type":
					Expect(t.Name()).To(Equal("some-other-type"))
					Expect(t.Type()).To(Equal("docker-image-ng"))
					Expect(t.Source()).To(Equal(atc.Source{"some": "other-repository"}))
					Expect(t.Version()).To(BeNil())
				}
			}

			Expect(ids).To(HaveLen(2))
		})

		Context("when a resource type becomes inactive", func() {
			BeforeEach(func() {
				var (
					created bool
					err     error
				)

				pipeline, created, err = defaultTeam.SavePipeline(
					"pipeline-with-types",
					atc.Config{
						ResourceTypes: atc.ResourceTypes{
							{
								Name:   "some-type",
								Type:   "docker-image",
								Source: atc.Source{"some": "repository"},
							},
						},
					},
					pipeline.ConfigVersion(),
					dbng.PipelineUnpaused,
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(created).To(BeFalse())
			})

			It("does not return inactive resource types", func() {
				Expect(resourceTypes).To(HaveLen(1))
				Expect(resourceTypes[0].Name()).To(Equal("some-type"))
			})
		})
	})
})
