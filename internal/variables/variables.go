package variables

import (
	"fmt"
	"sort"
	"strings"
)

func Variables(vars map[string]string) {
	keys := make([]string, 0, len(vars))
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	printedMaps := ""
	for _, value := range keys {
		if strings.ContainsAny(value, ".") {
			splitted := strings.Split(value, ".")
			if !strings.Contains(printedMaps, splitted[0]) {
				printedMaps += splitted[0]
				println(splitted[0] + "")
			}
			println("\t" + strings.Join(splitted[1:], ".") + ": " + vars[value])
		} else {
			fmt.Printf("%s: %s\n", value, vars[value])
		}
		print("")
	}
}
