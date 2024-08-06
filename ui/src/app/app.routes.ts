import {Routes} from '@angular/router';
import {HomeViewComponent} from "./views/home-view/home-view.component";
import {TaskViewComponent} from "./views/task-view/task-view.component";

export const routes: Routes = [
  {
    title: 'Home',
    path: '',
    component: HomeViewComponent,
  },
  {
    title: 'Task',
    path: 'tasks/:id',
    component: TaskViewComponent,
  }
];
