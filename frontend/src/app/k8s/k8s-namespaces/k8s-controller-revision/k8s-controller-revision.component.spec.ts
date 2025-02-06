import { ComponentFixture, TestBed } from '@angular/core/testing';

import { K8sControllerRevisionComponent } from './k8s-controller-revision.component';

describe('K8sControllerRevisionComponent', () => {
  let component: K8sControllerRevisionComponent;
  let fixture: ComponentFixture<K8sControllerRevisionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [K8sControllerRevisionComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(K8sControllerRevisionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
