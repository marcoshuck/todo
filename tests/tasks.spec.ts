import {expect, test} from '@playwright/test';
import {createTask, deleteTask, getTask, listTasks, undeleteTask, updateTask} from "./tasks.utils";
import {Task} from "../api/tasks/v1/tasks_pb";


test('POST /v1/tasks', async ({request}) => {
    let input = {
        title: 'An awesome task',
        description: 'An awesome description for an awesome task',
    };

    let output = await createTask(request, input);

    // Assertions
    expect(output.title).toEqual(input.title);
    expect(output.description).toEqual(input.description);
    expect(output.id).toBeGreaterThan(0);
    expect(output.createTime).not.toBeNull();
    expect(output.updateTime).not.toBeNull();
})

test('GET /v1/tasks/:id', async ({request}) => {
    let input = {
        title: 'An awesome task',
        description: 'An awesome description for an awesome task',
    };

    let expected = await createTask(request, input);


    // Get the data
    const {data: output, response} = await getTask(request, expected.id);
    expect(response.ok()).toBeTruthy();

    // Assertions
    expect(output.title).toEqual(expected.title);
    expect(output.description).toEqual(expected.description);
    expect(output.id).toEqual(expected.id);
})

test('GET /v1/tasks', async ({request}) => {
    let input = {
        title: 'An awesome task',
        description: 'An awesome description for an awesome task',
    };

    const list: Task[] = [];

    for (let i = 0; i < 10; i++) {
        let task = await createTask(request, input);
        list.push(task);
        await new Promise(r => setTimeout(r, 1000));
    }

    let response = await listTasks(request, 5, undefined);

    expect(response.data.nextPageToken).not.toBe('');


    for (const task of list) {
        await deleteTask(request, task.id);
    }
})

test('DELETE /v1/tasks/:id', async ({request}) => {
    let input = {
        title: 'An awesome task',
        description: 'An awesome description for an awesome task',
    };

    let expected = await createTask(request, input);

    // Get the task
    let {data: output, response} = await getTask(request, expected.id);
    expect(response.ok()).toBeTruthy();

    // Delete the task
    await deleteTask(request, output.id);

    response = await request.get(`/v1/tasks/${output.id}`)
    expect(response.ok()).toBeFalsy();
    expect(response.status()).toEqual(404);
})

test('POST /v1/tasks:undelete', async ({request}) => {
    let input = {
        title: 'An awesome task',
        description: 'An awesome description for an awesome task',
    };

    let expected = await createTask(request, input);

    // Get the task
    let {data: output, response} = await getTask(request, expected.id);
    expect(response.ok()).toBeTruthy();

    // Delete the task
    await deleteTask(request, output.id);

    // It does not exist, hence not found
    response = await request.get(`/v1/tasks/${output.id}`)
    expect(response.ok()).toBeFalsy();
    expect(response.status()).toEqual(404);

    // But if we undelete
    await undeleteTask(request, output.id);

    // We rever the state back to found
    response = await request.get(`/v1/tasks/${output.id}`)
    expect(response.ok()).toBeTruthy();
})

test('PATCH /v1/tasks/:id', async ({request}) => {
    let input = {
        title: 'An awesome task',
        description: 'An awesome description for an awesome task',
    };

    let expected = await createTask(request, input);

    // Get the task
    let {data: output, response} = await getTask(request, expected.id);
    expect(response.ok()).toBeTruthy();

    // Partially update a task
    const patch = {
        id: Number(output.id),
        description: 'A modified description for an awesome task'
    }
    await updateTask(request, output.id, patch);

    ({data: output, response} = await getTask(request, expected.id));
    expect(response.ok()).toBeTruthy();
    expect(output.description).toEqual(patch.description);
})