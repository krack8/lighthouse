import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OnboardClusterPrerequisiteComponent } from './onboard-cluster-prerequisite.component';

describe('OnboardClusterPrerequisiteComponent', () => {
  let component: OnboardClusterPrerequisiteComponent;
  let fixture: ComponentFixture<OnboardClusterPrerequisiteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OnboardClusterPrerequisiteComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(OnboardClusterPrerequisiteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
