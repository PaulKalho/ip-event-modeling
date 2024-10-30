import * as jspb from 'google-protobuf'

import * as services_tasks_svc_v1_types_pb from '../../../services/tasks_svc/v1/types_pb'; // proto import: "services/tasks_svc/v1/types.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"


export class CreateTaskRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateTaskRequest;

  getDescription(): string;
  setDescription(value: string): CreateTaskRequest;
  hasDescription(): boolean;
  clearDescription(): CreateTaskRequest;

  getPublic(): boolean;
  setPublic(value: boolean): CreateTaskRequest;
  hasPublic(): boolean;
  clearPublic(): CreateTaskRequest;

  getDueAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDueAt(value?: google_protobuf_timestamp_pb.Timestamp): CreateTaskRequest;
  hasDueAt(): boolean;
  clearDueAt(): CreateTaskRequest;

  getInitialStatus(): services_tasks_svc_v1_types_pb.TaskStatus;
  setInitialStatus(value: services_tasks_svc_v1_types_pb.TaskStatus): CreateTaskRequest;
  hasInitialStatus(): boolean;
  clearInitialStatus(): CreateTaskRequest;

  getSubtasksList(): Array<CreateTaskRequest.SubTask>;
  setSubtasksList(value: Array<CreateTaskRequest.SubTask>): CreateTaskRequest;
  clearSubtasksList(): CreateTaskRequest;
  addSubtasks(value?: CreateTaskRequest.SubTask, index?: number): CreateTaskRequest.SubTask;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTaskRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTaskRequest): CreateTaskRequest.AsObject;
  static serializeBinaryToWriter(message: CreateTaskRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTaskRequest;
  static deserializeBinaryFromReader(message: CreateTaskRequest, reader: jspb.BinaryReader): CreateTaskRequest;
}

export namespace CreateTaskRequest {
  export type AsObject = {
    name: string,
    description?: string,
    pb_public?: boolean,
    dueAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    initialStatus?: services_tasks_svc_v1_types_pb.TaskStatus,
    subtasksList: Array<CreateTaskRequest.SubTask.AsObject>,
  }

  export class SubTask extends jspb.Message {
    getName(): string;
    setName(value: string): SubTask;

    getDone(): boolean;
    setDone(value: boolean): SubTask;
    hasDone(): boolean;
    clearDone(): SubTask;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SubTask.AsObject;
    static toObject(includeInstance: boolean, msg: SubTask): SubTask.AsObject;
    static serializeBinaryToWriter(message: SubTask, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SubTask;
    static deserializeBinaryFromReader(message: SubTask, reader: jspb.BinaryReader): SubTask;
  }

  export namespace SubTask {
    export type AsObject = {
      name: string,
      done?: boolean,
    }

    export enum DoneCase { 
      _DONE_NOT_SET = 0,
      DONE = 2,
    }
  }


  export enum DescriptionCase { 
    _DESCRIPTION_NOT_SET = 0,
    DESCRIPTION = 2,
  }

  export enum PublicCase { 
    _PUBLIC_NOT_SET = 0,
    PUBLIC = 4,
  }

  export enum InitialStatusCase { 
    _INITIAL_STATUS_NOT_SET = 0,
    INITIAL_STATUS = 6,
  }
}

export class CreateTaskResponse extends jspb.Message {
  getId(): string;
  setId(value: string): CreateTaskResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTaskResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTaskResponse): CreateTaskResponse.AsObject;
  static serializeBinaryToWriter(message: CreateTaskResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTaskResponse;
  static deserializeBinaryFromReader(message: CreateTaskResponse, reader: jspb.BinaryReader): CreateTaskResponse;
}

export namespace CreateTaskResponse {
  export type AsObject = {
    id: string,
  }
}

export class UpdateTaskRequest extends jspb.Message {
  getId(): string;
  setId(value: string): UpdateTaskRequest;

  getName(): string;
  setName(value: string): UpdateTaskRequest;
  hasName(): boolean;
  clearName(): UpdateTaskRequest;

  getDescription(): string;
  setDescription(value: string): UpdateTaskRequest;
  hasDescription(): boolean;
  clearDescription(): UpdateTaskRequest;

  getDueAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDueAt(value?: google_protobuf_timestamp_pb.Timestamp): UpdateTaskRequest;
  hasDueAt(): boolean;
  clearDueAt(): UpdateTaskRequest;

  getStatus(): services_tasks_svc_v1_types_pb.TaskStatus;
  setStatus(value: services_tasks_svc_v1_types_pb.TaskStatus): UpdateTaskRequest;
  hasStatus(): boolean;
  clearStatus(): UpdateTaskRequest;

  getPublic(): boolean;
  setPublic(value: boolean): UpdateTaskRequest;
  hasPublic(): boolean;
  clearPublic(): UpdateTaskRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTaskRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTaskRequest): UpdateTaskRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateTaskRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTaskRequest;
  static deserializeBinaryFromReader(message: UpdateTaskRequest, reader: jspb.BinaryReader): UpdateTaskRequest;
}

export namespace UpdateTaskRequest {
  export type AsObject = {
    id: string,
    name?: string,
    description?: string,
    dueAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    status?: services_tasks_svc_v1_types_pb.TaskStatus,
    pb_public?: boolean,
  }

  export enum NameCase { 
    _NAME_NOT_SET = 0,
    NAME = 2,
  }

  export enum DescriptionCase { 
    _DESCRIPTION_NOT_SET = 0,
    DESCRIPTION = 3,
  }

  export enum DueAtCase { 
    _DUE_AT_NOT_SET = 0,
    DUE_AT = 4,
  }

  export enum StatusCase { 
    _STATUS_NOT_SET = 0,
    STATUS = 5,
  }

  export enum PublicCase { 
    _PUBLIC_NOT_SET = 0,
    PUBLIC = 6,
  }
}

export class UpdateTaskResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTaskResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTaskResponse): UpdateTaskResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateTaskResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTaskResponse;
  static deserializeBinaryFromReader(message: UpdateTaskResponse, reader: jspb.BinaryReader): UpdateTaskResponse;
}

export namespace UpdateTaskResponse {
  export type AsObject = {
  }
}

export class GetTaskRequest extends jspb.Message {
  getId(): string;
  setId(value: string): GetTaskRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTaskRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTaskRequest): GetTaskRequest.AsObject;
  static serializeBinaryToWriter(message: GetTaskRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTaskRequest;
  static deserializeBinaryFromReader(message: GetTaskRequest, reader: jspb.BinaryReader): GetTaskRequest;
}

export namespace GetTaskRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetTaskResponse extends jspb.Message {
  getId(): string;
  setId(value: string): GetTaskResponse;

  getName(): string;
  setName(value: string): GetTaskResponse;

  getDescription(): string;
  setDescription(value: string): GetTaskResponse;

  getSubtasksList(): Array<GetTaskResponse.SubTask>;
  setSubtasksList(value: Array<GetTaskResponse.SubTask>): GetTaskResponse;
  clearSubtasksList(): GetTaskResponse;
  addSubtasks(value?: GetTaskResponse.SubTask, index?: number): GetTaskResponse.SubTask;

  getStatus(): services_tasks_svc_v1_types_pb.TaskStatus;
  setStatus(value: services_tasks_svc_v1_types_pb.TaskStatus): GetTaskResponse;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): GetTaskResponse;
  hasCreatedAt(): boolean;
  clearCreatedAt(): GetTaskResponse;

  getPublic(): boolean;
  setPublic(value: boolean): GetTaskResponse;

  getDueAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDueAt(value?: google_protobuf_timestamp_pb.Timestamp): GetTaskResponse;
  hasDueAt(): boolean;
  clearDueAt(): GetTaskResponse;

  getCreatedBy(): string;
  setCreatedBy(value: string): GetTaskResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTaskResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTaskResponse): GetTaskResponse.AsObject;
  static serializeBinaryToWriter(message: GetTaskResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTaskResponse;
  static deserializeBinaryFromReader(message: GetTaskResponse, reader: jspb.BinaryReader): GetTaskResponse;
}

