import {Component, signal} from '@angular/core';
import {MainLayoutComponent} from "../../layouts/main-layout/main-layout.component";
import {TaskService} from "../../tasks/task.service";
import {ActivatedRoute} from "@angular/router";
import {Location} from "@angular/common";
import {Task} from "../../../api/tasks/v1/tasks_pb";
import {TaskDetailsComponent} from "../../tasks/task-details/task-details.component";
import {IconButtonStarComponent} from "../../components/icon-button-star/icon-button-star.component";
import {MatIcon} from "@angular/material/icon";
import {MatIconButton} from "@angular/material/button";
import {MatToolbar} from "@angular/material/toolbar";

@Component({
  selector: 'app-task-view',
  standalone: true,
  imports: [
    MainLayoutComponent,
    TaskDetailsComponent,
    IconButtonStarComponent,
    MatIcon,
    MatIconButton,
    MatToolbar
  ],
  templateUrl: './task-view.component.html',
  styleUrl: './task-view.component.scss'
})
export class TaskViewComponent {
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
    const id = this.activatedRoute.snapshot.paramMap.get('id');
    if (!id) {
      console.error('Failed to process id')
      this.goBack();
    }
    this.taskService.get(BigInt(id!)).subscribe((task => {
      this.task.set(task);
    }));
  }

  catch(err: unknown) {
    console.error(err);
    this.goBack();
  }

  protected goBack() {
    this.location.back();
  }
}
