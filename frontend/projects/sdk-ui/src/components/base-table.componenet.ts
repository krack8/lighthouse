import { ChangeDetectorRef } from '@angular/core';

export class BaseTableComponenet {
  initialSelection = [];
  allowMultiSelect = true;
  selection: any;
  tableValue = [];

  constructor(protected cdr: ChangeDetectorRef) {}

  /** Whether the number of selected elements matches the total number of rows. */
  protected isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.tableValue.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  protected masterToggle() {
    this.isAllSelected() ? this.selection.clear() : this.tableValue.forEach(row => this.selection.select(row));
  }
}