export namespace GetTaskResponse {
  export type AsObject = {
    id: string,
    name: string,
    description: string,
    subtasksList: Array<GetTaskResponse.SubTask.AsObject>,
    status: services_tasks_svc_v1_types_pb.TaskStatus,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    pb_public: boolean,
    dueAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    createdBy: string,
  }

  export class SubTask extends jspb.Message {
    getId(): string;
    setId(value: string): SubTask;

    getName(): string;
    setName(value: string): SubTask;

    getDone(): boolean;
    setDone(value: boolean): SubTask;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SubTask.AsObject;
    static toObject(includeInstance: boolean, msg: SubTask): SubTask.AsObject;
    static serializeBinaryToWriter(message: SubTask, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SubTask;
    static deserializeBinaryFromReader(message: SubTask, reader: jspb.BinaryReader): SubTask;
  }

  export namespace SubTask {
    export type AsObject = {
      id: string,
      name: string,
      done: boolean,
    }
  }


  export enum DueAtCase { 
    _DUE_AT_NOT_SET = 0,
    DUE_AT = 8,
  }
}

export class GetAllTasksRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllTasksRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllTasksRequest): GetAllTasksRequest.AsObject;
  static serializeBinaryToWriter(message: GetAllTasksRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllTasksRequest;
  static deserializeBinaryFromReader(message: GetAllTasksRequest, reader: jspb.BinaryReader): GetAllTasksRequest;
}

export namespace GetAllTasksRequest {
  export type AsObject = {
  }
}

export class GetAllTasksResponse extends jspb.Message {
  getTasksList(): Array<GetAllTasksResponse.Task>;
  setTasksList(value: Array<GetAllTasksResponse.Task>): GetAllTasksResponse;
  clearTasksList(): GetAllTasksResponse;
  addTasks(value?: GetAllTasksResponse.Task, index?: number): GetAllTasksResponse.Task;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllTasksResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllTasksResponse): GetAllTasksResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllTasksResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllTasksResponse;
  static deserializeBinaryFromReader(message: GetAllTasksResponse, reader: jspb.BinaryReader): GetAllTasksResponse;
}

export namespace GetAllTasksResponse {
  export type AsObject = {
    tasksList: Array<GetAllTasksResponse.Task.AsObject>,
  }

  export class Task extends jspb.Message {
    getId(): string;
    setId(value: string): Task;

    getName(): string;
    setName(value: string): Task;

    getDescription(): string;
    setDescription(value: string): Task;

    getSubtasksList(): Array<GetAllTasksResponse.Task.SubTask>;
    setSubtasksList(value: Array<GetAllTasksResponse.Task.SubTask>): Task;
    clearSubtasksList(): Task;
    addSubtasks(value?: GetAllTasksResponse.Task.SubTask, index?: number): GetAllTasksResponse.Task.SubTask;

    getStatus(): services_tasks_svc_v1_types_pb.TaskStatus;
    setStatus(value: services_tasks_svc_v1_types_pb.TaskStatus): Task;

    getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Task;
    hasCreatedAt(): boolean;
    clearCreatedAt(): Task;

    getPublic(): boolean;
    setPublic(value: boolean): Task;

    getDueAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setDueAt(value?: google_protobuf_timestamp_pb.Timestamp): Task;
    hasDueAt(): boolean;
    clearDueAt(): Task;

    getCreatedBy(): string;
    setCreatedBy(value: string): Task;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Task.AsObject;
    static toObject(includeInstance: boolean, msg: Task): Task.AsObject;
    static serializeBinaryToWriter(message: Task, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Task;
    static deserializeBinaryFromReader(message: Task, reader: jspb.BinaryReader): Task;
  }

  export namespace Task {
    export type AsObject = {
      id: string,
      name: string,
      description: string,
      subtasksList: Array<GetAllTasksResponse.Task.SubTask.AsObject>,
      status: services_tasks_svc_v1_types_pb.TaskStatus,
      createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
      pb_public: boolean,
      dueAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
      createdBy: string,
    }

    export class SubTask extends jspb.Message {
      getId(): string;
      setId(value: string): SubTask;

      getName(): string;
      setName(value: string): SubTask;

      getDone(): boolean;
      setDone(value: boolean): SubTask;

