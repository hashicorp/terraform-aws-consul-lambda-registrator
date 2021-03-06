package main

import (
	"os"
	"testing"

	testsuite "github.com/hashicorp/terraform-aws-consul-lambda-registrator/test/acceptance/framework/suite"
)

var suite testsuite.Suite

func TestMain(m *testing.M) {
	suite = testsuite.NewSuite(m)
	os.Exit(suite.Run())
}
