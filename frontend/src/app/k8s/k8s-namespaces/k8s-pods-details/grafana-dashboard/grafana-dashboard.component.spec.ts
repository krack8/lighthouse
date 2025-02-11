import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GrafanaDashboardComponent } from './grafana-dashboard.component';

describe('GrafanaDashboardComponent', () => {
  let component: GrafanaDashboardComponent;
  let fixture: ComponentFixture<GrafanaDashboardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [GrafanaDashboardComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(GrafanaDashboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
