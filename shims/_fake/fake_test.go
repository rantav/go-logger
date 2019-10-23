package fake

import (
	"github.com/AppsFlyer/go-logger"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("meets the interface", func() {
	var _ log.Logger = &FakeLogger{}
})
