package service

import (
	"context"
	tasksv1 "github.com/marcoshuck/todo/api/tasks/v1"
	"github.com/marcoshuck/todo/pkg/domain"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/metric/noop"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	reader tasksv1.TasksReaderServiceServer
}

func (suite *TasksServiceTestSuite) SetupSuite() {

}

func (suite *TasksServiceTestSuite) SetupTest() {
	var err error
	suite.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	suite.Require().NoError(err)
	suite.Require().NoError(suite.db.Migrator().AutoMigrate(&domain.Task{}))

	suite.writer = NewTasksWriter(suite.db, zap.NewNop(), noop.NewMeterProvider().Meter(""))
	suite.reader = NewTasksReader(suite.db, zap.NewNop(), noop.NewMeterProvider().Meter(""))
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

func (suite *TasksServiceTestSuite) TestGet_Success() {
	ctx := context.Background()

	expected, err := suite.writer.CreateTask(ctx, &tasksv1.CreateTaskRequest{
		Task: &tasksv1.Task{
			Title: "A test",
		},
	})
	suite.Require().NoError(err)

	response, err := suite.reader.GetTask(ctx, &tasksv1.GetTaskRequest{Id: expected.GetId()})
	suite.Assert().NoError(err)

	suite.Assert().Equal(expected.GetTitle(), response.GetTitle())
}

func (suite *TasksServiceTestSuite) TestGet_NotFound() {
	ctx := context.Background()

	_, err := suite.reader.GetTask(ctx, &tasksv1.GetTaskRequest{Id: 199452})
	suite.Assert().Error(err)
	suite.Assert().ErrorIs(err, status.Error(codes.NotFound, "task not found"))
}

func (suite *TasksServiceTestSuite) TestList_Empty() {
	ctx := context.Background()

	response, err := suite.reader.ListTasks(ctx, &tasksv1.ListTasksRequest{
		PageSize:  0,
		PageToken: "",
	})
	suite.Assert().NoError(err)
	suite.Assert().Empty(response.GetTasks())
}

func (suite *TasksServiceTestSuite) TestList_Success() {
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		_, err := suite.writer.CreateTask(ctx, &tasksv1.CreateTaskRequest{
			Task: &tasksv1.Task{
				Title: "A test",
			},
		})
		suite.Require().NoError(err)
	}

	var expected int64
	suite.Require().NoError(suite.db.Model(&domain.Task{}).Count(&expected).Error)

	response, err := suite.reader.ListTasks(ctx, &tasksv1.ListTasksRequest{
		PageSize:  0,
		PageToken: "",
	})
	suite.Assert().NoError(err)
	suite.Assert().NotEmpty(response.GetTasks())
	suite.Assert().Len(response.GetTasks(), int(expected))
}
