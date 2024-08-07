import {Component, effect, model, signal, WritableSignal} from '@angular/core';
import {MainLayoutComponent} from "../../layouts/main-layout/main-layout.component";
import {
  MatCard,
  MatCardActions,
  MatCardContent,
  MatCardFooter,
  MatCardHeader,
  MatCardTitle
} from "@angular/material/card";
import {TaskService} from "../../tasks/task.service";
import {ListTasksResponse, Task} from "../../../api/tasks/v1/tasks_pb";
import {MatPaginator, PageEvent} from "@angular/material/paginator";
import {TaskListComponent} from "../../tasks/task-list/task-list.component";
import {environment} from "../../../environments/environment";

@Component({
  selector: 'app-home-view',
  standalone: true,
  imports: [
    MainLayoutComponent,
    MatCard,
    MatCardHeader,
    MatCardContent,
    TaskListComponent,
    MatCardFooter,
    MatCardActions,
    MatCardTitle,
    MatPaginator
  ],
  templateUrl: './home-view.component.html',
  styleUrl: './home-view.component.scss'
})
export class HomeViewComponent {
  protected pageSizes: number[] = [10, 20, 30, 50, 100];
  protected pageSize = model(10);
  protected tasks: WritableSignal<Task[]> = signal([]);
  private nextPageToken?: string;

  constructor(private readonly taskService: TaskService) {
    this.listTasks();
    effect(() => {
      this.nextPageToken = undefined;
      this.listTasks();
    });
  }

  handlePageEvent(e: PageEvent) {
    this.pageSize.set(e.pageSize);
  }

  private listTasks() {
    this.taskService.list(this.pageSize(), this.nextPageToken).subscribe((response: ListTasksResponse) => {
      this.tasks.set(response.tasks);
      this.nextPageToken = response.nextPageToken;
    })
  }

  protected readonly environment = environment;
}
