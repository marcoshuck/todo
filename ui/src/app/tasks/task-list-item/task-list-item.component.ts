import {Component, Input, OnInit, signal} from '@angular/core';
import {MatCheckbox} from "@angular/material/checkbox";
import {MatGridList, MatGridTile} from "@angular/material/grid-list";
import {Task} from "../../../api/tasks/v1/tasks_pb";
import {Router} from "@angular/router";
import {IconButtonStarComponent} from "../../components/icon-button-star/icon-button-star.component";

@Component({
  selector: 'app-task-list-item',
  standalone: true,
  imports: [
    MatGridList,
    MatGridTile,
    MatCheckbox,
    IconButtonStarComponent
  ],
  templateUrl: './task-list-item.component.html',
  styleUrl: './task-list-item.component.scss'
})
export class TaskListItemComponent implements OnInit {
  @Input({required: true})
  public task!: Task;

  protected completed = signal(false);
  protected starred = signal(false);

  constructor(private router: Router) {
  }

  ngOnInit(): void {
    this.completed.set(!!this.task.completedAt);
  }

  star() {
    this.starred.set(!this.starred());
  }

  async seeMore() {
    await this.router.navigateByUrl(`/tasks/${this.task.id}`);
  }
}
