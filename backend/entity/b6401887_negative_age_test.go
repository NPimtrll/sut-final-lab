package entity_test

import (
	"testing"

	"github.com/NPimtrll/sut-final-lab/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerAge(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`age is not integer`, func(*testing.T) {
		costomer := entity.Customers{
			Name:       "Test",
			Age:        "20",
			CustomerID: "CM12345678",
		}

		ok, err := govalidator.ValidateStruct(costomer)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Age is not integer"))
	})
}
