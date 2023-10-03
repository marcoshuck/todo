package server

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
}

func (suite *ServerTestSuite) SetupSuite() {

}

func (suite *ServerTestSuite) SetupTest() {

}

func (suite *ServerTestSuite) TearDownTest() {

}

func (suite *ServerTestSuite) TearDownSuite() {

}

func (suite *ServerTestSuite) TestReadConfig() {
	cfg, err := ReadConfig()
	suite.Assert().NoError(err)
	suite.Assert().NotZero(cfg)
}

func (suite *ServerTestSuite) TestSetup() {

}
