package limit

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//Limit function is gin middleware to limit current requests
func Limit(max int) gin.HandlerFunc {
	if max <= 0 {
		log.Panic("max must be more than 0")
	}
	sema := make(chan struct{}, max)
	return func(c *gin.Context) {
		var called, fulled bool
		defer func() {
			if called == false && fulled == false {
				<-sema
			}
			if r := recover(); r != nil { // We don't handle panic
				panic(r)
			}
		}()

		select {
		case sema <- struct{}{}:
			c.Next()
			called = true
			<-sema
		default:
			fulled = true
			c.Status(http.StatusBadGateway)
		}
	}
}
