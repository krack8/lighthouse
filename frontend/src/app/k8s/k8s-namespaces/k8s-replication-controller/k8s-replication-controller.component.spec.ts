import { ComponentFixture, TestBed } from '@angular/core/testing';

import { K8sReplicationControllerComponent } from './k8s-replication-controller.component';

describe('K8sReplicationControllerComponent', () => {
  let component: K8sReplicationControllerComponent;
  let fixture: ComponentFixture<K8sReplicationControllerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [K8sReplicationControllerComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(K8sReplicationControllerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
