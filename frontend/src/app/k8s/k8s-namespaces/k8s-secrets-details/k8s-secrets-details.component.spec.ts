import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { K8sSecretsDetailsComponent } from './k8s-secrets-details.component';

describe('K8sSecretsDetailsComponent', () => {
  let component: K8sSecretsDetailsComponent;
  let fixture: ComponentFixture<K8sSecretsDetailsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [K8sSecretsDetailsComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(K8sSecretsDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
