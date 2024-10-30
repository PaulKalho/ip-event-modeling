//
//  Generated code. Do not modify.
//  source: services/tasks_svc/v1/task_svc.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use createTaskRequestDescriptor instead')
const CreateTaskRequest$json = {
  '1': 'CreateTaskRequest',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    {'1': 'description', '3': 2, '4': 1, '5': 9, '9': 0, '10': 'description', '17': true},
    {'1': 'public', '3': 4, '4': 1, '5': 8, '9': 1, '10': 'public', '17': true},
    {'1': 'due_at', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'dueAt'},
    {'1': 'initial_status', '3': 6, '4': 1, '5': 14, '6': '.services.tasks_svc.v1.TaskStatus', '9': 2, '10': 'initialStatus', '17': true},
    {'1': 'subtasks', '3': 8, '4': 3, '5': 11, '6': '.services.tasks_svc.v1.CreateTaskRequest.SubTask', '10': 'subtasks'},
  ],
  '3': [CreateTaskRequest_SubTask$json],
  '8': [
    {'1': '_description'},
    {'1': '_public'},
    {'1': '_initial_status'},
  ],
};

@$core.Deprecated('Use createTaskRequestDescriptor instead')
const CreateTaskRequest_SubTask$json = {
  '1': 'SubTask',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    {'1': 'done', '3': 2, '4': 1, '5': 8, '9': 0, '10': 'done', '17': true},
  ],
  '8': [
    {'1': '_done'},
  ],
};

/// Descriptor for `CreateTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTaskRequestDescriptor = $convert.base64Decode(
    'ChFDcmVhdGVUYXNrUmVxdWVzdBISCgRuYW1lGAEgASgJUgRuYW1lEiUKC2Rlc2NyaXB0aW9uGA'
    'IgASgJSABSC2Rlc2NyaXB0aW9uiAEBEhsKBnB1YmxpYxgEIAEoCEgBUgZwdWJsaWOIAQESMQoG'
    'ZHVlX2F0GAUgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIFZHVlQXQSTQoOaW5pdG'
    'lhbF9zdGF0dXMYBiABKA4yIS5zZXJ2aWNlcy50YXNrc19zdmMudjEuVGFza1N0YXR1c0gCUg1p'
    'bml0aWFsU3RhdHVziAEBEkwKCHN1YnRhc2tzGAggAygLMjAuc2VydmljZXMudGFza3Nfc3ZjLn'
    'YxLkNyZWF0ZVRhc2tSZXF1ZXN0LlN1YlRhc2tSCHN1YnRhc2tzGj8KB1N1YlRhc2sSEgoEbmFt'
    'ZRgBIAEoCVIEbmFtZRIXCgRkb25lGAIgASgISABSBGRvbmWIAQFCBwoFX2RvbmVCDgoMX2Rlc2'
    'NyaXB0aW9uQgkKB19wdWJsaWNCEQoPX2luaXRpYWxfc3RhdHVz');

