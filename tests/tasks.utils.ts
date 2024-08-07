import {APIRequestContext, expect} from "@playwright/test";
import {ListTasksResponseSchema, Task, TaskSchema} from '../api/tasks/v1/tasks_pb';
import {fromJsonString} from "@bufbuild/protobuf";

export async function createTask(request: APIRequestContext, input: any): Promise<Task> {
    // Send the request and wait for the response.
    const response = await request.post('/v1/tasks', {
        data: input,
    });

    // Status: OK
    expect(response.ok()).toBeTruthy();

    // Read the body
    const body = await response.body();

    return fromJsonString(TaskSchema, body.toString());
}

export async function getTask(request: APIRequestContext, id: bigint) {
    const response = await request.get(`/v1/tasks/${id}`);
    const body = await response.body();
    return {
        response: response,
        data: fromJsonString(TaskSchema, body.toString()),
    }
}

export async function listTasks(request: APIRequestContext, size: number, nextPageToken: string | undefined) {
    const response = await request.get(`/v1/tasks`, {
        params: {
            "page_size": size,
            "page_token": nextPageToken,
        },
    });
    const body = await response.body();

    return {
        response: response,
        data: fromJsonString(ListTasksResponseSchema, body.toString()),
    }
}

export async function deleteTask(request: APIRequestContext, id: bigint): Promise<void> {
    const response = await request.delete(`/v1/tasks/${id}`);
    expect(response.ok()).toBeTruthy();
}

export async function undeleteTask(request: APIRequestContext, id: bigint): Promise<void> {
    const response = await request.post('/v1/tasks:undelete', {
        data: {
            id: Number(id),
        }
    });
    expect(response.ok()).toBeTruthy();
}

export async function updateTask(request: APIRequestContext, id: bigint, payload: any): Promise<Task> {
    const response = await request.patch(`/v1/tasks/${id}`, {
        data: payload,
    });
    // Read the body
    const body = await response.body();
    return fromJsonString(TaskSchema, body.toString());
}