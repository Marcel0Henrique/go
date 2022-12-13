package env

import (
	"os"
)

// Salva as varias de ambiente
func Setenv() {
	os.Setenv("mysqlUser", "root")
	os.Setenv("mysqlPass", "1275")
}
