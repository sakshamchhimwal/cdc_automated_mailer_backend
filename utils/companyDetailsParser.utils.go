package utils

import (
	"math"
	"strings"
	"unicode"
)

func countSpecialChars(line string) float32 {
	var spCharCount = 0
	for i := 0; i < len(line); i++ {
		if line[i] != ' ' && !unicode.IsLetter(rune(line[i])) {
			spCharCount += 1
		}
	}
	return float32(spCharCount)
}

func CompanyDetailsParser(companyText string) string {
	companyText = strings.Replace(companyText, "\n", "", -1)
	companyText = strings.Replace(companyText, "\t", "", -1)
	companyText = strings.Replace(companyText, ".", "\n", -1)
	companyText = strings.Replace(companyText, "  ", "", -1)
	splitString := strings.Split(companyText, "\n")
	var parsedString = ""

	var traversalLen = math.Min(float64(len(splitString)), 200)

	for i := 0; i < int(traversalLen); i++ {
		var tempString = ""
		for j := 0; j < len(splitString[i]); j++ {
			if len(tempString) == 0 {
				tempString += string(splitString[i][j])
			} else {
				if tempString[len(tempString)-1] == ' ' && splitString[i][j] != ' ' {
					tempString += string(splitString[i][j])
				} else if tempString[len(tempString)-1] != ' ' {
					tempString += string(splitString[i][j])
				}
			}
		}
		if len(tempString) > 30 {
			if countSpecialChars(tempString)/float32(len(tempString)) < 0.15 {
				parsedString += "\n" + tempString
			}
		}
	}

	return parsedString

}
