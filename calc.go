package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

var floatType = reflect.TypeOf(float64(0))

// Add adds a number to a variable
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

// Sub subtracts a number from a variable
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

// Mul multiplies the value of the variable by a number
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

// Div divides the variable by a number
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

// Sqrt gets the square root of a variable and sets it to the variable
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

// getFloat converts to float
func getFloat(unk interface{}) float64 {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0
	}
	fv := v.Convert(floatType)
	return fv.Float()
}
