package entity_test

import (
	"testing"

	"github.com/NPimtrll/sut-final-lab/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerValid(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Name is valid`, func(*testing.T) {
		customer := entity.Customers{
			Name:       "Test",
			Age:        "20",
			CustomerID: "CM12345678",
		}

		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
