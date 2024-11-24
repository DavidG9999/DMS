package tests

import (
	"math/rand/v2"
	"strings"
)

func randomInnKpp() string {
	numbers := []rune("0123456789")
	symbol := []rune("/")
	var b strings.Builder
	for i := 0; i < 9; i++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	for j := 9; j < 10; j++ {
		b.WriteRune(symbol[rand.IntN(len(symbol))])
	}
	for j := 10; j < 20; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	return b.String()
}

func invalidInnKppMin() string {
	numbers := []rune("0123456789")
	symbol := []rune("/")
	var b strings.Builder
	for i := 0; i < 9; i++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	for j := 9; j < 10; j++ {
		b.WriteRune(symbol[rand.IntN(len(symbol))])
	}
	for j := 10; j < 19; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	return b.String()
}

func invalidInnKppMax() string {
	numbers := []rune("0123456789")
	symbol := []rune("/")
	var b strings.Builder
	for i := 0; i < 9; i++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	for j := 9; j < 10; j++ {
		b.WriteRune(symbol[rand.IntN(len(symbol))])
	}
	for j := 10; j < 21; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	return b.String()
}

func randomBank() struct {
	bankName     string
	bankIdNubmer string
} {

	randBank := []struct {
		bankName     string
		bankIdNubmer string
	}{
		{
			bankName:     "СберБанк",
			bankIdNubmer: "045004641",
		},
		{
			bankName:     "Банк ВТБ",
			bankIdNubmer: "044525411",
		},
		{
			bankName:     "Альфа-Банк",
			bankIdNubmer: "044525593",
		},
		{
			bankName:     "ГазпромБанк",
			bankIdNubmer: "044525823",
		},
		{
			bankName:     "МКБ",
			bankIdNubmer: "044525659",
		},
		{
			bankName:     "Совкомбанк",
			bankIdNubmer: "044525360",
		},
		{
			bankName:     "Т-Банк",
			bankIdNubmer: "044525974",
		},
		{
			bankName:     "Россельхозбанк",
			bankIdNubmer: "044525111",
		},
		{
			bankName:     "Банк ДОМ.РФ",
			bankIdNubmer: "044525266",
		},
		{
			bankName:     "Росбанк",
			bankIdNubmer: "044525256",
		},
	}

	return randBank[rand.IntN(len(randBank))]
}

func randomBankAccountNumber() string {
	chars := []rune("0123456789")
	length := 20
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.IntN(len(chars))])
	}
	return b.String()
}

func updateRandomBankAccountNumber() *string {
	chars := []rune("0123456789")
	length := 20
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.IntN(len(chars))])
	}
	result := b.String()
	return &result
}

func invalidBankAccountNumberMin() string {
	chars := []rune("0123456789")
	length := 19
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.IntN(len(chars))])
	}
	return b.String()
}

func invalidBankAccountNumberMax() string {
	chars := []rune("0123456789")
	length := 21
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.IntN(len(chars))])
	}
	return b.String()
}

func licenseClass() string {
	return "A, A1, B, B1, M, C, C1, D"
}

func randomLicense() string {
	numbers := []rune("0123456789")
	var b strings.Builder
	for i := 0; i < 10; i++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	return b.String()
}

func invalidLicenseMin() string {
	numbers := []rune("0123456789")
	var b strings.Builder
	for i := 0; i < 9; i++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	return b.String()
}

func invalidLicenseMax() string {
	numbers := []rune("0123456789")
	var b strings.Builder
	for i := 0; i < 11; i++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	return b.String()
}

func randomTruck() struct {
	brand string
	model string
} {
	randTruck := []struct {
		brand string
		model string
	}{
		{
			brand: "MAN",
			model: "TGM",
		},
		{
			brand: "MAN",
			model: "TGL",
		},
		{
			brand: "MAN",
			model: "TGS",
		},
		{
			brand: "MAN",
			model: "TGS WW",
		},
		{
			brand: "MAN",
			model: "TGX",
		},
		{
			brand: "MAN",
			model: "TGA",
		},
		{
			brand: "Volvo",
			model: "FH",
		},
		{
			brand: "Volvo",
			model: "FM",
		},
		{
			brand: "Volvo",
			model: "FMX",
		},
		{
			brand: "Volvo",
			model: "FH16",
		},
		{
			brand: "KAMAZ",
			model: "5490 NEO",
		},
		{
			brand: "KAMAZ",
			model: "54901 NEO",
		},
		{
			brand: "KAMAZ",
			model: "65209",
		},
		{
			brand: "KAMAZ",
			model: "65659",
		},
	}
	return randTruck[rand.IntN(len(randTruck))]
}

func randomStateNumber() string {
	numbers := []rune("0123456789")
	alphabet := []rune("АВЕКМНОРСТУХ")
	var b strings.Builder
	for i := 0; i < 1; i++ {
		b.WriteRune(alphabet[rand.IntN(len(alphabet))])
	}
	for j := 1; j < 4; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	for j := 4; j < 6; j++ {
		b.WriteRune(alphabet[rand.IntN(len(alphabet))])
	}
	for j := 6; j < 9; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	return b.String()
}

func updateRandomStateNumber() *string {
	numbers := []rune("0123456789")
	alphabet := []rune("АВЕКМНОРСТУХ")
	var b strings.Builder
	for i := 0; i < 1; i++ {
		b.WriteRune(alphabet[rand.IntN(len(alphabet))])
	}
	for j := 1; j < 4; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	for j := 4; j < 6; j++ {
		b.WriteRune(alphabet[rand.IntN(len(alphabet))])
	}
	for j := 6; j < 9; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	result := b.String()
	return &result
}

func invalidStateNumberMin() string {
	numbers := []rune("0123456789")
	alphabet := []rune("ЙЦУКЕНГШЩЗХФЫВАПРОЛДЖЯЧСМИТБ")
	var b strings.Builder
	for i := 0; i < 2; i++ {
		b.WriteRune(alphabet[rand.IntN(len(alphabet))])
	}
	for j := 2; j < 4; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	for j := 4; j < 5; j++ {
		b.WriteRune(alphabet[rand.IntN(len(alphabet))])
	}
	for j := 5; j < 7; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	return b.String()
}

func invalidStateNumberMax() string {
	numbers := []rune("0123456789")
	alphabet := []rune("ЙЦУКЕНГШЩЗХФЫВАПРОЛДЖЯЧСМИТБ")
	var b strings.Builder
	for i := 0; i < 2; i++ {
		b.WriteRune(alphabet[rand.IntN(len(alphabet))])
	}
	for j := 2; j < 5; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	for j := 5; j < 6; j++ {
		b.WriteRune(alphabet[rand.IntN(len(alphabet))])
	}
	for j := 6; j < 10; j++ {
		b.WriteRune(numbers[rand.IntN(len(numbers))])
	}
	return b.String()
}
