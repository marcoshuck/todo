import {expect, test} from '@playwright/test';
import {Task} from '../gen/typescript/api/tasks/v1/tasks_pb';

test('POST /v1/tasks', async ({request}) => {
    let input = {
        title: 'An awesome task',
        description: 'An awesome description for an awesome task',
    };

    const response = await request.post('/v1/tasks', {
        data: input,
    });
    expect(response.ok()).toBeTruthy();

    const body = await response.body();
    const output: Task = Task.fromJsonString(body.toString());

    expect(output.title).toEqual(input.title);
    expect(output.description).toEqual(input.description);
    expect(output.id).toBeGreaterThan(0);
    expect(output.createTime).not.toBeNull();
    expect(output.updateTime).not.toBeNull();
})