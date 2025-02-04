import { ComponentFixture, TestBed } from '@angular/core/testing';

import { K8sPodDisruptionBudgetsComponent } from './k8s-pod-disruption-budgets.component';

describe('K8sPodDisruptionBudgetsComponent', () => {
  let component: K8sPodDisruptionBudgetsComponent;
  let fixture: ComponentFixture<K8sPodDisruptionBudgetsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [K8sPodDisruptionBudgetsComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(K8sPodDisruptionBudgetsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
