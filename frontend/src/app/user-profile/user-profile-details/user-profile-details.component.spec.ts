import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { UserProfileDetailsComponent } from './user-profile-details.component';

describe('UserProfileDetailsComponent', () => {
  let component: UserProfileDetailsComponent;
  let fixture: ComponentFixture<UserProfileDetailsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [UserProfileDetailsComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UserProfileDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
