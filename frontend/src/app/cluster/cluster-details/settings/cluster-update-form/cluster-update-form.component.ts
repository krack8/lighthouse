import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import icInfo from '@iconify/icons-ic/info';
import icVisibility from '@iconify/icons-ic/visibility';
import icVisibilityOff from '@iconify/icons-ic/visibility-off';
import { ClusterService } from '@cluster/cluster.service';
import { ToastrService } from '@sdk-ui/ui';
import { DnsValidator } from '@shared-ui/validators';

@Component({
  selector: 'kc-cluster-update-form',
  templateUrl: './cluster-update-form.component.html',
  styleUrls: ['./cluster-update-form.component.scss']
})
export class ClusterUpdateFormComponent implements OnInit {
  @Input() clusterId: string;
  @Output() closeFormEvent = new EventEmitter<void>();

  icInfo = icInfo;
  icVisibilityOff = icVisibilityOff;
  icVisibility = icVisibility;

  isLoading: boolean = true;
  clusterDetails: any = null;

  clusterConfigForm!: FormGroup;
  isSubmitting: boolean = false;

  // Additional Information Features
  enableVolumeSnapshot = false;
  enablePrometheus = false;
  enableGrafana = false;
  enableLoki = false;
  enableArgo = false;
  enableKiali = false;

  // Passwords view
  showPrometheusPassword = false;
  showGrafanaPassword = false;
  showLokiPassword = false;
  showArgoPassword = false;
  showKialiToken = false;

  serviceMeshToolType = 'ISTIO';
  serviceMeshSslEnabled = false;

  constructor(
    private fb: FormBuilder,
    private _clusterService: ClusterService,
    private toastrService: ToastrService
  ) {}

  ngOnInit(): void {
    this.clusterConfigForm = this.fb.group({
      clusterIssuerName: [''],
      snapshotClassName: [''],
      snapshotClassRWMName: [''],

      prometheusUrl: [''],
      prometheusUsername: [''],
      prometheusPassword: [''],

      grafanaUrl: [''],
      grafanaUsername: [''],
      grafanaPassword: [''],

      lokiUrl: [''],
      lokiUsername: [''],
      lokiPassword: [''],
      lokiOrgId: [''],

      argocdURL: [''],
      argocdUsername: [''],
      argocdPassword: [''],
      argocdPort: [''],

      serviceMeshEnabled: [false],

      serviceMeshGatewayWildcardDomain: [''],
      istioGatewayName: [''],
      serviceMeshGatewayWildcardDomainTls: [''],

      kialiURL: [''],
      kialiToken: ['']
    });
    this.getClusterDetails();
  }

  // Additional Config Start
  onAddiVolumeSnapshotChange(checked: boolean): void {
    if (checked) {
      this.clusterConfigForm.get('snapshotClassName').setValidators([Validators.required, Validators.pattern(/^[a-zA-Z0-9-]+$/)]);
      this.clusterConfigForm.get('snapshotClassName').updateValueAndValidity();
      this.clusterConfigForm.get('snapshotClassRWMName').setValidators([Validators.required, Validators.pattern(/^[a-zA-Z0-9-]+$/)]);
      this.clusterConfigForm.get('snapshotClassRWMName').updateValueAndValidity();

      if (this.clusterDetails?.snapshotClassName) {
        this.clusterConfigForm.get('snapshotClassName').patchValue(this.clusterDetails?.snapshotClassName);
        this.clusterConfigForm.get('snapshotClassRWMName').patchValue(this.clusterDetails?.snapshotClassRWMName);
      }
    } else {
      this.clusterConfigForm.get('snapshotClassName').setValue('');
      this.clusterConfigForm.get('snapshotClassName').clearValidators();
      this.clusterConfigForm.get('snapshotClassName').updateValueAndValidity();
      this.clusterConfigForm.get('snapshotClassRWMName').setValue('');
      this.clusterConfigForm.get('snapshotClassRWMName').clearValidators();
      this.clusterConfigForm.get('snapshotClassRWMName').updateValueAndValidity();
    }
    this.clusterConfigForm.markAsDirty();
  }

