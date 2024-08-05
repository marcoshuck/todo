import {ComponentFixture, TestBed} from '@angular/core/testing';

import {IconButtonStarComponent} from './icon-button-star.component';

describe('IconButtonStarComponent', () => {
  let component: IconButtonStarComponent;
  let fixture: ComponentFixture<IconButtonStarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [IconButtonStarComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(IconButtonStarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
