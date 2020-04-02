package website

import (
	"fmt"
	"strings"
)

type HttpError struct {
	err interface{}
	what string
}

func (h HttpError) Error() string {
	return fmt.Sprintf("%s\n%s\n%v\n", h.what,strings.Repeat("-",len(h.what)),h.err)
}
