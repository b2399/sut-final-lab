package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string // ต้องไม่เป็นค่าว่าง
	Email      string
	CustomerID string // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}

func TestCustomerPass(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check data", func(t *testing.T) {
		c := Customer{
			Name:       "Butsarakam",
			Email:      "butsarakam@gmail.com",
			CustomerID: "M1234567",
		}

		ok, err := govalidator.ValidateStruct(c)

		g.Expect(ok).To(BeTrue())

		g.Expect(err).To(BeNil())

	})
}