@$core.Deprecated('Use createTaskResponseDescriptor instead')
const CreateTaskResponse$json = {
  '1': 'CreateTaskResponse',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `CreateTaskResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTaskResponseDescriptor = $convert.base64Decode(
    'ChJDcmVhdGVUYXNrUmVzcG9uc2USDgoCaWQYASABKAlSAmlk');

@$core.Deprecated('Use updateTaskRequestDescriptor instead')
const UpdateTaskRequest$json = {
  '1': 'UpdateTaskRequest',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '9': 0, '10': 'name', '17': true},
    {'1': 'description', '3': 3, '4': 1, '5': 9, '9': 1, '10': 'description', '17': true},
    {'1': 'due_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '9': 2, '10': 'dueAt', '17': true},
    {'1': 'status', '3': 5, '4': 1, '5': 14, '6': '.services.tasks_svc.v1.TaskStatus', '9': 3, '10': 'status', '17': true},
    {'1': 'public', '3': 6, '4': 1, '5': 8, '9': 4, '10': 'public', '17': true},
  ],
  '8': [
    {'1': '_name'},
    {'1': '_description'},
    {'1': '_due_at'},
    {'1': '_status'},
    {'1': '_public'},
  ],
};

/// Descriptor for `UpdateTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateTaskRequestDescriptor = $convert.base64Decode(
    'ChFVcGRhdGVUYXNrUmVxdWVzdBIOCgJpZBgBIAEoCVICaWQSFwoEbmFtZRgCIAEoCUgAUgRuYW'
    '1liAEBEiUKC2Rlc2NyaXB0aW9uGAMgASgJSAFSC2Rlc2NyaXB0aW9uiAEBEjYKBmR1ZV9hdBgE'
    'IAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBIAlIFZHVlQXSIAQESPgoGc3RhdHVzGA'
    'UgASgOMiEuc2VydmljZXMudGFza3Nfc3ZjLnYxLlRhc2tTdGF0dXNIA1IGc3RhdHVziAEBEhsK'
    'BnB1YmxpYxgGIAEoCEgEUgZwdWJsaWOIAQFCBwoFX25hbWVCDgoMX2Rlc2NyaXB0aW9uQgkKB1'
    '9kdWVfYXRCCQoHX3N0YXR1c0IJCgdfcHVibGlj');

@$core.Deprecated('Use updateTaskResponseDescriptor instead')
const UpdateTaskResponse$json = {
  '1': 'UpdateTaskResponse',
};

/// Descriptor for `UpdateTaskResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateTaskResponseDescriptor = $convert.base64Decode(
    'ChJVcGRhdGVUYXNrUmVzcG9uc2U=');

@$core.Deprecated('Use getTaskRequestDescriptor instead')
const GetTaskRequest$json = {
  '1': 'GetTaskRequest',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `GetTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTaskRequestDescriptor = $convert.base64Decode(
    'Cg5HZXRUYXNrUmVxdWVzdBIOCgJpZBgBIAEoCVICaWQ=');

@$core.Deprecated('Use getTaskResponseDescriptor instead')
const GetTaskResponse$json = {
  '1': 'GetTaskResponse',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    {'1': 'description', '3': 3, '4': 1, '5': 9, '10': 'description'},
    {'1': 'subtasks', '3': 4, '4': 3, '5': 11, '6': '.services.tasks_svc.v1.GetTaskResponse.SubTask', '10': 'subtasks'},
    {'1': 'status', '3': 5, '4': 1, '5': 14, '6': '.services.tasks_svc.v1.TaskStatus', '10': 'status'},
    {'1': 'created_at', '3': 6, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    {'1': 'public', '3': 7, '4': 1, '5': 8, '10': 'public'},
    {'1': 'due_at', '3': 8, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '9': 0, '10': 'dueAt', '17': true},
    {'1': 'created_by', '3': 9, '4': 1, '5': 9, '10': 'createdBy'},
  ],
  '3': [GetTaskResponse_SubTask$json],
  '8': [
    {'1': '_due_at'},
  ],
};

@$core.Deprecated('Use getTaskResponseDescriptor instead')
const GetTaskResponse_SubTask$json = {
  '1': 'SubTask',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    {'1': 'done', '3': 3, '4': 1, '5': 8, '10': 'done'},
  ],
};

/// Descriptor for `GetTaskResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTaskResponseDescriptor = $convert.base64Decode(
    'Cg9HZXRUYXNrUmVzcG9uc2USDgoCaWQYASABKAlSAmlkEhIKBG5hbWUYAiABKAlSBG5hbWUSIA'
    'oLZGVzY3JpcHRpb24YAyABKAlSC2Rlc2NyaXB0aW9uEkoKCHN1YnRhc2tzGAQgAygLMi4uc2Vy'
    'dmljZXMudGFza3Nfc3ZjLnYxLkdldFRhc2tSZXNwb25zZS5TdWJUYXNrUghzdWJ0YXNrcxI5Cg'
    'ZzdGF0dXMYBSABKA4yIS5zZXJ2aWNlcy50YXNrc19zdmMudjEuVGFza1N0YXR1c1IGc3RhdHVz'
    'EjkKCmNyZWF0ZWRfYXQYBiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgljcmVhdG'
    'VkQXQSFgoGcHVibGljGAcgASgIUgZwdWJsaWMSNgoGZHVlX2F0GAggASgLMhouZ29vZ2xlLnBy'
    'b3RvYnVmLlRpbWVzdGFtcEgAUgVkdWVBdIgBARIdCgpjcmVhdGVkX2J5GAkgASgJUgljcmVhdG'
    'VkQnkaQQoHU3ViVGFzaxIOCgJpZBgBIAEoCVICaWQSEgoEbmFtZRgCIAEoCVIEbmFtZRISCgRk'
    'b25lGAMgASgIUgRkb25lQgkKB19kdWVfYXQ=');

@$core.Deprecated('Use getAllTasksRequestDescriptor instead')
const GetAllTasksRequest$json = {
  '1': 'GetAllTasksRequest',
};

/// Descriptor for `GetAllTasksRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getAllTasksRequestDescriptor = $convert.base64Decode(
    'ChJHZXRBbGxUYXNrc1JlcXVlc3Q=');

