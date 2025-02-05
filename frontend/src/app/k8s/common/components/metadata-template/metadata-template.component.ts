import { CommonModule } from '@angular/common';
import { Component, Input, OnInit } from '@angular/core';
import { MatExpansionModule } from '@angular/material/expansion';
import { AgoPipe } from '../../../../../../projects/shared-ui/src/pipes/ago.pipe';
import { ExpansionDataViewerTemplateComponent } from '../expansion-data-viewer-template/expansion-data-viewer-template.component';

@Component({
  selector: 'kc-metadata-template',
  templateUrl: './metadata-template.component.html',
  styleUrls: ['./metadata-template.component.scss'],
  standalone: true,
  imports: [CommonModule, MatExpansionModule, AgoPipe, ExpansionDataViewerTemplateComponent]
})
export class MetadataTemplateComponent implements OnInit {
  @Input() data: any;
  panelOpenState: boolean;
  constructor() {}

  ngOnInit(): void {}
}
