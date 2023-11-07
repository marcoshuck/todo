package server

import (
	"github.com/gojaguar/jaguar/config"
	"github.com/marcoshuck/todo/pkg/conf"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
}

func (suite *ServerTestSuite) SetupSuite() {
	suite.Require().NoError(os.Setenv("APPLICATION_NAME", "todo"))
	suite.Require().NoError(os.Setenv("DATABASE_NAME", "todo_db"))
}

func (suite *ServerTestSuite) SetupTest() {

}

func (suite *ServerTestSuite) TearDownTest() {

}

func (suite *ServerTestSuite) TearDownSuite() {
	suite.Require().NoError(os.Unsetenv("APPLICATION_NAME"))
	suite.Require().NoError(os.Unsetenv("DATABASE_NAME"))
	suite.Require().NoError(os.Remove("todo_db.db"))
}

func (suite *ServerTestSuite) TestSetup() {
	cfg, err := conf.ReadServerConfig()
	suite.Require().NoError(err)
	suite.Require().NotZero(cfg)

	cfg.DB.Engine = config.EngineSQLite

	app, err := Setup(cfg)
	suite.Assert().NoError(err)
	suite.Assert().NotZero(app)
	suite.Assert().NotNil(app.listener)
	suite.Assert().NotNil(app.logger)
	suite.Assert().NotNil(app.server)
	suite.Assert().NotNil(app.db)
	suite.Assert().NotNil(app.services.Tasks)
}
