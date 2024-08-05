import {Component, input} from '@angular/core';
import {TaskListItemComponent} from "../task-list-item/task-list-item.component";
import {Task} from '../../../api/tasks/v1/tasks_pb';

@Component({
  selector: 'app-task-list',
  standalone: true,
  imports: [
    TaskListItemComponent
  ],
  templateUrl: './task-list.component.html',
  styleUrl: './task-list.component.scss'
})
export class TaskListComponent {
  public tasks = input.required<Task[]>();
}
