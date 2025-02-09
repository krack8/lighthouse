import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { K8sStatefulSetsComponent } from './k8s-stateful-sets.component';

describe('K8sStatefulSetsComponent', () => {
  let component: K8sStatefulSetsComponent;
  let fixture: ComponentFixture<K8sStatefulSetsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [K8sStatefulSetsComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(K8sStatefulSetsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
