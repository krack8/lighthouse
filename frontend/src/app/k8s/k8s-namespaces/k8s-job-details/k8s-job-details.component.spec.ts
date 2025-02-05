import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { K8sJobDetailsComponent } from './k8s-job-details.component';

describe('K8sJobDetailsComponent', () => {
  let component: K8sJobDetailsComponent;
  let fixture: ComponentFixture<K8sJobDetailsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [K8sJobDetailsComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(K8sJobDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
