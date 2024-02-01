package entity_test

import (
	"testing"

	"github.com/NPimtrll/sut-final-lab/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`customer_id is not pattern`, func(*testing.T) {
		costomer := entity.Customers{
			Name:       "Test",
			Age:        "20",
			CustomerID: "CS12345678",
		}

		ok, err := govalidator.ValidateStruct(costomer)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("ComtomerID is not pattern"))
	})
}
