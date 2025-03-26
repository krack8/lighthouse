import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterRenameDialogComponent } from './cluster-rename-dialog.component';

describe('ClusterRenameDialogComponent', () => {
  let component: ClusterRenameDialogComponent;
  let fixture: ComponentFixture<ClusterRenameDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ClusterRenameDialogComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClusterRenameDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
