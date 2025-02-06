import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Router } from '@angular/router';
import icClose from '@iconify/icons-ic/twotone-close';
import icDownload from '@iconify/icons-ic/twotone-download';

@Component({
  selector: 'kc-cluster-release-note-dialog',
  templateUrl: './cluster-release-note-dialog.component.html',
  styleUrls: ['./cluster-release-note-dialog.component.scss']
})
export class ClusterReleaseNoteDialogComponent implements OnInit {
  icClose = icClose;
  icDownload = icDownload;
  warningStep: boolean = false;
  releaseNote: any;
  releaseSuccess: boolean = null;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any,
    public dialogRef: MatDialogRef<ClusterReleaseNoteDialogComponent>,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.releaseNote = this.data?.newReleaseNote;
  }

  toggleWarning() {
    this.warningStep = !this.warningStep;
  }

  downloadLogs() {
    const content = this.data.errorMessage;
    const blob = new Blob([content], { type: 'text/plain' });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'error-logs.txt';
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    window.URL.revokeObjectURL(url);
  }

  redirectLink() {
    this.router.navigate(['ticket']);
    this.dialogRef.close();
  }
}
