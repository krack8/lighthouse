import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterUpdateFormComponent } from './cluster-update-form.component';

describe('ClusterUpdateFormComponent', () => {
  let component: ClusterUpdateFormComponent;
  let fixture: ComponentFixture<ClusterUpdateFormComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ClusterUpdateFormComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ClusterUpdateFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
