import { Component, Input, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'kc-json-data-viewer-template',
  templateUrl: './json-data-viewer-template.component.html',
  styleUrls: ['./json-data-viewer-template.component.scss'],
  standalone: true,
  imports: [CommonModule]
})
export class JsonDataViewerTemplateComponent implements OnInit {
  @Input() data: any;

  constructor() {}

  ngOnInit(): void {}

  isInt(value: string): boolean {
    const parsedValue = parseInt(value, 10);
    return !isNaN(parsedValue) && String(parsedValue) === value;
  }

  isObject(value: any): boolean {
    return typeof value === 'object' && value !== null;
  }
}