      serializeBinary(): Uint8Array;
      toObject(includeInstance?: boolean): SubTask.AsObject;
      static toObject(includeInstance: boolean, msg: SubTask): SubTask.AsObject;
      static serializeBinaryToWriter(message: SubTask, writer: jspb.BinaryWriter): void;
      static deserializeBinary(bytes: Uint8Array): SubTask;
      static deserializeBinaryFromReader(message: SubTask, reader: jspb.BinaryReader): SubTask;
    }

    export namespace SubTask {
      export type AsObject = {
        id: string,
        name: string,
        done: boolean,
      }
    }


    export enum DueAtCase { 
      _DUE_AT_NOT_SET = 0,
      DUE_AT = 8,
    }
  }

}

export class CreateSubtaskRequest extends jspb.Message {
  getTaskId(): string;
  setTaskId(value: string): CreateSubtaskRequest;

  getSubtask(): CreateSubtaskRequest.Subtask | undefined;
  setSubtask(value?: CreateSubtaskRequest.Subtask): CreateSubtaskRequest;
  hasSubtask(): boolean;
  clearSubtask(): CreateSubtaskRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSubtaskRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSubtaskRequest): CreateSubtaskRequest.AsObject;
  static serializeBinaryToWriter(message: CreateSubtaskRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSubtaskRequest;
  static deserializeBinaryFromReader(message: CreateSubtaskRequest, reader: jspb.BinaryReader): CreateSubtaskRequest;
}

export namespace CreateSubtaskRequest {
  export type AsObject = {
    taskId: string,
    subtask?: CreateSubtaskRequest.Subtask.AsObject,
  }

  export class Subtask extends jspb.Message {
    getName(): string;
    setName(value: string): Subtask;

    getDone(): boolean;
    setDone(value: boolean): Subtask;
    hasDone(): boolean;
    clearDone(): Subtask;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Subtask.AsObject;
    static toObject(includeInstance: boolean, msg: Subtask): Subtask.AsObject;
    static serializeBinaryToWriter(message: Subtask, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Subtask;
    static deserializeBinaryFromReader(message: Subtask, reader: jspb.BinaryReader): Subtask;
  }

  export namespace Subtask {
    export type AsObject = {
      name: string,
      done?: boolean,
    }

    export enum DoneCase { 
      _DONE_NOT_SET = 0,
      DONE = 2,
    }
  }

}

export class CreateSubtaskResponse extends jspb.Message {
  getSubtaskId(): string;
  setSubtaskId(value: string): CreateSubtaskResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSubtaskResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSubtaskResponse): CreateSubtaskResponse.AsObject;
  static serializeBinaryToWriter(message: CreateSubtaskResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSubtaskResponse;
  static deserializeBinaryFromReader(message: CreateSubtaskResponse, reader: jspb.BinaryReader): CreateSubtaskResponse;
}

export namespace CreateSubtaskResponse {
  export type AsObject = {
    subtaskId: string,
  }
}

export class UpdateSubtaskRequest extends jspb.Message {
  getTaskId(): string;
  setTaskId(value: string): UpdateSubtaskRequest;

  getSubtaskId(): string;
  setSubtaskId(value: string): UpdateSubtaskRequest;

  getSubtask(): UpdateSubtaskRequest.Subtask | undefined;
  setSubtask(value?: UpdateSubtaskRequest.Subtask): UpdateSubtaskRequest;
  hasSubtask(): boolean;
  clearSubtask(): UpdateSubtaskRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateSubtaskRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateSubtaskRequest): UpdateSubtaskRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateSubtaskRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateSubtaskRequest;
  static deserializeBinaryFromReader(message: UpdateSubtaskRequest, reader: jspb.BinaryReader): UpdateSubtaskRequest;
}

export namespace UpdateSubtaskRequest {
  export type AsObject = {
    taskId: string,
    subtaskId: string,
    subtask?: UpdateSubtaskRequest.Subtask.AsObject,
  }

  export class Subtask extends jspb.Message {
    getName(): string;
    setName(value: string): Subtask;
    hasName(): boolean;
    clearName(): Subtask;

