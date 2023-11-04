package server

import (
	"github.com/gojaguar/jaguar/config"
	"github.com/marcoshuck/todo/pkg/conf"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"os"
	"testing"
)

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
	cancelSetEnvAppName func()
	cancelSetEnvDBName  func()
}

func (suite *ServerTestSuite) SetupSuite() {
	var err error
	err, suite.cancelSetEnvAppName = conf.setEnv("APPLICATION_NAME", "todo")
	suite.Require().NoError(err)

	err, suite.cancelSetEnvDBName = conf.setEnv("DATABASE_NAME", "todo_db")
	suite.Require().NoError(err)
}

func (suite *ServerTestSuite) SetupTest() {

}

func (suite *ServerTestSuite) TearDownTest() {

}

func (suite *ServerTestSuite) TearDownSuite() {
	suite.cancelSetEnvAppName()
	suite.cancelSetEnvDBName()
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

func (suite *ServerTestSuite) TestRun() {
	var srv testGrpcServer

	suite.Assert().NoError(Run(Application{
		server:   &srv,
		logger:   zap.NewNop(),
		listener: bufconn.Listen(1),
	}))
	suite.Assert().Equal(testGrpcServer(1), srv)
}

type testGrpcServer int

func (t *testGrpcServer) Serve(listener net.Listener) error {
	*t++
	return nil
}
