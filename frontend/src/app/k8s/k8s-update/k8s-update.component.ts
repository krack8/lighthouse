import { Component, OnInit, ViewChild } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { K8sNamespacesService } from '@k8s/k8s-namespaces/k8s-namespaces.service';
import { ToastrService } from '@sdk-ui/ui';
import { MatTabChangeEvent, MatTabGroup } from '@angular/material/tabs';
import { saveAs } from 'file-saver';
import { dump as toYaml, load as fromYaml } from 'js-yaml';
import 'ace-builds';
import 'ace-builds/src-noconflict/mode-json';
import 'ace-builds/src-noconflict/mode-yaml';
import { Observable } from 'rxjs';
import { K8sPersistentVolumeService } from '@k8s/k8s-persistent-volume/k8s-persistent-volume.service';
import { K8sClusterRoleBindingService } from '@k8s/k8s-cluster-role-binding/k8s-cluster-role-binding.service';
import { K8sClusterRoleService } from '@k8s/k8s-cluster-role/k8s-cluster-role.service';
import { K8sStorageClassService } from '@k8s/k8s-storage-class/k8s-storage-class.service';
import { K8sClusterCustomResourcesService } from '@k8s/k8s-cluster-custom-resources/k8s-cluster-custom-resources.service';
import { K8sNodesService } from '@k8s/k8s-nodes/k8s-nodes.service';
import { K8sService } from '@k8s/k8s.service';

@Component({
  selector: 'kc-k8s-update',
  templateUrl: './k8s-update.component.html',
  styleUrls: ['./k8s-update.component.scss']
})
export class K8sUpdateComponent implements OnInit {
  @ViewChild('tabGroup', { static: false }) tab: MatTabGroup;
  isSubmitting = false;
  payload: any;
  inputData = '';
  preInputData;
  isEditMode: boolean;
  applyManifestFor: string;
  queryParams: any = {};
  selectedTabIndex: number = 0;
  errorIndex: number;
  tabChange: boolean = true;
  acceptedFileTypes: string = '.yaml, .yml';
  originalInputData;

  constructor(
    private namespaceService: K8sNamespacesService,
    private pvService: K8sPersistentVolumeService,
    private clusterRoleService: K8sClusterRoleService,
    private clusterRoleBindingService: K8sClusterRoleBindingService,
    private storageClassService: K8sStorageClassService,
    private customResourceService: K8sClusterCustomResourcesService,
    private nodeService: K8sNodesService,
    private k8sService: K8sService,
    public toastrService: ToastrService,
    private dialog: MatDialogRef<K8sUpdateComponent>
  ) {}

  ngOnInit(): void {
    if (this.isEditMode) {
      try {
        if (!this.preInputData.spec) {
          delete this.preInputData.spec;
        }
        if (!this.preInputData.metadata.annotations) {
          delete this.preInputData.metadata.annotations;
        }
        if (!this.preInputData.metadata.labels) {
          delete this.preInputData.metadata.labels;
        }
        if (!this.preInputData.selfLink) {
          delete this.preInputData.metadata.selfLink;
        }
        this.originalInputData = toYaml(this.preInputData);

        this.inputData = this.originalInputData;
      } catch (e) {
        console.log(e);
      }
    }
  }

  onFileSelected(event: any) {
    const file: File = event.target.files[0];
    const reader: FileReader = new FileReader();
    reader.onload = (e: any) => {
      const fileData: any = e.target.result;
      this.inputData = fileData;
    };
    reader.readAsText(file);
  }

  exportFile() {
    if (this.tab.selectedIndex == 0) {
      const yamlData = this.inputData;
      const blob = new Blob([yamlData], { type: 'text/yaml;charset=utf-8' });
      saveAs(blob, 'data.yaml');
    }

    if (this.tab.selectedIndex == 1) {
      const jsonFile = this.inputData;
      const blob = new Blob([jsonFile], { type: 'application/json;charset=utf-8' });
      saveAs(blob, 'data.json');
    }
  }

