import { Injectable } from '@angular/core';
import {
  CreateTaskRequest,
  ListTasksResponse,
  Task,
  UpdateTaskRequest
} from "../../api/tasks/v1/tasks_pb";
import {FetchBackend, HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {environment} from "../../environments/environment";

declare interface ITaskService {
  create(task: Task): Observable<Task>;
  get(id: bigint): Observable<Task>;
  list(size: number, cursor?: string): Observable<ListTasksResponse>;
  update(id: bigint, task: Task): Observable<Task>;
  delete(id: bigint): Observable<Task>;
  undelete(id: bigint): Observable<Task>;
}

@Injectable({
  providedIn: 'root'
})
export class TaskService implements ITaskService {
  private http: HttpClient;
  private readonly baseURL: string;

  constructor() {
    this.baseURL = `${environment.apiUrl}/v1/tasks`;
    this.http = new HttpClient(new FetchBackend())
  }

  create(task: Task): Observable<Task> {
    const body: CreateTaskRequest = {
      task: task,
    }
    return this.http.post<Task>(this.baseURL, body);
  }

  get(id: bigint): Observable<Task> {
    return this.http.get<Task>(`${this.baseURL}/${id}`);
  }

  list(size: number, cursor?: string): Observable<ListTasksResponse> {
    const params = {
      'page_size': size || 50,
      'page_token': cursor || '',
    }
    return this.http.get<ListTasksResponse>(this.baseURL, {
      params: params,
    });
  }

  update(id: bigint, task: Task): Observable<Task> {
    const body: UpdateTaskRequest = {
      task: task,
      updateMask: task,
    }
    return this.http.patch<Task>(`${this.baseURL}/${id}`, body);
  }

  delete(id: bigint): Observable<Task> {
    return this.http.delete<Task>(`${this.baseURL}/${id}`);
  }

  undelete(id: bigint): Observable<Task> {
    return this.http.post<Task>(`${this.baseURL}/${id}:undelete`, null);
  }
}