@$core.Deprecated('Use getAllTasksResponseDescriptor instead')
const GetAllTasksResponse$json = {
  '1': 'GetAllTasksResponse',
  '2': [
    {'1': 'tasks', '3': 1, '4': 3, '5': 11, '6': '.services.tasks_svc.v1.GetAllTasksResponse.Task', '10': 'tasks'},
  ],
  '3': [GetAllTasksResponse_Task$json],
};

@$core.Deprecated('Use getAllTasksResponseDescriptor instead')
const GetAllTasksResponse_Task$json = {
  '1': 'Task',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    {'1': 'description', '3': 3, '4': 1, '5': 9, '10': 'description'},
    {'1': 'subtasks', '3': 4, '4': 3, '5': 11, '6': '.services.tasks_svc.v1.GetAllTasksResponse.Task.SubTask', '10': 'subtasks'},
    {'1': 'status', '3': 5, '4': 1, '5': 14, '6': '.services.tasks_svc.v1.TaskStatus', '10': 'status'},
    {'1': 'created_at', '3': 6, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    {'1': 'public', '3': 7, '4': 1, '5': 8, '10': 'public'},
    {'1': 'due_at', '3': 8, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '9': 0, '10': 'dueAt', '17': true},
    {'1': 'created_by', '3': 9, '4': 1, '5': 9, '10': 'createdBy'},
  ],
  '3': [GetAllTasksResponse_Task_SubTask$json],
  '8': [
    {'1': '_due_at'},
  ],
};

@$core.Deprecated('Use getAllTasksResponseDescriptor instead')
const GetAllTasksResponse_Task_SubTask$json = {
  '1': 'SubTask',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    {'1': 'done', '3': 3, '4': 1, '5': 8, '10': 'done'},
  ],
};

/// Descriptor for `GetAllTasksResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getAllTasksResponseDescriptor = $convert.base64Decode(
    'ChNHZXRBbGxUYXNrc1Jlc3BvbnNlEkUKBXRhc2tzGAEgAygLMi8uc2VydmljZXMudGFza3Nfc3'
    'ZjLnYxLkdldEFsbFRhc2tzUmVzcG9uc2UuVGFza1IFdGFza3Ma1AMKBFRhc2sSDgoCaWQYASAB'
    'KAlSAmlkEhIKBG5hbWUYAiABKAlSBG5hbWUSIAoLZGVzY3JpcHRpb24YAyABKAlSC2Rlc2NyaX'
    'B0aW9uElMKCHN1YnRhc2tzGAQgAygLMjcuc2VydmljZXMudGFza3Nfc3ZjLnYxLkdldEFsbFRh'
    'c2tzUmVzcG9uc2UuVGFzay5TdWJUYXNrUghzdWJ0YXNrcxI5CgZzdGF0dXMYBSABKA4yIS5zZX'
    'J2aWNlcy50YXNrc19zdmMudjEuVGFza1N0YXR1c1IGc3RhdHVzEjkKCmNyZWF0ZWRfYXQYBiAB'
    'KAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgljcmVhdGVkQXQSFgoGcHVibGljGAcgAS'
    'gIUgZwdWJsaWMSNgoGZHVlX2F0GAggASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEgA'
    'UgVkdWVBdIgBARIdCgpjcmVhdGVkX2J5GAkgASgJUgljcmVhdGVkQnkaQQoHU3ViVGFzaxIOCg'
    'JpZBgBIAEoCVICaWQSEgoEbmFtZRgCIAEoCVIEbmFtZRISCgRkb25lGAMgASgIUgRkb25lQgkK'
    'B19kdWVfYXQ=');

