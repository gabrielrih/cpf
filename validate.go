package cpf

import (
	"strconv"
	"strings"
)

const CPF_LENGTH = 11
const FACTOR_FIRST_VERIFIER_DIGIT = 10
const FACTOR_SECOND_VERIFIER_DIGIT = 11

func Validate(rawcpf string) bool {
	if len(rawcpf) == 0 { //empty
		return false
	}
	cpf := cleanCPF(rawcpf)
	if len(cpf) != CPF_LENGTH {
		return false
	}
	if areAllTheDigitsTheSame(cpf) {
		return false
	}
	firstVerifierDigit := calculateDigit(cpf, FACTOR_FIRST_VERIFIER_DIGIT)
	secondVerifierDigit := calculateDigit(cpf, FACTOR_SECOND_VERIFIER_DIGIT)
	calculatedVerifierDigit := strconv.Itoa(firstVerifierDigit) + strconv.Itoa(secondVerifierDigit)
	verifierDigit := extractVerifierDigits(cpf)
	return calculatedVerifierDigit == verifierDigit
}

func cleanCPF(cpf string) string {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	return cpf
}

func areAllTheDigitsTheSame(cpf string) bool {
	var lastDigit string
	isFirst := true
	for index := 0; index < len(cpf); index++ {
		currentDigit := cpf[index : index+1]
		if !isFirst {
			if currentDigit != lastDigit {
				return false
			}
		}
		isFirst = false
		lastDigit = currentDigit //alimenta para a próxima iteração
	}
	return true
}

func calculateDigit(cpf string, factor int) int {
	resultado := 0
	for index := 0; index < len(cpf); index++ {
		if factor > 1 {
			currentDigit, _ := strconv.Atoi(cpf[index : index+1])
			resultado += currentDigit * factor
			factor -= 1
		}
	}
	rest := resultado % 11
	if rest < 2 {
		return 0
	} else {
		return 11 - rest
	}
}

func extractVerifierDigits(cpf string) string {
	return cpf[9:11]
}
