// @generated by protoc-gen-es v1.10.0
// @generated from file api/tasks/v1/tasks.proto (package api.tasks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import {FieldMask, proto3, Timestamp} from "@bufbuild/protobuf";

/**
 * @generated from message api.tasks.v1.Task
 */
export const Task = /*@__PURE__*/ proto3.makeMessageType(
  "api.tasks.v1.Task",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "deadline", kind: "message", T: Timestamp },
    { no: 5, name: "completed_at", kind: "message", T: Timestamp },
    { no: 1000, name: "create_time", kind: "message", T: Timestamp },
    { no: 1001, name: "update_time", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message api.tasks.v1.CreateTaskRequest
 */
export const CreateTaskRequest = /*@__PURE__*/ proto3.makeMessageType(
  "api.tasks.v1.CreateTaskRequest",
  () => [
    { no: 2, name: "task", kind: "message", T: Task },
  ],
);

/**
 * @generated from message api.tasks.v1.DeleteTaskRequest
 */
export const DeleteTaskRequest = /*@__PURE__*/ proto3.makeMessageType(
  "api.tasks.v1.DeleteTaskRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message api.tasks.v1.UndeleteTaskRequest
 */
export const UndeleteTaskRequest = /*@__PURE__*/ proto3.makeMessageType(
  "api.tasks.v1.UndeleteTaskRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message api.tasks.v1.UpdateTaskRequest
 */
export const UpdateTaskRequest = /*@__PURE__*/ proto3.makeMessageType(
  "api.tasks.v1.UpdateTaskRequest",
  () => [
    { no: 1, name: "task", kind: "message", T: Task },
    { no: 2, name: "update_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message api.tasks.v1.GetTaskRequest
 */
export const GetTaskRequest = /*@__PURE__*/ proto3.makeMessageType(
  "api.tasks.v1.GetTaskRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message api.tasks.v1.ListTasksRequest
 */
export const ListTasksRequest = /*@__PURE__*/ proto3.makeMessageType(
  "api.tasks.v1.ListTasksRequest",
  () => [
    { no: 2, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message api.tasks.v1.ListTasksResponse
 */
export const ListTasksResponse = /*@__PURE__*/ proto3.makeMessageType(
  "api.tasks.v1.ListTasksResponse",
  () => [
    { no: 1, name: "tasks", kind: "message", T: Task, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

