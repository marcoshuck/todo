package service

import (
	"context"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/domain"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/metric/noop"
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
	db     *gorm.DB
	writer tasksv1.TasksWriterServiceServer
}

func (suite *TasksServiceTestSuite) SetupSuite() {

}

func (suite *TasksServiceTestSuite) SetupTest() {
	var err error
	suite.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	suite.Require().NoError(err)
	suite.Require().NoError(suite.db.Migrator().AutoMigrate(&domain.Task{}))

	suite.writer = NewTasksWriter(suite.db, zap.NewNop(), noop.NewMeterProvider().Meter(""))
}

func (suite *TasksServiceTestSuite) TearDownTest() {
	db, err := suite.db.DB()
	suite.Require().NoError(err)
	suite.Assert().NoError(db.Close())
}

func (suite *TasksServiceTestSuite) TearDownSuite() {

}

func (suite *TasksServiceTestSuite) TestCreate_Success() {
	var before int64
	suite.Require().NoError(suite.db.Model(&domain.Task{}).Count(&before).Error)

	const title = "test"
	res, err := suite.writer.CreateTask(context.Background(), &tasksv1.CreateTaskRequest{
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
