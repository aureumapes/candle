// By AureumApes

package commands

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Add(command string) {
	splitedCmd := strings.Split(command, " ")
	variableOld, ok := variables[splitedCmd[2]]
	if !ok {
		fmt.Println(splitedCmd[2] + ": Unknown Variable")
	}
	if isNumeric(splitedCmd[1]) {
		variable, _ := strconv.ParseFloat(splitedCmd[1], 64)
		variables[splitedCmd[2]] = getFloat(variableOld) + variable

	}
}

func Sub(command string) {
	splitedCmd := strings.Split(command, " ")
	variableOld, ok := variables[splitedCmd[2]]
	if !ok {
		fmt.Println(splitedCmd[2] + ": Unknown Variable")
	}
	if isNumeric(splitedCmd[1]) {
		variable, _ := strconv.ParseFloat(splitedCmd[1], 64)
		variables[splitedCmd[2]] = getFloat(variableOld) - variable

	}
}

func Mult(command string) {
	splitedCmd := strings.Split(command, " ")
	variableOld, ok := variables[splitedCmd[2]]
	if !ok {
		fmt.Println(splitedCmd[2] + ": Unknown Variable")
	}
	if isNumeric(splitedCmd[1]) {
		variable, _ := strconv.ParseFloat(splitedCmd[1], 64)
		variables[splitedCmd[2]] = getFloat(variableOld) * variable

	}
}

func Div(command string) {
	splitedCmd := strings.Split(command, " ")
	variableOld, ok := variables[splitedCmd[2]]
	if !ok {
		fmt.Println(splitedCmd[2] + ": Unknown Variable")
	}
	if isNumeric(splitedCmd[1]) {
		variable, _ := strconv.ParseFloat(splitedCmd[1], 64)
		variables[splitedCmd[2]] = getFloat(variableOld) / variable
	}
}

var floatType = reflect.TypeOf(float64(0))

func getFloat(unk interface{}) float64 {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0
	}
	fv := v.Convert(floatType)
	return fv.Float()
}
