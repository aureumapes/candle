package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func Add(command string) {
	splitCmd := strings.Split(command, " ")
	variableOld, ok := variables[splitCmd[2]]
	if !ok {
		fmt.Println(splitCmd[2] + ": Unknown Variable")
	}
	if isNumeric(splitCmd[1]) {
		variable, _ := strconv.ParseFloat(splitCmd[1], 64)
		variables[splitCmd[2]] = getFloat(variableOld) + variable
	}
}

func Sub(command string) {
	splitCmd := strings.Split(command, " ")
	variableOld, ok := variables[splitCmd[2]]
	if !ok {
		fmt.Println(splitCmd[2] + ": Unknown Variable")
	}
	if isNumeric(splitCmd[1]) {
		variable, _ := strconv.ParseFloat(splitCmd[1], 64)
		variables[splitCmd[2]] = getFloat(variableOld) - variable

	}
}

func Mul(command string) {
	splitCmd := strings.Split(command, " ")
	variableOld, ok := variables[splitCmd[2]]
	if !ok {
		fmt.Println(splitCmd[2] + ": Unknown Variable")
	}
	if isNumeric(splitCmd[1]) {
		variable, _ := strconv.ParseFloat(splitCmd[1], 64)
		variables[splitCmd[2]] = getFloat(variableOld) * variable

	}
}

func Div(command string) {
	splitCmd := strings.Split(command, " ")
	variableOld, ok := variables[splitCmd[2]]
	if !ok {
		fmt.Println(splitCmd[2] + ": Unknown Variable")
	}
	if isNumeric(splitCmd[1]) {
		variable, _ := strconv.ParseFloat(splitCmd[1], 64)
		variables[splitCmd[2]] = getFloat(variableOld) / variable
	}
}

func Sqrt(command string) {
	splitCmd := strings.Split(command, " ")
	_, ok := variables[splitCmd[1]]
	if ok {
		if isNumeric(splitCmd[2]) {
			a, _ := strconv.ParseFloat(splitCmd[2], 64)
			variables[splitCmd[1]] = math.Sqrt(a)
		} else {
			variables[splitCmd[1]] = math.Sqrt(getFloat(variables[splitCmd[2]]))
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