  onAddiPrometheusChange(checked: boolean): void {
    if (checked) {
      this.clusterConfigForm.get('prometheusUrl').setValidators([Validators.required, DnsValidator.url]);
      this.clusterConfigForm.get('prometheusUrl').updateValueAndValidity();

      if (this.clusterDetails.prometheusUrl) {
        this.clusterConfigForm.get('prometheusUrl').setValue(this.clusterDetails.prometheusUrl || '');
        this.clusterConfigForm.get('prometheusUsername').setValue(this.clusterDetails.prometheusUsername || '');
        this.clusterConfigForm.get('prometheusPassword').setValue(this.clusterDetails.prometheusPassword || '');
      }
    } else {
      this.clusterConfigForm.get('prometheusUrl').clearValidators();
      this.clusterConfigForm.get('prometheusUrl').updateValueAndValidity();

      this.clusterConfigForm.get('prometheusUrl').setValue('');
      this.clusterConfigForm.get('prometheusUsername').setValue('');
      this.clusterConfigForm.get('prometheusPassword').setValue('');
    }
    this.clusterConfigForm.markAsDirty();
  }

  onAddiGrafanaChange(checked: boolean): void {
    if (checked) {
      this.clusterConfigForm.get('grafanaUrl').setValidators([Validators.required, DnsValidator.url]);
      this.clusterConfigForm.get('grafanaUrl').updateValueAndValidity();

      if (this.clusterDetails.grafanaUrl) {
        this.clusterConfigForm.get('grafanaUrl').setValue(this.clusterDetails.grafanaUrl || '');
        this.clusterConfigForm.get('grafanaUsername').setValue(this.clusterDetails.grafanaUsername || '');
        this.clusterConfigForm.get('grafanaPassword').setValue(this.clusterDetails.grafanaPassword || '');
      }
    } else {
      this.clusterConfigForm.get('grafanaUrl').clearValidators();
      this.clusterConfigForm.get('grafanaUrl').updateValueAndValidity();

      this.clusterConfigForm.get('grafanaUrl').setValue('');
      this.clusterConfigForm.get('grafanaUsername').setValue('');
      this.clusterConfigForm.get('grafanaPassword').setValue('');
    }
    this.clusterConfigForm.markAsDirty();
  }

  onAddiLokiChange(checked: boolean): void {
    if (checked) {
      this.clusterConfigForm.get('lokiUrl').setValidators([Validators.required, DnsValidator.url]);
      this.clusterConfigForm.get('lokiUrl').updateValueAndValidity();
      if (this.clusterDetails.lokiUrl) {
        this.clusterConfigForm.get('lokiUrl').setValue(this.clusterDetails.lokiUrl || '');
        this.clusterConfigForm.get('lokiUsername').setValue(this.clusterDetails.lokiUsername || '');
        this.clusterConfigForm.get('lokiPassword').setValue(this.clusterDetails.lokiPassword || '');
        this.clusterConfigForm.get('lokiOrgId').setValue(this.clusterDetails?.lokiOrgId);
      }
    } else {
      this.clusterConfigForm.get('lokiUrl').clearValidators();
      this.clusterConfigForm.get('lokiUrl').updateValueAndValidity();

      this.clusterConfigForm.get('lokiUrl').setValue('');
      this.clusterConfigForm.get('lokiUsername').setValue('');
      this.clusterConfigForm.get('lokiPassword').setValue('');
      this.clusterConfigForm.get('lokiOrgId').setValue('');
    }
    this.clusterConfigForm.markAsDirty();
  }

