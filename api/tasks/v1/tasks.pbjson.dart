//
//  Generated code. Do not modify.
//  source: api/tasks/v1/tasks.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

import '../../../google/protobuf/field_mask.pbjson.dart' as $1;
import '../../../google/protobuf/timestamp.pbjson.dart' as $0;

@$core.Deprecated('Use taskDescriptor instead')
const Task$json = {
  '1': 'Task',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'title', '3': 2, '4': 1, '5': 9, '8': {}, '10': 'title'},
    {'1': 'description', '3': 3, '4': 1, '5': 9, '10': 'description'},
    {'1': 'deadline', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'deadline'},
    {'1': 'completed_at', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'completedAt'},
    {'1': 'create_time', '3': 1000, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createTime'},
    {'1': 'update_time', '3': 1001, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updateTime'},
  ],
  '7': {},
};

/// Descriptor for `Task`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List taskDescriptor = $convert.base64Decode(
    'CgRUYXNrEg4KAmlkGAEgASgDUgJpZBIdCgV0aXRsZRgCIAEoCUIHukgEcgIQA1IFdGl0bGUSIA'
    'oLZGVzY3JpcHRpb24YAyABKAlSC2Rlc2NyaXB0aW9uEjYKCGRlYWRsaW5lGAQgASgLMhouZ29v'
    'Z2xlLnByb3RvYnVmLlRpbWVzdGFtcFIIZGVhZGxpbmUSPQoMY29tcGxldGVkX2F0GAUgASgLMh'
    'ouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFILY29tcGxldGVkQXQSPAoLY3JlYXRlX3RpbWUY'
    '6AcgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIKY3JlYXRlVGltZRI8Cgt1cGRhdG'
    'VfdGltZRjpByABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgp1cGRhdGVUaW1lOjXq'
    'QTIKFXRvZG8uaHVjay5jb20uYXIvVGFzaxIMdGFza3Mve3Rhc2t9KgV0YXNrczIEdGFzaw==');

@$core.Deprecated('Use createTaskRequestDescriptor instead')
const CreateTaskRequest$json = {
  '1': 'CreateTaskRequest',
  '2': [
    {'1': 'task', '3': 2, '4': 1, '5': 11, '6': '.api.tasks.v1.Task', '8': {}, '10': 'task'},
  ],
  '9': [
    {'1': 1, '2': 2},
  ],
  '10': ['parent'],
};

/// Descriptor for `CreateTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTaskRequestDescriptor = $convert.base64Decode(
    'ChFDcmVhdGVUYXNrUmVxdWVzdBIrCgR0YXNrGAIgASgLMhIuYXBpLnRhc2tzLnYxLlRhc2tCA+'
    'BBAlIEdGFza0oECAEQAlIGcGFyZW50');

@$core.Deprecated('Use deleteTaskRequestDescriptor instead')
const DeleteTaskRequest$json = {
  '1': 'DeleteTaskRequest',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '8': {}, '10': 'id'},
  ],
};

/// Descriptor for `DeleteTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteTaskRequestDescriptor = $convert.base64Decode(
    'ChFEZWxldGVUYXNrUmVxdWVzdBI0CgJpZBgBIAEoA0Ik4EEC+kEXChV0b2RvLmh1Y2suY29tLm'
    'FyL1Rhc2u6SAQiAiAAUgJpZA==');

@$core.Deprecated('Use undeleteTaskRequestDescriptor instead')
const UndeleteTaskRequest$json = {
  '1': 'UndeleteTaskRequest',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '8': {}, '10': 'id'},
  ],
};

/// Descriptor for `UndeleteTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List undeleteTaskRequestDescriptor = $convert.base64Decode(
    'ChNVbmRlbGV0ZVRhc2tSZXF1ZXN0EjQKAmlkGAEgASgDQiTgQQL6QRcKFXRvZG8uaHVjay5jb2'
    '0uYXIvVGFza7pIBCICIABSAmlk');

