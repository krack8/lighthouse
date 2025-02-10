import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import icClose from '@iconify/icons-ic/close';

@Component({
  selector: 'kc-json-data-viewer',
  templateUrl: './json-data-viewer.component.html',
  styleUrls: ['./json-data-viewer.component.scss']
})
export class JsonDataViewerComponent implements OnInit {
  constructor(
    public dialogRef: MatDialogRef<JsonDataViewerComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {}

  icClose = icClose;

  ngOnInit(): void {}

  close() {
    this.dialogRef.close();
  }

  isInt(value: string): boolean {
    const parsedValue = parseInt(value, 10);
    return !isNaN(parsedValue) && String(parsedValue) === value;
  }

  isObject(value: any): boolean {
    return typeof value === 'object' && value !== null;
  }
}
