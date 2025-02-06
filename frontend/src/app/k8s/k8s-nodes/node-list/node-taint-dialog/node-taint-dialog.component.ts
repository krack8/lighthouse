import { Component, Inject, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialog, MatDialogRef } from '@angular/material/dialog';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icAdd from '@iconify/icons-ic/add';
import { ToastrService } from '@sdk-ui/ui';
import { K8sNodesService } from '@k8s/k8s-nodes/k8s-nodes.service';
import { ConfirmDialogComponent } from '@shared-ui/ui';

@Component({
  selector: 'kc-node-taint-dialog',
  templateUrl: './node-taint-dialog.component.html',
  styleUrls: ['./node-taint-dialog.component.scss']
})
export class NodeTaintDialogComponent implements OnInit {
  taintList: any = [];
  untaintList: any = {
    keys: []
  };
  addTaint: boolean = false;
  icDelete = icDelete;
  icAdd = icAdd;
  submittingForm: boolean = false;
  deleteAction: boolean = false;
  nodeName: string;
  form: FormGroup;
  anyActionPerfomed: boolean = false; // if any action is performed the node list is re-loaded otherwise not

  constructor(
    @Inject(MAT_DIALOG_DATA) private data,
    public dialogRef: MatDialogRef<NodeTaintDialogComponent>,
    private fb: FormBuilder,
    private toastr: ToastrService,
    public dialog: MatDialog,
    private nodeService: K8sNodesService
  ) {}

  ngOnInit(): void {
    this.form = this.fb.group({
      taints: this.fb.array([])
    });

    this.nodeName = this.data?.metadata?.name;
    if (this.data?.spec?.taints) {
      this.taintList = this.data?.spec?.taints;
    }

    this.addNewTaint();
  }

  toogleAddTaint() {
    this.addTaint = !this.addTaint;
  }

  get taints() {
    return this.form.get('taints') as FormArray;
  }

  addNewTaint() {
    const textboxGroup = this.fb.group({
      key: ['', [Validators.required]],
      value: ['', [Validators.required]],
      effect: ['', [Validators.required]]
    });

    this.taints.push(textboxGroup);
  }

  toggleCheckbox(checked: Event, taint: any) {
    if (checked) {
      this.untaintList.keys.push(taint.key);
    }
    if (!checked) {
      this.untaintList.keys = this.untaintList.keys.filter(k => k !== taint.key);
    }
  }

  removeTaint(index: number) {
    this.taints.removeAt(index);
  }

  closeDialog() {
    this.dialogRef.close(this.anyActionPerfomed);
  }

  saveTaints() {
    this.submittingForm = true;
    const body = {
      taint: this.form.value.taints
    };

    const dialogRef = this.dialog.open(ConfirmDialogComponent, {
      minHeight: '270px',
      width: '720px',
      disableClose: true,
      data: `Are you sure you want the taint(s) to the node, ${this.nodeName} ?`
    });

    dialogRef.afterClosed().subscribe(res => {
      if (res) {
        this.nodeService.taintNode(this.nodeName, body).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Node tainted successfully');
              this.taintList = [...this.taintList, ...body.taint];
              this.submittingForm = false;
              this.taints.clear();
              this.addNewTaint();
              this.addTaint = false;
              this.anyActionPerfomed = true;
            } else {
              this.toastr.error('Node taint was unsuccessful', res.message);
              this.submittingForm = false;
            }
          },
          err => {
            this.toastr.error('Node taint was unsuccessful');
            this.submittingForm = false;
          }
        );
      }
    });
  }

  untaintNode(untaintAll?: boolean) {
    let body = this.untaintList;
    let warnMsg = 'Are you sure you want to untaint the selected taints?';
    if (untaintAll) {
      body = {
        keys: []
      };
      warnMsg = 'Are you sure you want to untaint all taints?';
    }
    const dialogRef = this.dialog.open(ConfirmDialogComponent, {
      minHeight: '270px',
      width: '720px',
      disableClose: true,
      data: warnMsg
    });

    dialogRef.afterClosed().subscribe(res => {
      if (res) {
        this.deleteAction = true;
        this.nodeService.untaintNode(this.nodeName, body).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Node untainted successfully');
              if (untaintAll) {
                this.taintList = [];
              } else {
                this.taintList = this.taintList.filter(t => !body.keys.includes(t.key));
              }
              this.untaintList.keys = [];
              this.deleteAction = false;
              this.anyActionPerfomed = true;
            } else {
              this.toastr.error('Node untaint was unsuccessful', res.message);
              this.deleteAction = false;
            }
          },
          err => {
            this.toastr.error('Node untaint was unsuccessful');
            this.deleteAction = false;
          }
        );
      }
    });
  }
}
