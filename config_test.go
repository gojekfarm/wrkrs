package wrkrs

import (
	"github.com/customerio/gospec"
	. "github.com/customerio/gospec"
)

func ConfigSpec(c gospec.Context) {
	var recoverOnPanic = func(f func()) (err interface{}) {
		defer func() {
			if cause := recover(); cause != nil {
				err = cause
			}
		}()

		f()

		return
	}

	c.Specify("sets redis pool size which defaults to 1", func() {
		c.Expect(Config.Pool.MaxIdle, Equals, 1)

		Configure(map[string]string{
			"server":  "localhost:6379",
			"process": "1",
			"pool":    "20",
		})

		c.Expect(Config.Pool.MaxIdle, Equals, 20)
	})

	c.Specify("can specify custom process", func() {
		c.Expect(Config.processId, Equals, "1")

		Configure(map[string]string{
			"server":  "localhost:6379",
			"process": "2",
		})

		c.Expect(Config.processId, Equals, "2")
	})

	c.Specify("requires a server parameter", func() {
		err := recoverOnPanic(func() {
			Configure(map[string]string{"process": "2"})
		})

		c.Expect(err, Equals, "Configure requires a 'server' option, which identifies a Redis instance")
	})

	c.Specify("requires a process parameter", func() {
		err := recoverOnPanic(func() {
			Configure(map[string]string{"server": "localhost:6379"})
		})

		c.Expect(err, Equals, "Configure requires a 'process' option, which uniquely identifies this instance")
	})

	c.Specify("adds ':' to the end of the namespace", func() {
		c.Expect(Config.Namespace, Equals, "")

		Configure(map[string]string{
			"server":    "localhost:6379",
			"process":   "1",
			"namespace": "prod",
		})

		c.Expect(Config.Namespace, Equals, "prod:")
	})

	c.Specify("defaults poll interval to 15 seconds", func() {
		Configure(map[string]string{
			"server":  "localhost:6379",
			"process": "1",
		})

		c.Expect(Config.PollInterval, Equals, 15)
	})

	c.Specify("allows customization of poll interval", func() {
		Configure(map[string]string{
			"server":        "localhost:6379",
			"process":       "1",
			"poll_interval": "1",
		})

		c.Expect(Config.PollInterval, Equals, 1)
	})

	c.Specify("defaults dial timeout to 500 ms", func() {
		Configure(map[string]string{
			"server":        "localhost:6379",
			"process":       "1",
			"poll_interval": "1",
		})

		c.Expect(Config.DialTimeout, Equals, 500)
	})

	c.Specify("allows customisation of dial timeout", func() {
		Configure(map[string]string{
			"server":        "localhost:6379",
			"process":       "1",
			"poll_interval": "1",
			"dial_timeout":  "100",
		})

		c.Expect(Config.DialTimeout, Equals, 100)
	})

	c.Specify("defaults read timeout to 500 ms", func() {
		Configure(map[string]string{
			"server":        "localhost:6379",
			"process":       "1",
			"poll_interval": "1",
		})

		c.Expect(Config.ReadTimeout, Equals, 500)
	})

	c.Specify("allows customisation of read timeout", func() {
		Configure(map[string]string{
			"server":        "localhost:6379",
			"process":       "1",
			"poll_interval": "1",
			"read_timeout":  "100",
		})

		c.Expect(Config.ReadTimeout, Equals, 100)
	})

	c.Specify("defaults write timeout to 500 ms", func() {
		Configure(map[string]string{
			"server":        "localhost:6379",
			"process":       "1",
			"poll_interval": "1",
		})

		c.Expect(Config.WriteTimeout, Equals, 500)
	})

	c.Specify("allows customisation of write timeout", func() {
		Configure(map[string]string{
			"server":        "localhost:6379",
			"process":       "1",
			"poll_interval": "1",
			"write_timeout": "100",
		})

		c.Expect(Config.WriteTimeout, Equals, 100)
	})
}
