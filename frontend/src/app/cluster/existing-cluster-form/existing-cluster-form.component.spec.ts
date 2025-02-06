import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ExistingClusterFormComponent } from './existing-cluster-form.component';

describe('ExistingClusterFormComponent', () => {
  let component: ExistingClusterFormComponent;
  let fixture: ComponentFixture<ExistingClusterFormComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ExistingClusterFormComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ExistingClusterFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
