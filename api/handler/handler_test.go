package handler

import (
	"log"
	"testing"

	"github.com/bufbuild/protovalidate-go"
)

var testValidator *protovalidate.Validator

func TestMain(m *testing.M) {
	var err error
	testValidator, err = protovalidate.New()
	if err != nil {
		log.Fatalf("failed to create validator: %s", err)
	}

	m.Run()
}
