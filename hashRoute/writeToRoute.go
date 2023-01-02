package hashRoute

import (
	"fmt"
	"os"
	"strings"
)

func WriteToRoute2(apiGroup string, name string, method string, controller string) {
	data, err := os.ReadFile("app/routes/router.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(data)

	// Convert the file data to a string
	str := string(data)

	// Find the start and end indices of the "registration" function
	startIndex := strings.Index(str, apiGroup)
	endIndex := strings.Index(str, "}")

	// Extract the function body as a string
	functionBody := str[startIndex : endIndex+1]

	// Find the index of the first curly brace in the function body
	braceIndex := strings.Index(functionBody, ")")

	// Modify the function body by inserting the new code after the first curly brace
	toReplace := "[APIGROUP].[METHOD](\"/[NAME]\", [CONTROLLER])"
	toReplace = strings.Replace(toReplace, "[APIGROUP]", apiGroup, 1)
	toReplace = strings.Replace(toReplace, "[METHOD]", method, 1)
	toReplace = strings.Replace(toReplace, "[NAME]", name, 1)
	toReplace = strings.Replace(toReplace, "[CONTROLLER]", controller, 1)
	functionBody = functionBody[:braceIndex+1] + "\n\t" + toReplace + functionBody[braceIndex+1:]

	// Replace the function body in the original string
	str = str[:startIndex] + functionBody + str[endIndex+1:]

	// Convert the modified string back to a byte slice
	modifiedData := []byte(str)

	// Write the modified data back to the file
	err = os.WriteFile("app/routes/router.go", modifiedData, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
