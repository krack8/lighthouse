import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIntroComponent } from './cluster-intro.component';

describe('ClusterIntroComponent', () => {
  let component: ClusterIntroComponent;
  let fixture: ComponentFixture<ClusterIntroComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterIntroComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterIntroComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
