package cpf

import "testing"

func TestAllDigitsEqualsWithoutMask(test *testing.T) {
	cpf := "11111111111"
	isValid := Validate(cpf)
	if isValid {
		test.Fatalf("CPF %v deveria ser inválido!", cpf)
	}
}

func TestAllDigitsEqualsWithMask(test *testing.T) {
	cpf := "111.111.111-11"
	isValid := Validate(cpf)
	if isValid {
		test.Fatalf("CPF %v deveria ser inválido!", cpf)
	}
}

func TestValidCPFWithMask(test *testing.T) {
	cpf := "029.464.780-50"
	isValid := Validate(cpf)
	if !isValid {
		test.Fatalf("CPF com máscara %v deveria ser válido!", cpf)
	}
}

func TestValidCPFWithoutMask(test *testing.T) {
	cpf := "02849444081"
	isValid := Validate(cpf)
	if !isValid {
		test.Fatalf("CPF sem máscara %v deveria ser válido!", cpf)
	}
}