@$core.Deprecated('Use updateTaskRequestDescriptor instead')
const UpdateTaskRequest$json = {
  '1': 'UpdateTaskRequest',
  '2': [
    {'1': 'task', '3': 1, '4': 1, '5': 11, '6': '.api.tasks.v1.Task', '8': {}, '10': 'task'},
    {'1': 'update_mask', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.FieldMask', '10': 'updateMask'},
  ],
};

/// Descriptor for `UpdateTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateTaskRequestDescriptor = $convert.base64Decode(
    'ChFVcGRhdGVUYXNrUmVxdWVzdBIxCgR0YXNrGAEgASgLMhIuYXBpLnRhc2tzLnYxLlRhc2tCCe'
    'BBArpIA8ABAVIEdGFzaxI7Cgt1cGRhdGVfbWFzaxgCIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5G'
    'aWVsZE1hc2tSCnVwZGF0ZU1hc2s=');

@$core.Deprecated('Use getTaskRequestDescriptor instead')
const GetTaskRequest$json = {
  '1': 'GetTaskRequest',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '8': {}, '10': 'id'},
  ],
};

/// Descriptor for `GetTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTaskRequestDescriptor = $convert.base64Decode(
    'Cg5HZXRUYXNrUmVxdWVzdBI0CgJpZBgBIAEoA0Ik4EEC+kEXChV0b2RvLmh1Y2suY29tLmFyL1'
    'Rhc2u6SAQiAiAAUgJpZA==');

@$core.Deprecated('Use listTasksRequestDescriptor instead')
const ListTasksRequest$json = {
  '1': 'ListTasksRequest',
  '2': [
    {'1': 'page_size', '3': 2, '4': 1, '5': 5, '8': {}, '10': 'pageSize'},
    {'1': 'page_token', '3': 3, '4': 1, '5': 9, '8': {}, '10': 'pageToken'},
  ],
  '9': [
    {'1': 1, '2': 2},
  ],
  '10': ['parent'],
};

/// Descriptor for `ListTasksRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listTasksRequestDescriptor = $convert.base64Decode(
    'ChBMaXN0VGFza3NSZXF1ZXN0EicKCXBhZ2Vfc2l6ZRgCIAEoBUIK4EEBukgEGgIgAFIIcGFnZV'
    'NpemUSIgoKcGFnZV90b2tlbhgDIAEoCUID4EEBUglwYWdlVG9rZW5KBAgBEAJSBnBhcmVudA==');

@$core.Deprecated('Use listTasksResponseDescriptor instead')
const ListTasksResponse$json = {
  '1': 'ListTasksResponse',
  '2': [
    {'1': 'tasks', '3': 1, '4': 3, '5': 11, '6': '.api.tasks.v1.Task', '10': 'tasks'},
    {'1': 'next_page_token', '3': 2, '4': 1, '5': 9, '10': 'nextPageToken'},
  ],
};

/// Descriptor for `ListTasksResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listTasksResponseDescriptor = $convert.base64Decode(
    'ChFMaXN0VGFza3NSZXNwb25zZRIoCgV0YXNrcxgBIAMoCzISLmFwaS50YXNrcy52MS5UYXNrUg'
    'V0YXNrcxImCg9uZXh0X3BhZ2VfdG9rZW4YAiABKAlSDW5leHRQYWdlVG9rZW4=');

const $core.Map<$core.String, $core.dynamic> TasksWriterServiceBase$json = {
  '1': 'TasksWriterService',
  '2': [
    {'1': 'CreateTask', '2': '.api.tasks.v1.CreateTaskRequest', '3': '.api.tasks.v1.Task', '4': {}},
    {'1': 'DeleteTask', '2': '.api.tasks.v1.DeleteTaskRequest', '3': '.api.tasks.v1.Task', '4': {}},
    {'1': 'UndeleteTask', '2': '.api.tasks.v1.UndeleteTaskRequest', '3': '.api.tasks.v1.Task', '4': {}},
    {'1': 'UpdateTask', '2': '.api.tasks.v1.UpdateTaskRequest', '3': '.api.tasks.v1.Task', '4': {}},
  ],
};

