import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { RoleDeleteConformationComponent } from './role-delete-conformation.component';

describe('RoleDeleteConformationComponent', () => {
  let component: RoleDeleteConformationComponent;
  let fixture: ComponentFixture<RoleDeleteConformationComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [RoleDeleteConformationComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(RoleDeleteConformationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
