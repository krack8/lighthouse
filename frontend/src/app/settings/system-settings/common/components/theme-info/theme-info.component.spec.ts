import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ThemeInfoComponent } from './theme-info.component';

describe('ThemeInfoComponent', () => {
  let component: ThemeInfoComponent;
  let fixture: ComponentFixture<ThemeInfoComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ThemeInfoComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ThemeInfoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
