package thaicurrency

import (
	"errors"
	"math"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	ErrTooManyDecimals   = errors.New("decimal point more than 2 digits")
	ErrInputExceedsLimit = errors.New("input exceeds maximum limit")

	maxInt64 = decimal.NewFromInt(math.MaxInt64)

	digitWords = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า", "สิบ"}
	placeWords = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}
)

func Decimal(amount decimal.Decimal) (string, error) {
	if err := validateAmountLimit(amount); err != nil {
		return "", err
	}
	formattedAmount := amount.StringFixed(2)
	return convert(formattedAmount)
}

func validateAmountLimit(amount decimal.Decimal) error {
	if amount.GreaterThan(maxInt64) {
		return ErrInputExceedsLimit
	}
	return nil
}

func convert(amount string) (string, error) {
	bahtPart, satangPart := splitAmount(amount)

	if !isValidSatang(satangPart) {
		return "", ErrTooManyDecimals
	}

	bahtText := formatBaht(bahtPart)
	satangText := formatSatang(satangPart)

	return buildThaiString(bahtText, satangText), nil
}

func splitAmount(amount string) (string, string) {
	index := strings.Index(amount, ".")
	if index == -1 {
		return amount, ""
	}

	satang := amount[index+1:]
	if len(satang) == 1 {
		satang += "0"
	}

	return amount[:index], satang
}

func isValidSatang(satang string) bool {
	return len(satang) == 2 || len(satang) == 0
}

func splitIntoMillionGroups(baht string) []string {
	digits := strings.Split(baht, "")
	numGroups := len(digits)/6 + 1
	groups := make([]string, numGroups)
	groupIdx := numGroups - 1

	for i := len(digits) - 1; i >= 0; i-- {
		groups[groupIdx] = digits[i] + groups[groupIdx]
		if len(groups[groupIdx]) >= 6 {
			groupIdx--
		}
	}

	return groups
}

func groupToThaiText(group string) string {
	var result string

	for j := 0; j < len(group); j++ {
		digit := int(group[j] - '0')
		if digit == 0 {
			continue
		}

		position := len(group) - j - 1

		switch {
		case position == 0 && digit == 1 && len(group) > 1:
			result += "เอ็ด"
		case position == 1 && digit == 2:
			result += "ยี่"
		case position == 1 && digit == 1:
		default:
			result += digitWords[digit]
		}

		result += placeWords[position]
	}

	return result
}

func formatBaht(baht string) string {
	groups := splitIntoMillionGroups(baht)
	textParts := make([]string, len(groups))

	for i, group := range groups {
		textParts[i] = groupToThaiText(group)
	}

	return strings.Join(textParts, "ล้าน")
}

func formatSatang(satang string) string {
	var result string

	for i := 0; i < len(satang); i++ {
		digit := int(satang[i] - '0')
		if digit == 0 {
			continue
		}

		position := len(satang) - i - 1

		switch {
		case position == 0 && digit == 1:
			result += "เอ็ด"
		case position == 1 && digit == 2:
			result += "ยี่"
		case position == 1 && digit == 1:
		default:
			result += digitWords[digit]
		}

		result += placeWords[position]
	}

	if result == "" {
		return "ถ้วน"
	}

	if result == "เอ็ด" {
		return "หนึ่งสตางค์"
	}

	return result + "สตางค์"
}

func buildThaiString(bahtText, satangText string) string {
	if bahtText == "" && satangText == "ถ้วน" {
		return "ศูนย์บาทถ้วน"
	}
	if bahtText == "" {
		return satangText
	}
	return bahtText + "บาท" + satangText
}
