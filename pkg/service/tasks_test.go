package service

import (
	"context"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/domain"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestTasksServiceSuite(t *testing.T) {
	suite.Run(t, new(TasksServiceTestSuite))
}

type TasksServiceTestSuite struct {
	suite.Suite
	db  *gorm.DB
	svc tasksv1.TasksServiceServer
}

func (suite *TasksServiceTestSuite) SetupSuite() {

}

func (suite *TasksServiceTestSuite) SetupTest() {
	var err error
	suite.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	suite.Require().NoError(err)
	suite.Require().NoError(suite.db.Migrator().AutoMigrate(&domain.Task{}))

	suite.svc = NewTasks(
		suite.db,
		zap.NewNop(),
		trace.NewNoopTracerProvider().Tracer(""),
		noop.NewMeterProvider().Meter(""),
	)
}

func (suite *TasksServiceTestSuite) TearDownTest() {
	db, err := suite.db.DB()
	suite.Require().NoError(err)
	suite.Assert().NoError(db.Close())
}

func (suite *TasksServiceTestSuite) TearDownSuite() {

}

func (suite *TasksServiceTestSuite) TestCreate_NilBody() {
	res, err := suite.svc.CreateTask(context.Background(), &tasksv1.CreateTaskRequest{
		Task: nil,
	})
	suite.Assert().Error(err)
	suite.Assert().ErrorContains(err, "invalid task")
	suite.Assert().Nil(res)
}

func (suite *TasksServiceTestSuite) TestCreate_InvalidArgument() {
	res, err := suite.svc.CreateTask(context.Background(), &tasksv1.CreateTaskRequest{
		Task: &tasksv1.Task{
			Title: "",
		},
	})
	suite.Assert().Error(err)
	suite.Assert().ErrorContains(err, "invalid task")
	suite.Assert().Nil(res)
}

func (suite *TasksServiceTestSuite) TestCreate_Success() {
	var before int64
	suite.Require().NoError(suite.db.Model(&domain.Task{}).Count(&before).Error)

	const title = "test"
	res, err := suite.svc.CreateTask(context.Background(), &tasksv1.CreateTaskRequest{
		Task: &tasksv1.Task{
			Title: title,
		},
	})
	suite.Assert().NoError(err)
	suite.Assert().NotNil(res)
	suite.Assert().Equal(title, res.GetTitle())

	var after int64
	suite.Require().NoError(suite.db.Model(&domain.Task{}).Count(&after).Error)
	suite.NotEqual(before, after)
	suite.Equal(before+1, after)
}
