import { CommonModule } from '@angular/common';
import { Component, Input, OnInit } from '@angular/core';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatIconModule } from '@angular/material/icon';
import icMoveUp from '@iconify/icons-ic/twotone-arrow-drop-up';
import { IconModule } from '@visurel/iconify-angular';
import { JsonDataViewerTemplateComponent } from '../json-data-viewer-template/json-data-viewer-template.component';
@Component({
  selector: 'kc-expansion-data-viewer-template',
  templateUrl: './expansion-data-viewer-template.component.html',
  styleUrls: ['./expansion-data-viewer-template.component.scss'],
  standalone: true,
  imports: [CommonModule, MatExpansionModule, MatIconModule, IconModule, JsonDataViewerTemplateComponent]
})
export class ExpansionDataViewerTemplateComponent implements OnInit {
  @Input() data: any;
  @Input() label: any;
  panelOpenState: boolean;
  icMoveUp = icMoveUp;
  isOpen = false;
  noData: boolean = false;
  constructor() {}

  ngOnInit(): void {
    if (!this.data || this.data.length === 0 || Object.keys(this.data).length === 0) {
      this.noData = true;
    }
  }

  togglePanel() {
    this.isOpen = !this.isOpen;
  }
}
