package tratamento

import (
	"strings"
)

//RemoveCaracteres limpa os caracteres da coluna
func RemoveCaracteres(data string) string {
    data = strings.Replace(data, ".", "", -1)
    data = strings.Replace(data, "-", "", -1)
    data = strings.Replace(data, "/", "", -1)
    return data
}