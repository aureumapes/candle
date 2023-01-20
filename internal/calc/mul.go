package calc

import (
	"github.com/Amiraxoba/candle/util"
	"strconv"
)

func Mul(args []string, vars map[string]string) map[string]string {
	if _, varExists := vars[args[0]]; varExists {
		if util.IsNumeric(vars[args[0]]) {
			if util.IsNumeric(args[1]) {
				convertedVar, _ := strconv.ParseFloat(vars[args[0]], 64)
				convertedMultiplication, _ := strconv.ParseFloat(args[1], 64)
				vars[args[0]] = strconv.FormatFloat(convertedVar*convertedMultiplication, 'f', -1, 64)
			} else if _, varExists = vars[args[1]]; varExists {
				convertedVar, _ := strconv.ParseFloat(vars[args[0]], 64)
				convertedMultiplication, _ := strconv.ParseFloat(vars[args[1]], 64)
				vars[args[0]] = strconv.FormatFloat(convertedVar*convertedMultiplication, 'f', -1, 64)
			} else {
				util.PrintError("\"" + args[1] + "\" is not numeric, nor an variable")
			}
		} else {
			util.PrintError("Variable \"" + args[0] + "\" is not numeric")
		}
	} else {
		util.PrintError("Unknown variable \"" + args[0] + "\"")
	}
	return vars
}
