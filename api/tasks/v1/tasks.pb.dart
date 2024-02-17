//
//  Generated code. Do not modify.
//  source: api/tasks/v1/tasks.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../../../google/protobuf/field_mask.pb.dart' as $1;
import '../../../google/protobuf/timestamp.pb.dart' as $0;

class Task extends $pb.GeneratedMessage {
  factory Task({
    $fixnum.Int64? id,
    $core.String? title,
    $core.String? description,
    $0.Timestamp? deadline,
    $0.Timestamp? completedAt,
    $0.Timestamp? createTime,
    $0.Timestamp? updateTime,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (title != null) {
      $result.title = title;
    }
    if (description != null) {
      $result.description = description;
    }
    if (deadline != null) {
      $result.deadline = deadline;
    }
    if (completedAt != null) {
      $result.completedAt = completedAt;
    }
    if (createTime != null) {
      $result.createTime = createTime;
    }
    if (updateTime != null) {
      $result.updateTime = updateTime;
    }
    return $result;
  }
  Task._() : super();
  factory Task.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Task.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Task', package: const $pb.PackageName(_omitMessageNames ? '' : 'api.tasks.v1'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'title')
    ..aOS(3, _omitFieldNames ? '' : 'description')
    ..aOM<$0.Timestamp>(4, _omitFieldNames ? '' : 'deadline', subBuilder: $0.Timestamp.create)
    ..aOM<$0.Timestamp>(5, _omitFieldNames ? '' : 'completedAt', subBuilder: $0.Timestamp.create)
    ..aOM<$0.Timestamp>(1000, _omitFieldNames ? '' : 'createTime', subBuilder: $0.Timestamp.create)
    ..aOM<$0.Timestamp>(1001, _omitFieldNames ? '' : 'updateTime', subBuilder: $0.Timestamp.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Task clone() => Task()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Task copyWith(void Function(Task) updates) => super.copyWith((message) => updates(message as Task)) as Task;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Task create() => Task._();
  Task createEmptyInstance() => create();
  static $pb.PbList<Task> createRepeated() => $pb.PbList<Task>();
  @$core.pragma('dart2js:noInline')
  static Task getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Task>(create);
  static Task? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get description => $_getSZ(2);
  @$pb.TagNumber(3)
  set description($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDescription() => $_has(2);
  @$pb.TagNumber(3)
  void clearDescription() => clearField(3);

  @$pb.TagNumber(4)
  $0.Timestamp get deadline => $_getN(3);
  @$pb.TagNumber(4)
  set deadline($0.Timestamp v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasDeadline() => $_has(3);
  @$pb.TagNumber(4)
  void clearDeadline() => clearField(4);
  @$pb.TagNumber(4)
  $0.Timestamp ensureDeadline() => $_ensure(3);

  @$pb.TagNumber(5)
  $0.Timestamp get completedAt => $_getN(4);
  @$pb.TagNumber(5)
  set completedAt($0.Timestamp v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasCompletedAt() => $_has(4);
  @$pb.TagNumber(5)
  void clearCompletedAt() => clearField(5);
  @$pb.TagNumber(5)
  $0.Timestamp ensureCompletedAt() => $_ensure(4);

  @$pb.TagNumber(1000)
  $0.Timestamp get createTime => $_getN(5);
  @$pb.TagNumber(1000)
  set createTime($0.Timestamp v) { setField(1000, v); }
  @$pb.TagNumber(1000)
  $core.bool hasCreateTime() => $_has(5);
  @$pb.TagNumber(1000)
  void clearCreateTime() => clearField(1000);
  @$pb.TagNumber(1000)
  $0.Timestamp ensureCreateTime() => $_ensure(5);

  @$pb.TagNumber(1001)
  $0.Timestamp get updateTime => $_getN(6);
  @$pb.TagNumber(1001)
  set updateTime($0.Timestamp v) { setField(1001, v); }
  @$pb.TagNumber(1001)
  $core.bool hasUpdateTime() => $_has(6);
  @$pb.TagNumber(1001)
  void clearUpdateTime() => clearField(1001);
  @$pb.TagNumber(1001)
  $0.Timestamp ensureUpdateTime() => $_ensure(6);
}

class CreateTaskRequest extends $pb.GeneratedMessage {
  factory CreateTaskRequest({
    Task? task,
  }) {
    final $result = create();
    if (task != null) {
      $result.task = task;
    }
    return $result;
  }
  CreateTaskRequest._() : super();
  factory CreateTaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateTaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateTaskRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'api.tasks.v1'), createEmptyInstance: create)
    ..aOM<Task>(2, _omitFieldNames ? '' : 'task', subBuilder: Task.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateTaskRequest clone() => CreateTaskRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateTaskRequest copyWith(void Function(CreateTaskRequest) updates) => super.copyWith((message) => updates(message as CreateTaskRequest)) as CreateTaskRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTaskRequest create() => CreateTaskRequest._();
  CreateTaskRequest createEmptyInstance() => create();
  static $pb.PbList<CreateTaskRequest> createRepeated() => $pb.PbList<CreateTaskRequest>();
  @$core.pragma('dart2js:noInline')
  static CreateTaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTaskRequest>(create);
  static CreateTaskRequest? _defaultInstance;

  /// Task is the the task to create.
  @$pb.TagNumber(2)
  Task get task => $_getN(0);
  @$pb.TagNumber(2)
  set task(Task v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasTask() => $_has(0);
  @$pb.TagNumber(2)
  void clearTask() => clearField(2);
  @$pb.TagNumber(2)
  Task ensureTask() => $_ensure(0);
}

class DeleteTaskRequest extends $pb.GeneratedMessage {
  factory DeleteTaskRequest({
    $fixnum.Int64? id,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    return $result;
  }
  DeleteTaskRequest._() : super();
  factory DeleteTaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteTaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeleteTaskRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'api.tasks.v1'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteTaskRequest clone() => DeleteTaskRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteTaskRequest copyWith(void Function(DeleteTaskRequest) updates) => super.copyWith((message) => updates(message as DeleteTaskRequest)) as DeleteTaskRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeleteTaskRequest create() => DeleteTaskRequest._();
  DeleteTaskRequest createEmptyInstance() => create();
  static $pb.PbList<DeleteTaskRequest> createRepeated() => $pb.PbList<DeleteTaskRequest>();
  @$core.pragma('dart2js:noInline')
  static DeleteTaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteTaskRequest>(create);
  static DeleteTaskRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class UndeleteTaskRequest extends $pb.GeneratedMessage {
  factory UndeleteTaskRequest({
    $fixnum.Int64? id,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    return $result;
  }
  UndeleteTaskRequest._() : super();
  factory UndeleteTaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UndeleteTaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'UndeleteTaskRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'api.tasks.v1'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UndeleteTaskRequest clone() => UndeleteTaskRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UndeleteTaskRequest copyWith(void Function(UndeleteTaskRequest) updates) => super.copyWith((message) => updates(message as UndeleteTaskRequest)) as UndeleteTaskRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UndeleteTaskRequest create() => UndeleteTaskRequest._();
  UndeleteTaskRequest createEmptyInstance() => create();
  static $pb.PbList<UndeleteTaskRequest> createRepeated() => $pb.PbList<UndeleteTaskRequest>();
  @$core.pragma('dart2js:noInline')
  static UndeleteTaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UndeleteTaskRequest>(create);
  static UndeleteTaskRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class UpdateTaskRequest extends $pb.GeneratedMessage {
  factory UpdateTaskRequest({
    Task? task,
    $1.FieldMask? updateMask,
  }) {
    final $result = create();
    if (task != null) {
      $result.task = task;
    }
    if (updateMask != null) {
      $result.updateMask = updateMask;
    }
    return $result;
  }
  UpdateTaskRequest._() : super();
  factory UpdateTaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateTaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'UpdateTaskRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'api.tasks.v1'), createEmptyInstance: create)
    ..aOM<Task>(1, _omitFieldNames ? '' : 'task', subBuilder: Task.create)
    ..aOM<$1.FieldMask>(2, _omitFieldNames ? '' : 'updateMask', subBuilder: $1.FieldMask.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateTaskRequest clone() => UpdateTaskRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateTaskRequest copyWith(void Function(UpdateTaskRequest) updates) => super.copyWith((message) => updates(message as UpdateTaskRequest)) as UpdateTaskRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UpdateTaskRequest create() => UpdateTaskRequest._();
  UpdateTaskRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateTaskRequest> createRepeated() => $pb.PbList<UpdateTaskRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdateTaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateTaskRequest>(create);
  static UpdateTaskRequest? _defaultInstance;

  @$pb.TagNumber(1)
  Task get task => $_getN(0);
  @$pb.TagNumber(1)
  set task(Task v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTask() => $_has(0);
  @$pb.TagNumber(1)
  void clearTask() => clearField(1);
  @$pb.TagNumber(1)
  Task ensureTask() => $_ensure(0);

  @$pb.TagNumber(2)
  $1.FieldMask get updateMask => $_getN(1);
  @$pb.TagNumber(2)
  set updateMask($1.FieldMask v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasUpdateMask() => $_has(1);
  @$pb.TagNumber(2)
  void clearUpdateMask() => clearField(2);
  @$pb.TagNumber(2)
  $1.FieldMask ensureUpdateMask() => $_ensure(1);
}

class GetTaskRequest extends $pb.GeneratedMessage {
  factory GetTaskRequest({
    $fixnum.Int64? id,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    return $result;
  }
  GetTaskRequest._() : super();
  factory GetTaskRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetTaskRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetTaskRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'api.tasks.v1'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'id')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetTaskRequest clone() => GetTaskRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetTaskRequest copyWith(void Function(GetTaskRequest) updates) => super.copyWith((message) => updates(message as GetTaskRequest)) as GetTaskRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetTaskRequest create() => GetTaskRequest._();
  GetTaskRequest createEmptyInstance() => create();
  static $pb.PbList<GetTaskRequest> createRepeated() => $pb.PbList<GetTaskRequest>();
  @$core.pragma('dart2js:noInline')
  static GetTaskRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTaskRequest>(create);
  static GetTaskRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class ListTasksRequest extends $pb.GeneratedMessage {
  factory ListTasksRequest({
    $core.int? pageSize,
    $core.String? pageToken,
  }) {
    final $result = create();
    if (pageSize != null) {
      $result.pageSize = pageSize;
    }
    if (pageToken != null) {
      $result.pageToken = pageToken;
    }
    return $result;
  }
  ListTasksRequest._() : super();
  factory ListTasksRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListTasksRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ListTasksRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'api.tasks.v1'), createEmptyInstance: create)
    ..a<$core.int>(2, _omitFieldNames ? '' : 'pageSize', $pb.PbFieldType.O3)
    ..aOS(3, _omitFieldNames ? '' : 'pageToken')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListTasksRequest clone() => ListTasksRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListTasksRequest copyWith(void Function(ListTasksRequest) updates) => super.copyWith((message) => updates(message as ListTasksRequest)) as ListTasksRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListTasksRequest create() => ListTasksRequest._();
  ListTasksRequest createEmptyInstance() => create();
  static $pb.PbList<ListTasksRequest> createRepeated() => $pb.PbList<ListTasksRequest>();
  @$core.pragma('dart2js:noInline')
  static ListTasksRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListTasksRequest>(create);
  static ListTasksRequest? _defaultInstance;

  @$pb.TagNumber(2)
  $core.int get pageSize => $_getIZ(0);
  @$pb.TagNumber(2)
  set pageSize($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(2)
  $core.bool hasPageSize() => $_has(0);
  @$pb.TagNumber(2)
  void clearPageSize() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get pageToken => $_getSZ(1);
  @$pb.TagNumber(3)
  set pageToken($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasPageToken() => $_has(1);
  @$pb.TagNumber(3)
  void clearPageToken() => clearField(3);
}

class ListTasksResponse extends $pb.GeneratedMessage {
  factory ListTasksResponse({
    $core.Iterable<Task>? tasks,
    $core.String? nextPageToken,
  }) {
    final $result = create();
    if (tasks != null) {
      $result.tasks.addAll(tasks);
    }
    if (nextPageToken != null) {
      $result.nextPageToken = nextPageToken;
    }
    return $result;
  }
  ListTasksResponse._() : super();
  factory ListTasksResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListTasksResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ListTasksResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'api.tasks.v1'), createEmptyInstance: create)
    ..pc<Task>(1, _omitFieldNames ? '' : 'tasks', $pb.PbFieldType.PM, subBuilder: Task.create)
    ..aOS(2, _omitFieldNames ? '' : 'nextPageToken')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListTasksResponse clone() => ListTasksResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListTasksResponse copyWith(void Function(ListTasksResponse) updates) => super.copyWith((message) => updates(message as ListTasksResponse)) as ListTasksResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListTasksResponse create() => ListTasksResponse._();
  ListTasksResponse createEmptyInstance() => create();
  static $pb.PbList<ListTasksResponse> createRepeated() => $pb.PbList<ListTasksResponse>();
  @$core.pragma('dart2js:noInline')
  static ListTasksResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListTasksResponse>(create);
  static ListTasksResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Task> get tasks => $_getList(0);

  @$pb.TagNumber(2)
  $core.String get nextPageToken => $_getSZ(1);
  @$pb.TagNumber(2)
  set nextPageToken($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNextPageToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearNextPageToken() => clearField(2);
}

class TasksWriterServiceApi {
  $pb.RpcClient _client;
  TasksWriterServiceApi(this._client);

  $async.Future<Task> createTask($pb.ClientContext? ctx, CreateTaskRequest request) =>
    _client.invoke<Task>(ctx, 'TasksWriterService', 'CreateTask', request, Task())
  ;
  $async.Future<Task> deleteTask($pb.ClientContext? ctx, DeleteTaskRequest request) =>
    _client.invoke<Task>(ctx, 'TasksWriterService', 'DeleteTask', request, Task())
  ;
  $async.Future<Task> undeleteTask($pb.ClientContext? ctx, UndeleteTaskRequest request) =>
    _client.invoke<Task>(ctx, 'TasksWriterService', 'UndeleteTask', request, Task())
  ;
  $async.Future<Task> updateTask($pb.ClientContext? ctx, UpdateTaskRequest request) =>
    _client.invoke<Task>(ctx, 'TasksWriterService', 'UpdateTask', request, Task())
  ;
}

class TasksReaderServiceApi {
  $pb.RpcClient _client;
  TasksReaderServiceApi(this._client);

  $async.Future<Task> getTask($pb.ClientContext? ctx, GetTaskRequest request) =>
    _client.invoke<Task>(ctx, 'TasksReaderService', 'GetTask', request, Task())
  ;
  $async.Future<ListTasksResponse> listTasks($pb.ClientContext? ctx, ListTasksRequest request) =>
    _client.invoke<ListTasksResponse>(ctx, 'TasksReaderService', 'ListTasks', request, ListTasksResponse())
  ;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
