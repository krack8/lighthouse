import { ComponentFixture, TestBed } from '@angular/core/testing';

import { K8sEndpointsComponent } from './k8s-endpoints.component';

describe('K8sEndpointsComponent', () => {
  let component: K8sEndpointsComponent;
  let fixture: ComponentFixture<K8sEndpointsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [K8sEndpointsComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(K8sEndpointsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
