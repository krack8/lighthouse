import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NodeTaintDialogComponent } from './node-taint-dialog.component';

describe('NodeTaintDialogComponent', () => {
  let component: NodeTaintDialogComponent;
  let fixture: ComponentFixture<NodeTaintDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [NodeTaintDialogComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(NodeTaintDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
