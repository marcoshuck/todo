import {ChangeDetectionStrategy, Component, signal} from '@angular/core';
import {MatToolbar} from "@angular/material/toolbar";
import {MatIcon} from "@angular/material/icon";
import {MatIconButton} from "@angular/material/button";
import {IconButtonStarComponent} from "../../components/icon-button-star/icon-button-star.component";
import {ActivatedRoute} from "@angular/router";
import {TaskService} from "../task.service";
import {Task} from '../../../api/tasks/v1/tasks_pb';
import {MatList, MatListItem} from "@angular/material/list";
import {MatFormField, MatHint, MatLabel} from "@angular/material/form-field";
import {MatDatepicker, MatDatepickerInput, MatDatepickerToggle} from "@angular/material/datepicker";
import {MatInput} from "@angular/material/input";
import {provideNativeDateAdapter} from "@angular/material/core";
import {Location} from '@angular/common';
import {MatProgressSpinner} from "@angular/material/progress-spinner";

@Component({
  selector: 'app-task-details',
  standalone: true,
  imports: [
    MatToolbar,
    MatIcon,
    MatIconButton,
    IconButtonStarComponent,
    MatList,
    MatListItem,
    MatFormField,
    MatDatepickerToggle,
    MatDatepicker,
    MatDatepickerInput,
    MatInput,
    MatLabel,
    MatHint,
    MatProgressSpinner,
  ],
  templateUrl: './task-details.component.html',
  styleUrl: './task-details.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
  providers: [provideNativeDateAdapter()],
})
export class TaskDetailsComponent {
  deadline?: Date = new Date();
  protected task = signal<Task>({
    id: 0n,
    title: "",
    description: "",
  })

  constructor(
    private taskService: TaskService,
    private activatedRoute: ActivatedRoute,
    private location: Location,
  ) {
    try {
      const id = this.activatedRoute.snapshot.paramMap.get('id');
      if (!id) {
        console.error('Failed to process id')
        this.goBack();
      }
      this.taskService.get(BigInt(id!)).subscribe((task => {
        this.task.set(task);
      }));
    } catch (err: unknown) {
      console.error(err);
      this.goBack();
    }
  }


  protected goBack() {
    this.location.back();
  }
}
