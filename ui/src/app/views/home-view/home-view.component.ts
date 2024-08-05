import { Component } from '@angular/core';
import {MainLayoutComponent} from "../../layouts/main-layout/main-layout.component";

@Component({
  selector: 'app-home-view',
  standalone: true,
  imports: [
    MainLayoutComponent
  ],
  templateUrl: './home-view.component.html',
  styleUrl: './home-view.component.scss'
})
export class HomeViewComponent {

}
