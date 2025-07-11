package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
