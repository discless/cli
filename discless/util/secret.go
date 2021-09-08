package util

import (
	"errors"
	"strings"
)

func ReplaceSecret(in string, secrets map[string]string) (string, error) {
	if strings.HasPrefix(in,"secret.") {
		secretk := (in)[7:]
		if v,ok := secrets[secretk]; ok {
			return v, nil
		} else {
			return in, errors.New("Could not find secret \"" + string(secretk) + "\" in the provided secrets.")
		}
	}
	return in, nil
}
