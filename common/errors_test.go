package common_test

import (
	"bytes"
	"errors"
	"github.com/bouk/monkey"
	"github.com/inhuman/helpers/common"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

const layout = "2006/01/02 15:04:05"

func TestCheckErrorFatal(t *testing.T) {

	err := getError()

	fakeLogFatal := func(msg ...interface{}) {
		log.Println("fake log fatal")
	}

	patch := monkey.Patch(log.Fatal, fakeLogFatal)
	defer patch.Unpatch()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	common.CheckErrorFatal(err)

	expect := time.Now().Format(layout) + " fake log fatal\n"

	assert.Equal(t, expect, buf.String())

}

func TestCheckErrorFatalNoError(t *testing.T) {

	err := getNilError()
	test := true

	common.CheckErrorFatal(err)

	assert.Equal(t, true, test)

}

func TestCheckErrorMessage(t *testing.T) {

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	err := getError()

	isError := common.CheckErrorMessage("error fired:", err)

	expect := time.Now().Format(layout) + " error fired: test error\n"
	assert.Equal(t, expect, buf.String())
	assert.Equal(t, true, isError)

}

func TestCheckErrorMessageNoError(t *testing.T) {

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	err := getNilError()

	isError := common.CheckErrorMessage("not printed", err)

	assert.Equal(t, "", buf.String())
	assert.Equal(t, false, isError)

}

func TestCheckError(t *testing.T) {

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	err := getError()

	isError := common.CheckError(err)

	expect := time.Now().Format(layout) + " error: test error\n"
	assert.Equal(t, expect, buf.String())
	assert.Equal(t, true, isError)

}

func TestCheckErrorNoError(t *testing.T) {

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	err := getNilError()

	isError := common.CheckError(err)

	assert.Equal(t, "", buf.String())
	assert.Equal(t, false, isError)
}

func getNilError() error {
	return nil
}

func getError() error {
	return errors.New("test error")
}
