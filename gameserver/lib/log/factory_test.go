package log_test

import (
	"testing"

	"github.com/heatxsink/go-logstash"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	l := logstash.New("192.168.1.103", 5000, 5)

	_, err := l.Connect()

	// log.Println(tcp.)

	assert.NoError(t, err)

	err = l.Writeln("{ 'foo' : 'bar' }")

	assert.NoError(t, err)

	// err := l.Connect()
}
