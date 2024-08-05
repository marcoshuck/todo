import {Routes} from '@angular/router';
import {HomeViewComponent} from "./views/home-view/home-view.component";

export const routes: Routes = [
  {
    title: 'Home',
    path: '',
    component: HomeViewComponent,
  }
];
