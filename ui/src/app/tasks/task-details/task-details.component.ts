import {ChangeDetectionStrategy, Component, input} from '@angular/core';
import {Task} from '../../../api/tasks/v1/tasks_pb';
import {MatList, MatListItem} from "@angular/material/list";
import {provideNativeDateAdapter} from "@angular/material/core";
import {MatProgressSpinner} from "@angular/material/progress-spinner";

@Component({
  selector: 'app-task-details',
  standalone: true,
  imports: [
    MatProgressSpinner,
    MatList,
    MatListItem
  ],
  templateUrl: './task-details.component.html',
  styleUrl: './task-details.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
  providers: [provideNativeDateAdapter()],
})
export class TaskDetailsComponent {
  public task = input.required<Task>();
}