  onAddiArgoChange(checked: boolean): void {
    if (checked) {
      this.clusterConfigForm.get('argocdURL').setValidators([Validators.required, DnsValidator.url]);
      this.clusterConfigForm.get('argocdURL').updateValueAndValidity();
      this.clusterConfigForm.get('argocdUsername').setValidators([Validators.required]);
      this.clusterConfigForm.get('argocdUsername').updateValueAndValidity();
      this.clusterConfigForm.get('argocdPassword').setValidators([Validators.required]);
      this.clusterConfigForm.get('argocdPassword').updateValueAndValidity();
      this.clusterConfigForm.get('argocdPort').setValidators([Validators.required]);
      this.clusterConfigForm.get('argocdPort').updateValueAndValidity();

      if (this.clusterDetails.argocdURL) {
        this.clusterConfigForm.get('argocdURL').setValue(this.clusterDetails.argocdURL || '');
        this.clusterConfigForm.get('argocdUsername').setValue(this.clusterDetails.argocdUsername || '');
        this.clusterConfigForm.get('argocdPassword').setValue(this.clusterDetails.argocdPassword || '');
        this.clusterConfigForm.get('argocdPort').setValue(this.clusterDetails?.argocdPort);
      }
    } else {
      this.clusterConfigForm.get('argocdURL').setValue('');
      this.clusterConfigForm.get('argocdUsername').setValue('');
      this.clusterConfigForm.get('argocdPassword').setValue('');
      this.clusterConfigForm.get('argocdPort').setValue('');

      this.clusterConfigForm.get('argocdURL').clearValidators();
      this.clusterConfigForm.get('argocdURL').updateValueAndValidity();
      this.clusterConfigForm.get('argocdUsername').clearValidators();
      this.clusterConfigForm.get('argocdUsername').updateValueAndValidity();
      this.clusterConfigForm.get('argocdPassword').clearValidators();
      this.clusterConfigForm.get('argocdPassword').updateValueAndValidity();
      this.clusterConfigForm.get('argocdPort').clearValidators();
      this.clusterConfigForm.get('argocdPort').updateValueAndValidity();
    }
    this.clusterConfigForm.markAsDirty();
  }

  onAddiServiceMeshChange(checked: boolean): void {
    if (checked) {
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomain').setValidators([Validators.required, DnsValidator.host]);
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomain').updateValueAndValidity();

      this.clusterConfigForm.get('istioGatewayName').setValidators([Validators.required]);
      this.clusterConfigForm.get('istioGatewayName').updateValueAndValidity();

      if (this.clusterDetails.serviceMeshEnabled === true) {
        this.clusterConfigForm.get('serviceMeshGatewayWildcardDomain').setValue(this.clusterDetails.serviceMeshGatewayWildcardDomain || '');
        this.clusterConfigForm.get('istioGatewayName').setValue(this.clusterDetails.istioGatewayName || '');

        if (this.clusterDetails?.serviceMeshGatewayWildcardDomainTls) {
          this.serviceMeshSslEnabled = true;
          this.onServiceMeshSslChange(true);
        }
        if (this.clusterDetails?.kialiURL) {
          this.enableKiali = true;
          this.onAddiKialiChange(true);
        }
      }
    } else {
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomain').setValue('');
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomain').clearValidators();
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomain').updateValueAndValidity();
      this.clusterConfigForm.get('istioGatewayName').setValue('');
      this.clusterConfigForm.get('istioGatewayName').clearValidators();
      this.clusterConfigForm.get('istioGatewayName').updateValueAndValidity();

      if (this.enableKiali) {
        this.enableKiali = false;
        this.clusterConfigForm.get('kialiURL').setValue('');
        this.clusterConfigForm.get('kialiToken').setValue('');

        this.clusterConfigForm.get('kialiURL').clearValidators();
        this.clusterConfigForm.get('kialiURL').updateValueAndValidity();
        this.clusterConfigForm.get('kialiToken').clearValidators();
        this.clusterConfigForm.get('kialiToken').updateValueAndValidity();
      }
      if (this.serviceMeshSslEnabled) {
        this.serviceMeshSslEnabled = false;
        this.clusterConfigForm.get('serviceMeshGatewayWildcardDomainTls').setValue('');
        this.clusterConfigForm.get('serviceMeshGatewayWildcardDomainTls').clearValidators();
        this.clusterConfigForm.get('serviceMeshGatewayWildcardDomainTls').updateValueAndValidity();
      }
    }
    this.clusterConfigForm.markAsDirty();
  }