@$core.Deprecated('Use createSubtaskRequestDescriptor instead')
const CreateSubtaskRequest$json = {
  '1': 'CreateSubtaskRequest',
  '2': [
    {'1': 'task_id', '3': 1, '4': 1, '5': 9, '10': 'taskId'},
    {'1': 'subtask', '3': 2, '4': 1, '5': 11, '6': '.services.tasks_svc.v1.CreateSubtaskRequest.Subtask', '10': 'subtask'},
  ],
  '3': [CreateSubtaskRequest_Subtask$json],
};

@$core.Deprecated('Use createSubtaskRequestDescriptor instead')
const CreateSubtaskRequest_Subtask$json = {
  '1': 'Subtask',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    {'1': 'done', '3': 2, '4': 1, '5': 8, '9': 0, '10': 'done', '17': true},
  ],
  '8': [
    {'1': '_done'},
  ],
};

/// Descriptor for `CreateSubtaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createSubtaskRequestDescriptor = $convert.base64Decode(
    'ChRDcmVhdGVTdWJ0YXNrUmVxdWVzdBIXCgd0YXNrX2lkGAEgASgJUgZ0YXNrSWQSTQoHc3VidG'
    'FzaxgCIAEoCzIzLnNlcnZpY2VzLnRhc2tzX3N2Yy52MS5DcmVhdGVTdWJ0YXNrUmVxdWVzdC5T'
    'dWJ0YXNrUgdzdWJ0YXNrGj8KB1N1YnRhc2sSEgoEbmFtZRgBIAEoCVIEbmFtZRIXCgRkb25lGA'
    'IgASgISABSBGRvbmWIAQFCBwoFX2RvbmU=');

@$core.Deprecated('Use createSubtaskResponseDescriptor instead')
const CreateSubtaskResponse$json = {
  '1': 'CreateSubtaskResponse',
  '2': [
    {'1': 'subtask_id', '3': 1, '4': 1, '5': 9, '10': 'subtaskId'},
  ],
};

/// Descriptor for `CreateSubtaskResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createSubtaskResponseDescriptor = $convert.base64Decode(
    'ChVDcmVhdGVTdWJ0YXNrUmVzcG9uc2USHQoKc3VidGFza19pZBgBIAEoCVIJc3VidGFza0lk');

@$core.Deprecated('Use updateSubtaskRequestDescriptor instead')
const UpdateSubtaskRequest$json = {
  '1': 'UpdateSubtaskRequest',
  '2': [
    {'1': 'task_id', '3': 1, '4': 1, '5': 9, '10': 'taskId'},
    {'1': 'subtask_id', '3': 2, '4': 1, '5': 9, '10': 'subtaskId'},
    {'1': 'subtask', '3': 3, '4': 1, '5': 11, '6': '.services.tasks_svc.v1.UpdateSubtaskRequest.Subtask', '10': 'subtask'},
  ],
  '3': [UpdateSubtaskRequest_Subtask$json],
};

@$core.Deprecated('Use updateSubtaskRequestDescriptor instead')
const UpdateSubtaskRequest_Subtask$json = {
  '1': 'Subtask',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '9': 0, '10': 'name', '17': true},
    {'1': 'done', '3': 2, '4': 1, '5': 8, '9': 1, '10': 'done', '17': true},
  ],
  '8': [
    {'1': '_name'},
    {'1': '_done'},
  ],
};

