import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MetadataTemplateComponent } from './metadata-template.component';

describe('MetadataTemplateComponent', () => {
  let component: MetadataTemplateComponent;
  let fixture: ComponentFixture<MetadataTemplateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [MetadataTemplateComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(MetadataTemplateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
