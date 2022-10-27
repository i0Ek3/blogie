package api

import (
	"expvar"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Expvar(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	first := true
	report := func(key string, value any) {
		if !first {
			_, err := fmt.Fprintf(c.Writer, ",\n")
			if err != nil {
				return
			}
		}
		first = false
		if str, ok := value.(string); ok {
			_, err := fmt.Fprintf(c.Writer, "%q: %q", key, str)
			if err != nil {
				return
			}
		} else {
			_, err := fmt.Fprintf(c.Writer, "%q: %v", key, value)
			if err != nil {
				return
			}
		}
	}

	_, err := fmt.Fprintf(c.Writer, "{\n")
	if err != nil {
		return
	}

	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	
	_, err = fmt.Fprintf(c.Writer, "\n}\n")
	if err != nil {
		return
	}
}
