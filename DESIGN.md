# ToDo Application - Design document

## API specification

This API adheres to the Google API Improvement Proposals (AIP) guidelines, ensuring a standardized and efficient design.
Below is a table outlining the various endpoints for CRUD operations within the application.

| Endpoint                  | Service            | Method       | AIP                                   |
|---------------------------|--------------------|--------------|---------------------------------------|
| GET /v1/tasks             | TasksReaderService | GetTask      | [AIP-131](https://google.aip.dev/131) |
| GET /v1/tasks/{id}        | TasksReaderService | ListTasks    | [AIP-132](https://google.aip.dev/132) |
| POST /v1/tasks            | TasksWriterService | CreateTask   | [AIP-133](https://google.aip.dev/133) |
| PATCH /v1/tasks/{task.id} | TasksWriterService | UpdateTask   | [AIP-134](https://google.aip.dev/134) |
| DELETE /v1/tasks/{id}     | TasksWriterService | DeleteTask   | [AIP-135](https://google.aip.dev/135) |
| POST /v1/tasks:undelete   | TasksWriterService | UndeleteTask | [AIP-164](https://google.aip.dev/164) |

OpenAPI reference: [api/api.swagger.yaml](api/api.swagger.yaml)



