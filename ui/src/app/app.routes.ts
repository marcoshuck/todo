import {Routes} from '@angular/router';
import {HomeViewComponent} from "./views/home-view/home-view.component";
import {TaskDetailsComponent} from "./tasks/task-details/task-details.component";

export const routes: Routes = [
  {
    title: 'Home',
    path: '',
    component: HomeViewComponent,
  },
  {
    title: 'Task',
    path: 'tasks/:id',
    component: TaskDetailsComponent,
  }
];
