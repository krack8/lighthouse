import { ComponentFixture, TestBed } from '@angular/core/testing';

import { K8sEndpointSliceComponent } from './k8s-endpoint-slice.component';

describe('K8sEndpointSliceComponent', () => {
  let component: K8sEndpointSliceComponent;
  let fixture: ComponentFixture<K8sEndpointSliceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [K8sEndpointSliceComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(K8sEndpointSliceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
