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
import {MatListItem, MatListOption, MatSelectionList} from "@angular/material/list";
import {TaskService} from "../../tasks/task.service";
import {ListTasksResponse, Task} from "../../../api/tasks/v1/tasks_pb";
import {MatFormField, MatLabel} from "@angular/material/form-field";
import {MatOption, MatSelect} from "@angular/material/select";
import {FormsModule} from "@angular/forms";
import {TaskListItemComponent} from "../../tasks/task-list-item/task-list-item.component";
import {MatPaginator, PageEvent} from "@angular/material/paginator";
import {TaskListComponent} from "../../tasks/task-list/task-list.component";

@Component({
  selector: 'app-home-view',
  standalone: true,
  imports: [
    MainLayoutComponent,
    MatCard,
    MatCardTitle,
    MatCardHeader,
    MatCardContent,
    MatCardFooter,
    MatSelectionList,
    MatLabel,
    MatListItem,
    MatListOption,
    MatFormField,
    MatSelect,
    MatOption,
    FormsModule,
    MatCardActions,
    TaskListItemComponent,
    MatPaginator,
    TaskListComponent
  ],
  templateUrl: './home-view.component.html',
  styleUrl: './home-view.component.scss'
})
export class HomeViewComponent {
  protected pageSizes: number[] = [1, 10, 20, 30, 50];
  protected pageSize = model(1);
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
}
