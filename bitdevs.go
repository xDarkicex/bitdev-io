package bitdevs

import (
		"fmt"
		"time"

		"github.com/labstack/echo/v4"
		"github.com/xDarkicex/bitdevs/middleware"
)

type BitDev struct {
		Author string
		Date time.Time
		Domain string
}

// NewBitDev Control struct for bitdevs helper package
func NewBitDev(author, domain string) *BitDev {
		return &BitDev{Author: author, Domain: domain, Date: time.Now()}
}

func (bitDev BitDev) Secure() echo.MiddlewareFunc {
		fmt.Println("Route Secured")
		return  Secure()
}

var (
		// Secure Adds secure headers
		Secure = middleware.SecureHeader
)
