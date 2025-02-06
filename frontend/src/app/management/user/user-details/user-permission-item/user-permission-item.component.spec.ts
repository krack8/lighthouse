import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserPermissionItemComponent } from './user-permission-item.component';

describe('UserPermissionItemComponent', () => {
  let component: UserPermissionItemComponent;
  let fixture: ComponentFixture<UserPermissionItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [UserPermissionItemComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(UserPermissionItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
