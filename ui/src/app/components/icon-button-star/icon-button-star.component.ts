import {Component, effect, model} from '@angular/core';
import {MatIcon} from "@angular/material/icon";
import {MatIconButton} from "@angular/material/button";

@Component({
  selector: 'app-icon-button-star',
  standalone: true,
  imports: [
    MatIcon,
    MatIconButton
  ],
  templateUrl: './icon-button-star.component.html',
  styleUrl: './icon-button-star.component.scss'
})
export class IconButtonStarComponent {
  public starred = model<boolean>(false);
  protected icon = 'star_border';

  constructor() {
    effect(() => {
      this.icon = this.starred() ? 'star' : 'star_border';
    });
  }

  star() {
    this.starred.set(!this.starred())
  }
}
