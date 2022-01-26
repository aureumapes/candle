// By AureumApes

package main

import (
	"fmt"
	"math"
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

func Sqrt(command string) {
	splitedCommand := strings.Split(command, " ")
	_, ok := variables[splitedCommand[1]]
	if ok {
		if isNumeric(splitedCommand[2]) {
			a, _ := strconv.ParseFloat(splitedCommand[2], 64)
			variables[splitedCommand[1]] = math.Sqrt(a)
		} else {
			variables[splitedCommand[1]] = math.Sqrt(getFloat(variables[splitedCommand[2]]))
		}
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
