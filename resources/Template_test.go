package resources

import (
	"time"

	"github.com/beevik/etree"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

const (
	vmTemplateXML = "xml/template.xml"
)

var _ = ginkgo.Describe("Template", func() {
	var (
		doc        *etree.Document
		vmTemplate *Template
		err        error
	)

	ginkgo.Describe("getters", func() {
		ginkgo.BeforeEach(func() {
			// create user with data
			doc = etree.NewDocument()
			err = doc.ReadFromFile(vmTemplateXML)
			vmTemplate = CreateTemplateFromXML(doc.Root())
		})

		ginkgo.It("should find all Template attributes", func() {
			gomega.Expect(err).NotTo(gomega.HaveOccurred()) // no error during BeforeEach
			gomega.Expect(vmTemplate).ShouldNot(gomega.BeNil())

			gomega.Expect(vmTemplate.ID()).To(gomega.Equal(5296))
			gomega.Expect(vmTemplate.Name()).To(gomega.Equal("Ubuntu"))

			var userID int
			userID, err = vmTemplate.User()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(userID).To(gomega.Equal(50))

			var groupID int
			groupID, err = vmTemplate.Group()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(groupID).To(gomega.Equal(101))

			var permissions *Permissions
			permissions, err = vmTemplate.Permissions()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(permissions.User.Use).To(gomega.Equal(true))
			gomega.Expect(permissions.User.Manage).To(gomega.Equal(true))
			gomega.Expect(permissions.User.Admin).To(gomega.Equal(false))
			gomega.Expect(permissions.Group.Use).To(gomega.Equal(true))
			gomega.Expect(permissions.Group.Manage).To(gomega.Equal(false))
			gomega.Expect(permissions.Group.Admin).To(gomega.Equal(false))
			gomega.Expect(permissions.Other.Use).To(gomega.Equal(false))
			gomega.Expect(permissions.Other.Manage).To(gomega.Equal(false))
			gomega.Expect(permissions.Other.Admin).To(gomega.Equal(false))

			regTime := time.Unix(int64(1526973359), 0)
			gomega.Expect(vmTemplate.RegistrationTime()).To(gomega.Equal(&regTime))
		})
	})

	ginkgo.Describe("create Template", func() {
		ginkgo.It("should create Template", func() {
			gomega.Expect(err).NotTo(gomega.HaveOccurred()) // no error during BeforeEach

			vmTemplate = CreateTemplateWithID(42)

			gomega.Expect(vmTemplate.ID()).To(gomega.Equal(42))
		})

		ginkgo.Context("when Template doesn't have given attribute", func() {
			ginkgo.BeforeEach(func() {
				vmTemplate = CreateTemplateWithID(42)
				err = nil
			})

			ginkgo.It("should return that Template doesn't have name", func() {
				_, err = vmTemplate.Name()
				gomega.Expect(err).To(gomega.HaveOccurred())
			})

			ginkgo.It("should return that Template doesn't have user", func() {
				_, err = vmTemplate.User()
				gomega.Expect(err).To(gomega.HaveOccurred())
			})

			ginkgo.It("should return that Template doesn't have group", func() {
				_, err = vmTemplate.Group()
				gomega.Expect(err).To(gomega.HaveOccurred())
			})

			ginkgo.It("should return that Template doesn't have permissions", func() {
				_, err = vmTemplate.Permissions()
				gomega.Expect(err).To(gomega.HaveOccurred())
			})

			ginkgo.It("should return that Template doesn't have registration time", func() {
				_, err = vmTemplate.RegistrationTime()
				gomega.Expect(err).To(gomega.HaveOccurred())
			})
		})
	})
})
