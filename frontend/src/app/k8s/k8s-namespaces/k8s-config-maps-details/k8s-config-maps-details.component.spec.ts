import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { K8sConfigMapsDetailsComponent } from './k8s-config-maps-details.component';

describe('K8sConfigMapsDetailsComponent', () => {
  let component: K8sConfigMapsDetailsComponent;
  let fixture: ComponentFixture<K8sConfigMapsDetailsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [K8sConfigMapsDetailsComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(K8sConfigMapsDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