  onServiceMeshSslChange(checked: boolean): void {
    if (checked) {
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomainTls').setValidators([Validators.required]);
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomainTls').updateValueAndValidity();

      if (this.clusterDetails.serviceMeshGatewayWildcardDomainTls) {
        this.clusterConfigForm
          .get('serviceMeshGatewayWildcardDomainTls')
          .setValue(this.clusterDetails.serviceMeshGatewayWildcardDomainTls || '');
      }
    } else {
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomainTls').setValue('');
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomainTls').clearValidators();
      this.clusterConfigForm.get('serviceMeshGatewayWildcardDomainTls').updateValueAndValidity();
    }
    this.clusterConfigForm.markAsDirty();
  }

  onAddiKialiChange(checked): void {
    if (checked) {
      this.clusterConfigForm.get('kialiURL').setValidators([Validators.required, DnsValidator.url]);
      this.clusterConfigForm.get('kialiURL').updateValueAndValidity();
      this.clusterConfigForm.get('kialiToken').setValidators([Validators.required]);
      this.clusterConfigForm.get('kialiToken').updateValueAndValidity();
      if (this.clusterDetails?.kialiURL) {
        this.clusterConfigForm.get('kialiURL').setValue(this.clusterDetails?.kialiURL);
        this.clusterConfigForm.get('kialiToken').setValue(this.clusterDetails?.kialiToken);
      }
    } else {
      this.clusterConfigForm.get('kialiURL').setValue('');
      this.clusterConfigForm.get('kialiToken').setValue('');
      this.clusterConfigForm.get('kialiURL').clearValidators();
      this.clusterConfigForm.get('kialiURL').updateValueAndValidity();
      this.clusterConfigForm.get('kialiToken').clearValidators();
      this.clusterConfigForm.get('kialiToken').updateValueAndValidity();
    }
    this.clusterConfigForm.markAsDirty();
  }

  patchFormData(clusterData: any) {
    if (clusterData?.clusterIssuerName) this.clusterConfigForm.get('clusterIssuerName').patchValue(clusterData?.clusterIssuerName);

    // Volume Snapshot
    if (clusterData?.snapshotClassName) {
      this.enableVolumeSnapshot = true;
      this.onAddiVolumeSnapshotChange(true);
    }
    // Promethus
    if (clusterData?.prometheusUrl) {
      this.enablePrometheus = true;
      this.onAddiPrometheusChange(true);
    }
    // Grafana
    if (clusterData?.grafanaUrl) {
      this.enableGrafana = true;
      this.onAddiGrafanaChange(true);
    }
    // Loki
    if (clusterData?.lokiUrl) {
      this.enableLoki = true;
      this.onAddiLokiChange(true);
    }
    // ArgoCD
    if (clusterData?.argocdURL) {
      this.enableArgo = true;
      this.onAddiArgoChange(true);
    }
    // Service Mesh
    if (clusterData?.serviceMeshEnabled) {
      this.clusterConfigForm.get('serviceMeshEnabled').patchValue(true);
      this.onAddiServiceMeshChange(true);
    }
  }

  getClusterDetails() {
    this._clusterService.getClusterById(this.clusterId).subscribe({
      next: (res: any) => {
        if (res.status === 'success') {
          this.clusterDetails = res?.data;
          this.patchFormData(this.clusterDetails);
        }
        this.isLoading = false;
      },
      error: err => {
        this.isLoading = false;
        this.toastrService.error(err?.message || "Something is wrong, can't fetch existing data.");
      }
    });
  }

  onUpdate(): void {
    this.isSubmitting = true;
    const formData = this.clusterConfigForm.value;
    formData['id'] = this.clusterId;
    this._clusterService.updateCluster(formData).subscribe(
      (res: any) => {
        if (res?.status === 'success') {
          this.clusterConfigForm.markAsPristine();
          this.toastrService.success('Success! Cluster configuration updated.');
        } else {
          this.toastrService.error(res?.message);
        }
        this.isSubmitting = false;
      },
      err => {
        this.isSubmitting = false;
        this.toastrService.error(err?.message);
      }
    );
  }
}