/// Descriptor for `UpdateSubtaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateSubtaskRequestDescriptor = $convert.base64Decode(
    'ChRVcGRhdGVTdWJ0YXNrUmVxdWVzdBIXCgd0YXNrX2lkGAEgASgJUgZ0YXNrSWQSHQoKc3VidG'
    'Fza19pZBgCIAEoCVIJc3VidGFza0lkEk0KB3N1YnRhc2sYAyABKAsyMy5zZXJ2aWNlcy50YXNr'
    'c19zdmMudjEuVXBkYXRlU3VidGFza1JlcXVlc3QuU3VidGFza1IHc3VidGFzaxpNCgdTdWJ0YX'
    'NrEhcKBG5hbWUYASABKAlIAFIEbmFtZYgBARIXCgRkb25lGAIgASgISAFSBGRvbmWIAQFCBwoF'
    'X25hbWVCBwoFX2RvbmU=');

@$core.Deprecated('Use updateSubtaskResponseDescriptor instead')
const UpdateSubtaskResponse$json = {
  '1': 'UpdateSubtaskResponse',
};

/// Descriptor for `UpdateSubtaskResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateSubtaskResponseDescriptor = $convert.base64Decode(
    'ChVVcGRhdGVTdWJ0YXNrUmVzcG9uc2U=');

@$core.Deprecated('Use deleteSubtaskRequestDescriptor instead')
const DeleteSubtaskRequest$json = {
  '1': 'DeleteSubtaskRequest',
  '2': [
    {'1': 'task_id', '3': 1, '4': 1, '5': 9, '10': 'taskId'},
    {'1': 'subtask_id', '3': 2, '4': 1, '5': 9, '10': 'subtaskId'},
  ],
};

/// Descriptor for `DeleteSubtaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteSubtaskRequestDescriptor = $convert.base64Decode(
    'ChREZWxldGVTdWJ0YXNrUmVxdWVzdBIXCgd0YXNrX2lkGAEgASgJUgZ0YXNrSWQSHQoKc3VidG'
    'Fza19pZBgCIAEoCVIJc3VidGFza0lk');

@$core.Deprecated('Use deleteSubtaskResponseDescriptor instead')
const DeleteSubtaskResponse$json = {
  '1': 'DeleteSubtaskResponse',
};

/// Descriptor for `DeleteSubtaskResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteSubtaskResponseDescriptor = $convert.base64Decode(
    'ChVEZWxldGVTdWJ0YXNrUmVzcG9uc2U=');

@$core.Deprecated('Use removeTaskDueDateRequestDescriptor instead')
const RemoveTaskDueDateRequest$json = {
  '1': 'RemoveTaskDueDateRequest',
  '2': [
    {'1': 'task_id', '3': 1, '4': 1, '5': 9, '10': 'taskId'},
  ],
};

/// Descriptor for `RemoveTaskDueDateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List removeTaskDueDateRequestDescriptor = $convert.base64Decode(
    'ChhSZW1vdmVUYXNrRHVlRGF0ZVJlcXVlc3QSFwoHdGFza19pZBgBIAEoCVIGdGFza0lk');

@$core.Deprecated('Use removeTaskDueDateResponseDescriptor instead')
const RemoveTaskDueDateResponse$json = {
  '1': 'RemoveTaskDueDateResponse',
};

/// Descriptor for `RemoveTaskDueDateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List removeTaskDueDateResponseDescriptor = $convert.base64Decode(
    'ChlSZW1vdmVUYXNrRHVlRGF0ZVJlc3BvbnNl');

@$core.Deprecated('Use deleteTaskRequestDescriptor instead')
const DeleteTaskRequest$json = {
  '1': 'DeleteTaskRequest',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `DeleteTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteTaskRequestDescriptor = $convert.base64Decode(
    'ChFEZWxldGVUYXNrUmVxdWVzdBIOCgJpZBgBIAEoCVICaWQ=');

@$core.Deprecated('Use deleteTaskResponseDescriptor instead')
const DeleteTaskResponse$json = {
  '1': 'DeleteTaskResponse',
};

/// Descriptor for `DeleteTaskResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteTaskResponseDescriptor = $convert.base64Decode(
    'ChJEZWxldGVUYXNrUmVzcG9uc2U=');