  reset() {
    if (this.selectedTabIndex === 1) {
      this.inputData = this.toRawJSON(this.preInputData);
    } else {
      this.inputData = toYaml(this.preInputData);
    }
  }

  update(): void {
    this.isSubmitting = true;

    try {
      JSON.parse(this.toRawJSON(fromYaml(this.inputData)));
    } catch (e) {
      this.toastrService.error(e.reason, 'Error');
      this.isSubmitting = false;
      return;
    }

    const payload = JSON.parse(this.toRawJSON(fromYaml(this.inputData)));

    let apiObservable: Observable<any>;

    switch (this.applyManifestFor) {
      case 'all':
        if (payload?.kind) {
          if (payload.kind === 'Node' || payload.kind === 'node') {
            this.toastrService.warn('Node creation via manifest is not allowed.');
            this.isSubmitting = false;
            return;
          }
          apiObservable = this.k8sService.applyManifest(payload, this.applyManifestFor);
        } else {
          this.toastrService.error('Kind needs to be specified properly');
          this.isSubmitting = false;
          return;
        }
        break;
      case 'persistent-volume':
        apiObservable = this.pvService.applyManifest(payload, this.applyManifestFor);
        break;
      case 'custom-resources':
        apiObservable = this.customResourceService.applyManifest(payload, this.applyManifestFor, this.queryParams);
        break;
      case 'crd':
        apiObservable = this.customResourceService.applyManifest(payload, this.applyManifestFor);
        break;
      case 'cluster-role':
        apiObservable = this.clusterRoleService.applyManifest(payload, this.applyManifestFor);
        break;
      case 'cluster-role-binding':
        apiObservable = this.clusterRoleBindingService.applyManifest(payload, this.applyManifestFor);
        break;
      case 'storage-class':
        apiObservable = this.storageClassService.applyManifest(payload, this.applyManifestFor);
        break;
      default:
        if (!payload?.metadata?.namespace && this.applyManifestFor != 'namespace') {
          this.toastrService.error('Namespace needs to be specified');
          this.isSubmitting = false;
          return;
        }
        apiObservable = this.namespaceService.applyManifest(payload, this.applyManifestFor);
    }

    apiObservable.subscribe(
      res => {
        if (res.status === 'success') {
          this.isSubmitting = false;
          this.toastrService.success('Manifest Applied Successfully', 'Success');
          this.dialog.close(res);
        } else {
          this.toastrService.error(res.msg, 'Error');
        }
      },
      err => {
        this.toastrService.error(err.error.message);
        this.isSubmitting = false;
      }
    );
  }

  cancel(): void {
    this.inputData = '';
  }

  areMultipleNamespacesSelected(): boolean {
    return true;
  }

  isClicked() {
    // this function is used to check if the changed type function is called by clicking or not as it is also called when the tab Index is changedin the function
    this.tabChange = true;
  }

  changeType(type: MatTabChangeEvent): void {
    if (this.selectedTabIndex === 0) {
      this.acceptedFileTypes = '.yaml, .yml';
    } else {
      this.acceptedFileTypes = '.json';
    }

    if (type.index === 0 && this.inputData && this.tabChange) {
      try {
        this.inputData = toYaml(JSON.parse(this.inputData));
      } catch (e) {
        this.toastrService.error(e.reason, 'Could not convert JSON to YAML');
        this.tab.selectedIndex = 1;
        this.tabChange = false;
      }
    }

    if (type.index === 1 && this.inputData && this.tabChange) {
      try {
        this.inputData = this.toRawJSON(fromYaml(this.inputData));
      } catch (e) {
        this.toastrService.error(e.reason, 'Could not convert YAML to JSON');
        this.tab.selectedIndex = 0;
        this.tabChange = false;
      }
    }
  }

  toRawJSON(object: {}): string {
    return JSON.stringify(object, null, '\t');
  }
}
