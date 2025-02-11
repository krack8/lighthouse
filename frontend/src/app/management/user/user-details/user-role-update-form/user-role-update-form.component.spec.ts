import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserRoleUpdateFormComponent } from './user-role-update-form.component';

describe('UserRoleUpdateFormComponent', () => {
  let component: UserRoleUpdateFormComponent;
  let fixture: ComponentFixture<UserRoleUpdateFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [UserRoleUpdateFormComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(UserRoleUpdateFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
