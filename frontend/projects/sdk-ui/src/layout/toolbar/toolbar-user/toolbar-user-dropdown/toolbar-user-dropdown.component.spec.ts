import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ToolbarUserDropdownComponent } from './toolbar-user-dropdown.component';

describe('ToolbarUserDropdownComponent', () => {
  let component: ToolbarUserDropdownComponent;
  let fixture: ComponentFixture<ToolbarUserDropdownComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ToolbarUserDropdownComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ToolbarUserDropdownComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
