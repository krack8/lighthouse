import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { K8sServiceAccountsDetailsComponent } from './k8s-service-accounts-details.component';

describe('K8sServiceAccountsDetailsComponent', () => {
  let component: K8sServiceAccountsDetailsComponent;
  let fixture: ComponentFixture<K8sServiceAccountsDetailsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [K8sServiceAccountsDetailsComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(K8sServiceAccountsDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
