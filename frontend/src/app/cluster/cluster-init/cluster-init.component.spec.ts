import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterInitComponent } from './cluster-init.component';

describe('ClusterInitComponent', () => {
  let component: ClusterInitComponent;
  let fixture: ComponentFixture<ClusterInitComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ ClusterInitComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClusterInitComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
