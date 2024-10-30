//
//  Generated code. Do not modify.
//  source: services/tasks_svc/v1/task_svc.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'task_svc.pb.dart' as $0;

export 'task_svc.pb.dart';

@$pb.GrpcServiceName('services.tasks_svc.v1.TaskService')
class TaskServiceClient extends $grpc.Client {
  static final _$createTask = $grpc.ClientMethod<$0.CreateTaskRequest, $0.CreateTaskResponse>(
      '/services.tasks_svc.v1.TaskService/CreateTask',
      ($0.CreateTaskRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CreateTaskResponse.fromBuffer(value));
  static final _$updateTask = $grpc.ClientMethod<$0.UpdateTaskRequest, $0.UpdateTaskResponse>(
      '/services.tasks_svc.v1.TaskService/UpdateTask',
      ($0.UpdateTaskRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.UpdateTaskResponse.fromBuffer(value));
  static final _$getTask = $grpc.ClientMethod<$0.GetTaskRequest, $0.GetTaskResponse>(
      '/services.tasks_svc.v1.TaskService/GetTask',
      ($0.GetTaskRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.GetTaskResponse.fromBuffer(value));
  static final _$getAllTasks = $grpc.ClientMethod<$0.GetAllTasksRequest, $0.GetAllTasksResponse>(
      '/services.tasks_svc.v1.TaskService/GetAllTasks',
      ($0.GetAllTasksRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.GetAllTasksResponse.fromBuffer(value));
  static final _$createSubtask = $grpc.ClientMethod<$0.CreateSubtaskRequest, $0.CreateSubtaskResponse>(
      '/services.tasks_svc.v1.TaskService/CreateSubtask',
      ($0.CreateSubtaskRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CreateSubtaskResponse.fromBuffer(value));
  static final _$updateSubtask = $grpc.ClientMethod<$0.UpdateSubtaskRequest, $0.UpdateSubtaskResponse>(
      '/services.tasks_svc.v1.TaskService/UpdateSubtask',
      ($0.UpdateSubtaskRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.UpdateSubtaskResponse.fromBuffer(value));
  static final _$deleteSubtask = $grpc.ClientMethod<$0.DeleteSubtaskRequest, $0.DeleteSubtaskResponse>(
      '/services.tasks_svc.v1.TaskService/DeleteSubtask',
      ($0.DeleteSubtaskRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.DeleteSubtaskResponse.fromBuffer(value));
  static final _$removeTaskDueDate = $grpc.ClientMethod<$0.RemoveTaskDueDateRequest, $0.RemoveTaskDueDateResponse>(
      '/services.tasks_svc.v1.TaskService/RemoveTaskDueDate',
      ($0.RemoveTaskDueDateRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.RemoveTaskDueDateResponse.fromBuffer(value));
  static final _$deleteTask = $grpc.ClientMethod<$0.DeleteTaskRequest, $0.DeleteTaskResponse>(
      '/services.tasks_svc.v1.TaskService/DeleteTask',
      ($0.DeleteTaskRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.DeleteTaskResponse.fromBuffer(value));

  TaskServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.CreateTaskResponse> createTask($0.CreateTaskRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createTask, request, options: options);
  }

  $grpc.ResponseFuture<$0.UpdateTaskResponse> updateTask($0.UpdateTaskRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateTask, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetTaskResponse> getTask($0.GetTaskRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getTask, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetAllTasksResponse> getAllTasks($0.GetAllTasksRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getAllTasks, request, options: options);
  }

  $grpc.ResponseFuture<$0.CreateSubtaskResponse> createSubtask($0.CreateSubtaskRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createSubtask, request, options: options);
  }

  $grpc.ResponseFuture<$0.UpdateSubtaskResponse> updateSubtask($0.UpdateSubtaskRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateSubtask, request, options: options);
  }

  $grpc.ResponseFuture<$0.DeleteSubtaskResponse> deleteSubtask($0.DeleteSubtaskRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteSubtask, request, options: options);
  }

  $grpc.ResponseFuture<$0.RemoveTaskDueDateResponse> removeTaskDueDate($0.RemoveTaskDueDateRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$removeTaskDueDate, request, options: options);
  }

  $grpc.ResponseFuture<$0.DeleteTaskResponse> deleteTask($0.DeleteTaskRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteTask, request, options: options);
  }
}

@$pb.GrpcServiceName('services.tasks_svc.v1.TaskService')
abstract class TaskServiceBase extends $grpc.Service {
  $core.String get $name => 'services.tasks_svc.v1.TaskService';

  TaskServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.CreateTaskRequest, $0.CreateTaskResponse>(
        'CreateTask',
        createTask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateTaskRequest.fromBuffer(value),
        ($0.CreateTaskResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UpdateTaskRequest, $0.UpdateTaskResponse>(
        'UpdateTask',
        updateTask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UpdateTaskRequest.fromBuffer(value),
        ($0.UpdateTaskResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetTaskRequest, $0.GetTaskResponse>(
        'GetTask',
        getTask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetTaskRequest.fromBuffer(value),
        ($0.GetTaskResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetAllTasksRequest, $0.GetAllTasksResponse>(
        'GetAllTasks',
        getAllTasks_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetAllTasksRequest.fromBuffer(value),
        ($0.GetAllTasksResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CreateSubtaskRequest, $0.CreateSubtaskResponse>(
        'CreateSubtask',
        createSubtask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateSubtaskRequest.fromBuffer(value),
        ($0.CreateSubtaskResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UpdateSubtaskRequest, $0.UpdateSubtaskResponse>(
        'UpdateSubtask',
        updateSubtask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UpdateSubtaskRequest.fromBuffer(value),
        ($0.UpdateSubtaskResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DeleteSubtaskRequest, $0.DeleteSubtaskResponse>(
        'DeleteSubtask',
        deleteSubtask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DeleteSubtaskRequest.fromBuffer(value),
        ($0.DeleteSubtaskResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RemoveTaskDueDateRequest, $0.RemoveTaskDueDateResponse>(
        'RemoveTaskDueDate',
        removeTaskDueDate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RemoveTaskDueDateRequest.fromBuffer(value),
        ($0.RemoveTaskDueDateResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DeleteTaskRequest, $0.DeleteTaskResponse>(
        'DeleteTask',
        deleteTask_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DeleteTaskRequest.fromBuffer(value),
        ($0.DeleteTaskResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.CreateTaskResponse> createTask_Pre($grpc.ServiceCall call, $async.Future<$0.CreateTaskRequest> request) async {
    return createTask(call, await request);
  }

  $async.Future<$0.UpdateTaskResponse> updateTask_Pre($grpc.ServiceCall call, $async.Future<$0.UpdateTaskRequest> request) async {
    return updateTask(call, await request);
  }

  $async.Future<$0.GetTaskResponse> getTask_Pre($grpc.ServiceCall call, $async.Future<$0.GetTaskRequest> request) async {
    return getTask(call, await request);
  }

  $async.Future<$0.GetAllTasksResponse> getAllTasks_Pre($grpc.ServiceCall call, $async.Future<$0.GetAllTasksRequest> request) async {
    return getAllTasks(call, await request);
  }

  $async.Future<$0.CreateSubtaskResponse> createSubtask_Pre($grpc.ServiceCall call, $async.Future<$0.CreateSubtaskRequest> request) async {
    return createSubtask(call, await request);
  }

  $async.Future<$0.UpdateSubtaskResponse> updateSubtask_Pre($grpc.ServiceCall call, $async.Future<$0.UpdateSubtaskRequest> request) async {
    return updateSubtask(call, await request);
  }

  $async.Future<$0.DeleteSubtaskResponse> deleteSubtask_Pre($grpc.ServiceCall call, $async.Future<$0.DeleteSubtaskRequest> request) async {
    return deleteSubtask(call, await request);
  }

  $async.Future<$0.RemoveTaskDueDateResponse> removeTaskDueDate_Pre($grpc.ServiceCall call, $async.Future<$0.RemoveTaskDueDateRequest> request) async {
    return removeTaskDueDate(call, await request);
  }

  $async.Future<$0.DeleteTaskResponse> deleteTask_Pre($grpc.ServiceCall call, $async.Future<$0.DeleteTaskRequest> request) async {
    return deleteTask(call, await request);
  }

  $async.Future<$0.CreateTaskResponse> createTask($grpc.ServiceCall call, $0.CreateTaskRequest request);
  $async.Future<$0.UpdateTaskResponse> updateTask($grpc.ServiceCall call, $0.UpdateTaskRequest request);
  $async.Future<$0.GetTaskResponse> getTask($grpc.ServiceCall call, $0.GetTaskRequest request);
  $async.Future<$0.GetAllTasksResponse> getAllTasks($grpc.ServiceCall call, $0.GetAllTasksRequest request);
  $async.Future<$0.CreateSubtaskResponse> createSubtask($grpc.ServiceCall call, $0.CreateSubtaskRequest request);
  $async.Future<$0.UpdateSubtaskResponse> updateSubtask($grpc.ServiceCall call, $0.UpdateSubtaskRequest request);
  $async.Future<$0.DeleteSubtaskResponse> deleteSubtask($grpc.ServiceCall call, $0.DeleteSubtaskRequest request);
  $async.Future<$0.RemoveTaskDueDateResponse> removeTaskDueDate($grpc.ServiceCall call, $0.RemoveTaskDueDateRequest request);
  $async.Future<$0.DeleteTaskResponse> deleteTask($grpc.ServiceCall call, $0.DeleteTaskRequest request);
}