    getDone(): boolean;
    setDone(value: boolean): Subtask;
    hasDone(): boolean;
    clearDone(): Subtask;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Subtask.AsObject;
    static toObject(includeInstance: boolean, msg: Subtask): Subtask.AsObject;
    static serializeBinaryToWriter(message: Subtask, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Subtask;
    static deserializeBinaryFromReader(message: Subtask, reader: jspb.BinaryReader): Subtask;
  }

  export namespace Subtask {
    export type AsObject = {
      name?: string,
      done?: boolean,
    }

    export enum NameCase { 
      _NAME_NOT_SET = 0,
      NAME = 1,
    }

    export enum DoneCase { 
      _DONE_NOT_SET = 0,
      DONE = 2,
    }
  }

}

export class UpdateSubtaskResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateSubtaskResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateSubtaskResponse): UpdateSubtaskResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateSubtaskResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateSubtaskResponse;
  static deserializeBinaryFromReader(message: UpdateSubtaskResponse, reader: jspb.BinaryReader): UpdateSubtaskResponse;
}

export namespace UpdateSubtaskResponse {
  export type AsObject = {
  }
}

export class DeleteSubtaskRequest extends jspb.Message {
  getTaskId(): string;
  setTaskId(value: string): DeleteSubtaskRequest;

  getSubtaskId(): string;
  setSubtaskId(value: string): DeleteSubtaskRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteSubtaskRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteSubtaskRequest): DeleteSubtaskRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteSubtaskRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteSubtaskRequest;
  static deserializeBinaryFromReader(message: DeleteSubtaskRequest, reader: jspb.BinaryReader): DeleteSubtaskRequest;
}

export namespace DeleteSubtaskRequest {
  export type AsObject = {
    taskId: string,
    subtaskId: string,
  }
}

export class DeleteSubtaskResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteSubtaskResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteSubtaskResponse): DeleteSubtaskResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteSubtaskResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteSubtaskResponse;
  static deserializeBinaryFromReader(message: DeleteSubtaskResponse, reader: jspb.BinaryReader): DeleteSubtaskResponse;
}

export namespace DeleteSubtaskResponse {
  export type AsObject = {
  }
}

export class RemoveTaskDueDateRequest extends jspb.Message {
  getTaskId(): string;
  setTaskId(value: string): RemoveTaskDueDateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveTaskDueDateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveTaskDueDateRequest): RemoveTaskDueDateRequest.AsObject;
  static serializeBinaryToWriter(message: RemoveTaskDueDateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveTaskDueDateRequest;
  static deserializeBinaryFromReader(message: RemoveTaskDueDateRequest, reader: jspb.BinaryReader): RemoveTaskDueDateRequest;
}

export namespace RemoveTaskDueDateRequest {
  export type AsObject = {
    taskId: string,
  }
}

export class RemoveTaskDueDateResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveTaskDueDateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveTaskDueDateResponse): RemoveTaskDueDateResponse.AsObject;
  static serializeBinaryToWriter(message: RemoveTaskDueDateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveTaskDueDateResponse;
  static deserializeBinaryFromReader(message: RemoveTaskDueDateResponse, reader: jspb.BinaryReader): RemoveTaskDueDateResponse;
}

export namespace RemoveTaskDueDateResponse {
  export type AsObject = {
  }
}

export class DeleteTaskRequest extends jspb.Message {
  getId(): string;
  setId(value: string): DeleteTaskRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteTaskRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteTaskRequest): DeleteTaskRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteTaskRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteTaskRequest;
  static deserializeBinaryFromReader(message: DeleteTaskRequest, reader: jspb.BinaryReader): DeleteTaskRequest;
}

export namespace DeleteTaskRequest {
  export type AsObject = {
    id: string,
  }
}

export class DeleteTaskResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteTaskResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteTaskResponse): DeleteTaskResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteTaskResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteTaskResponse;
  static deserializeBinaryFromReader(message: DeleteTaskResponse, reader: jspb.BinaryReader): DeleteTaskResponse;
}

export namespace DeleteTaskResponse {
  export type AsObject = {
  }
}

