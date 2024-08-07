// @generated by protoc-gen-es v2.0.0
// @generated from file api/tasks/v1/tasks.proto (package api.tasks.v1, syntax proto3)
/* eslint-disable */

import type {GenFile, GenMessage, GenService} from "@bufbuild/protobuf/codegenv1";
import type {Message} from "@bufbuild/protobuf";
import type {FieldMask, Timestamp} from "@bufbuild/protobuf/wkt";

/**
 * Describes the file api/tasks/v1/tasks.proto.
 */
export declare const file_api_tasks_v1_tasks: GenFile;

/**
 * @generated from message api.tasks.v1.Task
 */
export declare type Task = Message<"api.tasks.v1.Task"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: string title = 2;
   */
  title: string;

  /**
   * @generated from field: string description = 3;
   */
  description: string;

  /**
   * @generated from field: google.protobuf.Timestamp deadline = 4;
   */
  deadline?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp completed_at = 5;
   */
  completedAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp create_time = 1000;
   */
  createTime?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp update_time = 1001;
   */
  updateTime?: Timestamp;
};

/**
 * Describes the message api.tasks.v1.Task.
 * Use `create(TaskSchema)` to create a new message.
 */
export declare const TaskSchema: GenMessage<Task>;

/**
 * @generated from message api.tasks.v1.CreateTaskRequest
 */
export declare type CreateTaskRequest = Message<"api.tasks.v1.CreateTaskRequest"> & {
  /**
   * Task is the the task to create.
   *
   * @generated from field: api.tasks.v1.Task task = 2;
   */
  task?: Task;
};

/**
 * Describes the message api.tasks.v1.CreateTaskRequest.
 * Use `create(CreateTaskRequestSchema)` to create a new message.
 */
export declare const CreateTaskRequestSchema: GenMessage<CreateTaskRequest>;

/**
 * @generated from message api.tasks.v1.DeleteTaskRequest
 */
export declare type DeleteTaskRequest = Message<"api.tasks.v1.DeleteTaskRequest"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;
};

/**
 * Describes the message api.tasks.v1.DeleteTaskRequest.
 * Use `create(DeleteTaskRequestSchema)` to create a new message.
 */
export declare const DeleteTaskRequestSchema: GenMessage<DeleteTaskRequest>;

/**
 * @generated from message api.tasks.v1.UndeleteTaskRequest
 */
export declare type UndeleteTaskRequest = Message<"api.tasks.v1.UndeleteTaskRequest"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;
};

/**
 * Describes the message api.tasks.v1.UndeleteTaskRequest.
 * Use `create(UndeleteTaskRequestSchema)` to create a new message.
 */
export declare const UndeleteTaskRequestSchema: GenMessage<UndeleteTaskRequest>;

/**
 * @generated from message api.tasks.v1.UpdateTaskRequest
 */
export declare type UpdateTaskRequest = Message<"api.tasks.v1.UpdateTaskRequest"> & {
  /**
   * @generated from field: api.tasks.v1.Task task = 1;
   */
  task?: Task;

  /**
   * @generated from field: google.protobuf.FieldMask update_mask = 2;
   */
  updateMask?: FieldMask;
};

/**
 * Describes the message api.tasks.v1.UpdateTaskRequest.
 * Use `create(UpdateTaskRequestSchema)` to create a new message.
 */
export declare const UpdateTaskRequestSchema: GenMessage<UpdateTaskRequest>;

/**
 * @generated from message api.tasks.v1.GetTaskRequest
 */
export declare type GetTaskRequest = Message<"api.tasks.v1.GetTaskRequest"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;
};

/**
 * Describes the message api.tasks.v1.GetTaskRequest.
 * Use `create(GetTaskRequestSchema)` to create a new message.
 */
export declare const GetTaskRequestSchema: GenMessage<GetTaskRequest>;

/**
 * @generated from message api.tasks.v1.ListTasksRequest
 */
export declare type ListTasksRequest = Message<"api.tasks.v1.ListTasksRequest"> & {
  /**
   * @generated from field: int32 page_size = 2;
   */
  pageSize: number;

  /**
   * @generated from field: string page_token = 3;
   */
  pageToken: string;
};

/**
 * Describes the message api.tasks.v1.ListTasksRequest.
 * Use `create(ListTasksRequestSchema)` to create a new message.
 */
export declare const ListTasksRequestSchema: GenMessage<ListTasksRequest>;

/**
 * @generated from message api.tasks.v1.ListTasksResponse
 */
export declare type ListTasksResponse = Message<"api.tasks.v1.ListTasksResponse"> & {
  /**
   * @generated from field: repeated api.tasks.v1.Task tasks = 1;
   */
  tasks: Task[];

  /**
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;
};

/**
 * Describes the message api.tasks.v1.ListTasksResponse.
 * Use `create(ListTasksResponseSchema)` to create a new message.
 */
export declare const ListTasksResponseSchema: GenMessage<ListTasksResponse>;

/**
 * TasksWriterService holds the methods to persist, modify and remove Tasks.
 *
 * @generated from service api.tasks.v1.TasksWriterService
 */
export declare const TasksWriterService: GenService<{
  /**
   * CreateTask creates a Task.
   *
   * @generated from rpc api.tasks.v1.TasksWriterService.CreateTask
   */
  createTask: {
    methodKind: "unary";
    input: typeof CreateTaskRequestSchema;
    output: typeof TaskSchema;
  },
  /**
   * @generated from rpc api.tasks.v1.TasksWriterService.DeleteTask
   */
  deleteTask: {
    methodKind: "unary";
    input: typeof DeleteTaskRequestSchema;
    output: typeof TaskSchema;
  },
  /**
   * @generated from rpc api.tasks.v1.TasksWriterService.UndeleteTask
   */
  undeleteTask: {
    methodKind: "unary";
    input: typeof UndeleteTaskRequestSchema;
    output: typeof TaskSchema;
  },
  /**
   * @generated from rpc api.tasks.v1.TasksWriterService.UpdateTask
   */
  updateTask: {
    methodKind: "unary";
    input: typeof UpdateTaskRequestSchema;
    output: typeof TaskSchema;
  },
}>;

/**
 * TasksReaderService holds the methods to obtain Tasks.
 *
 * @generated from service api.tasks.v1.TasksReaderService
 */
export declare const TasksReaderService: GenService<{
  /**
   * GetTask returns a Task.
   *
   * @generated from rpc api.tasks.v1.TasksReaderService.GetTask
   */
  getTask: {
    methodKind: "unary";
    input: typeof GetTaskRequestSchema;
    output: typeof TaskSchema;
  },
  /**
   * ListTasks returns a list of Tasks.
   *
   * @generated from rpc api.tasks.v1.TasksReaderService.ListTasks
   */
  listTasks: {
    methodKind: "unary";
    input: typeof ListTasksRequestSchema;
    output: typeof ListTasksResponseSchema;
  },
}>;

