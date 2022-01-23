package commands

import (
	"strings"
	"fmt"
	"strconv"
	"reflect"
)

func Add(command string){
	splitedCmd := strings.Split(command, " ")
	variableOld, ok := variables[splitedCmd[3]]
	if splitedCmd[2] == "to"{
		if !ok {
			fmt.Println(splitedCmd[3] + ": Unknown Variable")
		}
		if isNumeric(splitedCmd[1]) {
			variable, _ := strconv.ParseFloat(splitedCmd[1], 64)
			variables[splitedCmd[3]] = getFloat(variableOld) + variable
		}
	}
}

var floatType = reflect.TypeOf(float64(0))

func getFloat(unk interface{}) (float64) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0
	}
	fv := v.Convert(floatType)
	return fv.Float()
}