@$core.Deprecated('Use tasksWriterServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> TasksWriterServiceBase$messageJson = {
  '.api.tasks.v1.CreateTaskRequest': CreateTaskRequest$json,
  '.api.tasks.v1.Task': Task$json,
  '.google.protobuf.Timestamp': $0.Timestamp$json,
  '.api.tasks.v1.DeleteTaskRequest': DeleteTaskRequest$json,
  '.api.tasks.v1.UndeleteTaskRequest': UndeleteTaskRequest$json,
  '.api.tasks.v1.UpdateTaskRequest': UpdateTaskRequest$json,
  '.google.protobuf.FieldMask': $1.FieldMask$json,
};

/// Descriptor for `TasksWriterService`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List tasksWriterServiceDescriptor = $convert.base64Decode(
    'ChJUYXNrc1dyaXRlclNlcnZpY2USYQoKQ3JlYXRlVGFzaxIfLmFwaS50YXNrcy52MS5DcmVhdG'
    'VUYXNrUmVxdWVzdBoSLmFwaS50YXNrcy52MS5UYXNrIh7aQQR0YXNrgtPkkwIROgR0YXNrIgkv'
    'djEvdGFza3MSXgoKRGVsZXRlVGFzaxIfLmFwaS50YXNrcy52MS5EZWxldGVUYXNrUmVxdWVzdB'
    'oSLmFwaS50YXNrcy52MS5UYXNrIhvaQQJpZILT5JMCECoOL3YxL3Rhc2tzL3tpZH0SaQoMVW5k'
    'ZWxldGVUYXNrEiEuYXBpLnRhc2tzLnYxLlVuZGVsZXRlVGFza1JlcXVlc3QaEi5hcGkudGFza3'
    'MudjEuVGFzayIi2kECaWSC0+STAhc6ASoiEi92MS90YXNrczp1bmRlbGV0ZRJ1CgpVcGRhdGVU'
    'YXNrEh8uYXBpLnRhc2tzLnYxLlVwZGF0ZVRhc2tSZXF1ZXN0GhIuYXBpLnRhc2tzLnYxLlRhc2'
    'siMtpBDmlkLHVwZGF0ZV9tYXNrgtPkkwIbOgR0YXNrMhMvdjEvdGFza3Mve3Rhc2suaWR9');

const $core.Map<$core.String, $core.dynamic> TasksReaderServiceBase$json = {
  '1': 'TasksReaderService',
  '2': [
    {'1': 'GetTask', '2': '.api.tasks.v1.GetTaskRequest', '3': '.api.tasks.v1.Task', '4': {}},
    {'1': 'ListTasks', '2': '.api.tasks.v1.ListTasksRequest', '3': '.api.tasks.v1.ListTasksResponse', '4': {}},
  ],
};

@$core.Deprecated('Use tasksReaderServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> TasksReaderServiceBase$messageJson = {
  '.api.tasks.v1.GetTaskRequest': GetTaskRequest$json,
  '.api.tasks.v1.Task': Task$json,
  '.google.protobuf.Timestamp': $0.Timestamp$json,
  '.api.tasks.v1.ListTasksRequest': ListTasksRequest$json,
  '.api.tasks.v1.ListTasksResponse': ListTasksResponse$json,
};

/// Descriptor for `TasksReaderService`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List tasksReaderServiceDescriptor = $convert.base64Decode(
    'ChJUYXNrc1JlYWRlclNlcnZpY2USWgoHR2V0VGFzaxIcLmFwaS50YXNrcy52MS5HZXRUYXNrUm'
    'VxdWVzdBoSLmFwaS50YXNrcy52MS5UYXNrIh3aQQR0YXNrgtPkkwIQEg4vdjEvdGFza3Mve2lk'
    'fRJfCglMaXN0VGFza3MSHi5hcGkudGFza3MudjEuTGlzdFRhc2tzUmVxdWVzdBofLmFwaS50YX'
    'Nrcy52MS5MaXN0VGFza3NSZXNwb25zZSIRgtPkkwILEgkvdjEvdGFza3M=');

