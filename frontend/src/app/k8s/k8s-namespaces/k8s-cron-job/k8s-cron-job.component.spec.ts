import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { K8sCronJobComponent } from './k8s-cron-job.component';

describe('K8sCronJobComponent', () => {
  let component: K8sCronJobComponent;
  let fixture: ComponentFixture<K8sCronJobComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [K8sCronJobComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(K8sCronJobComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
