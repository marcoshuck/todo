//
//  Generated code. Do not modify.
//  source: api/tasks/v1/tasks.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'tasks.pb.dart' as $2;
import 'tasks.pbjson.dart';

export 'tasks.pb.dart';

abstract class TasksWriterServiceBase extends $pb.GeneratedService {
  $async.Future<$2.Task> createTask($pb.ServerContext ctx, $2.CreateTaskRequest request);
  $async.Future<$2.Task> deleteTask($pb.ServerContext ctx, $2.DeleteTaskRequest request);
  $async.Future<$2.Task> undeleteTask($pb.ServerContext ctx, $2.UndeleteTaskRequest request);
  $async.Future<$2.Task> updateTask($pb.ServerContext ctx, $2.UpdateTaskRequest request);

  $pb.GeneratedMessage createRequest($core.String methodName) {
    switch (methodName) {
      case 'CreateTask': return $2.CreateTaskRequest();
      case 'DeleteTask': return $2.DeleteTaskRequest();
      case 'UndeleteTask': return $2.UndeleteTaskRequest();
      case 'UpdateTask': return $2.UpdateTaskRequest();
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String methodName, $pb.GeneratedMessage request) {
    switch (methodName) {
      case 'CreateTask': return this.createTask(ctx, request as $2.CreateTaskRequest);
      case 'DeleteTask': return this.deleteTask(ctx, request as $2.DeleteTaskRequest);
      case 'UndeleteTask': return this.undeleteTask(ctx, request as $2.UndeleteTaskRequest);
      case 'UpdateTask': return this.updateTask(ctx, request as $2.UpdateTaskRequest);
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => TasksWriterServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => TasksWriterServiceBase$messageJson;
}

abstract class TasksReaderServiceBase extends $pb.GeneratedService {
  $async.Future<$2.Task> getTask($pb.ServerContext ctx, $2.GetTaskRequest request);
  $async.Future<$2.ListTasksResponse> listTasks($pb.ServerContext ctx, $2.ListTasksRequest request);

  $pb.GeneratedMessage createRequest($core.String methodName) {
    switch (methodName) {
      case 'GetTask': return $2.GetTaskRequest();
      case 'ListTasks': return $2.ListTasksRequest();
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String methodName, $pb.GeneratedMessage request) {
    switch (methodName) {
      case 'GetTask': return this.getTask(ctx, request as $2.GetTaskRequest);
      case 'ListTasks': return this.listTasks(ctx, request as $2.ListTasksRequest);
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => TasksReaderServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => TasksReaderServiceBase$messageJson;
}

