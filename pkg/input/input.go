package input

import (
	"bufio"
	"os"
	"strings"
)

// Вводим данные для структур
func Input() []string {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	inp := strings.Split(s.Text(), " ")
	return inp
}
