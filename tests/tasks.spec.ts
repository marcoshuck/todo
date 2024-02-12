import {expect, test} from '@playwright/test';
import {createTask, deleteTask, getTask, undeleteTask} from "./tasks.utils";


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
    expect(output.createTime).toEqual(expected.createTime);
    expect(output.updateTime).toEqual(expected.updateTime);
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