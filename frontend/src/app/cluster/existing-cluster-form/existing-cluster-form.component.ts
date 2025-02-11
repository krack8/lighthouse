import { AfterContentInit, Component, Inject, OnInit } from '@angular/core';
import { FormsModule, ReactiveFormsModule, UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { fadeInUp400ms } from '@sdk-ui/animations/fade-in-up.animation';
import { ClusterService } from '@cluster/cluster.service';
import { PageLayoutModule, ToastrService } from '@sdk-ui/ui';
import { Router } from '@angular/router';
import icInfo from '@iconify/icons-ic/info';
import icVisibility from '@iconify/icons-ic/visibility';
import icVisibilityOff from '@iconify/icons-ic/visibility-off';
import { CommonModule, DOCUMENT } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { IconModule } from '@visurel/iconify-angular';
import { CdkClipboardModule } from '@cdk-ui/clipboard';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatCheckboxChange, MatCheckboxModule } from '@angular/material/checkbox';
import { MatSelectModule } from '@angular/material/select';
import { CdkStepperModule, StepperSelectionEvent } from '@angular/cdk/stepper';
import { CdkHorizontalStepperModule } from '@cdk-ui/horizontal-stepper';
import { MatTooltipModule } from '@angular/material/tooltip';
import { CdkHintModule } from '@cdk-ui/hint';
import { OnboardClusterPrerequisiteComponent } from './onboard-cluster-prerequisite/onboard-cluster-prerequisite.component';
import { DnsValidator } from '@shared-ui/validators';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';

@Component({
  selector: 'kc-existing-cluster-form',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    MatIconModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatCheckboxModule,
    MatSelectModule,
    MatProgressBarModule,
    MatTooltipModule,
    PageLayoutModule,
    CdkHorizontalStepperModule,
    CdkClipboardModule,
    CdkStepperModule,
    CdkHintModule,
    IconModule,
    OnboardClusterPrerequisiteComponent
  ],
  templateUrl: './existing-cluster-form.component.html',
  styleUrls: ['./existing-cluster-form.component.scss'],
  animations: [fadeInUp400ms]
})
export class ExistingClusterFormComponent implements OnInit, AfterContentInit {
  icInfo = icInfo;
  icVisibilityOff = icVisibilityOff;
  icVisibility = icVisibility;

  isSubmitting: boolean = false;
  clusterInfoFormGroup: UntypedFormGroup = this.fb.group({});

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

  sidenavContent: any;

  constructor(
    @Inject(DOCUMENT) public document: any,
    private fb: UntypedFormBuilder,
    private toolbarService: ToolbarService,
    private clusterService: ClusterService,
    private toastr: ToastrService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: 'On-Board Cluster' });
    this.clusterInfoFormGroup = this.fb.group({
      clusterName: ['', [Validators.required, Validators.minLength(4), Validators.pattern(/^[a-z][a-z0-9-]+$/)]],
      namespace: ['klovercloud', [Validators.minLength(4), Validators.pattern(/^[a-z][a-z0-9-]+$/)]],

      // 01. Ingress
      clusterIssuerName: [''],
      // 01.01 Service Ingress
      nginxIngressGatewayWildcardDomain: ['', [Validators.required, DnsValidator.host]],
      nginxIngressControllerType: ['NGINX', [Validators.required]],
      nginxIngressClassName: ['nginx', [Validators.required, Validators.pattern(/^[a-zA-Z0-9-]+$/)]],
      nginxIngressGatewayWildcardDomainTlsEnabled: [false],

      // 01.02 Application Ingress
      useSameIngressGatewayConfigForApplication: [true],

      // 02. Storage Information
      volumeStorageClassReadWriteOnce: ['', [Validators.required, Validators.pattern(/^[a-zA-Z0-9-]+$/)]],
      volumeStorageClassReadWriteMany: ['', [Validators.required, Validators.pattern(/^[a-zA-Z0-9-]+$/)]],

      // 03. Additional Information
      // 03.01 Volume Snapshot Information
      // 03.02 Monitoring Information
      // 03.02.01 Volume Snapshot Information
      // 03.02.02 Grafana Configuration
      // 03.03 Log Information
      // 03.03.01 Loki Server Configuration
      // 03.04 ArgoCD Configuration
      // 03.05 Service Mesh Configuration
      serviceMeshEnabled: [false, [Validators.required]]
      // Implementation Tool Type Controll
      // 03.05.01 Istio Configuration
      // 03.05.02 Kiali Configuration
    });
    this.clusterInfoFormGroup
      .get('clusterIssuerName')
      .valueChanges.pipe(debounceTime(300), distinctUntilChanged())
      .subscribe(val => {
        if (val) {
          if (
            this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsAutoSSL') &&
            this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsAutoSSL').disabled
          ) {
            this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsAutoSSL').enable();
          }
          // Application SSL
          if (
            this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsAutoSSL') &&
            this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsAutoSSL').disabled
          ) {
            this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsAutoSSL').enable();
          }
        } else {
          if (
            this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsAutoSSL') &&
            this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsAutoSSL').enable
          ) {
            if (this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsAutoSSL').value) {
              this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsAutoSSL').setValue(false);
            }
            this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsAutoSSL').disable();
          }
          // Application SSL
          if (
            this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsAutoSSL') &&
            this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsAutoSSL').enable
          ) {
            if (this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsAutoSSL').value) {
              this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsAutoSSL').setValue(false);
            }
            this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsAutoSSL').disable();
          }
        }
      });
  }

  ngAfterContentInit(): void {
    this.sidenavContent = this.document.querySelector('.sidenav-content');
  }

  onStepperChange(_: StepperSelectionEvent): void {
    if (this.sidenavContent?.scrollTop > 70) this.sidenavContent.scrollTop = 0;
  }

  // Service Ingress Start
  onServiceSslChange(event: MatCheckboxChange): void {
    if (event.checked === true) {
      this.clusterInfoFormGroup.addControl(
        'nginxIngressGatewayWildcardDomainTlsAutoSSL',
        this.fb.control({
          value: false,
          disabled: this.clusterInfoFormGroup.get('clusterIssuerName').value ? false : true
        })
      );
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsCert', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsKey', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsCa', this.fb.control(''));
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsSecretReflectionEnabled', this.fb.control(false));
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsSecretEnabled', this.fb.control(false));

      this.clusterInfoFormGroup
        .get('nginxIngressGatewayWildcardDomainTlsAutoSSL')
        .valueChanges.pipe(distinctUntilChanged())
        .subscribe(bool => {
          if (bool) {
            if (this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsSecretEnabled').value === true) {
              this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsSecret');
            } else {
              this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsCert');
              this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsKey');
              this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsCa');
              this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsSecretReflectionEnabled');
            }
            this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsSecretEnabled');
          } else {
            this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsCert', this.fb.control('', [Validators.required]));
            this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsKey', this.fb.control('', [Validators.required]));
            this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsCa', this.fb.control(''));
            this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsSecretReflectionEnabled', this.fb.control(false));
            this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsSecretEnabled', this.fb.control(false));
          }
        });
    } else {
      if (this.clusterInfoFormGroup.get('nginxIngressGatewayWildcardDomainTlsSecretEnabled')?.value === true) {
        this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsSecret');
      } else {
        this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsCert');
        this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsKey');
        this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsCa');
        this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsSecretReflectionEnabled');
      }
      this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsAutoSSL');
      this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsSecretEnabled');
    }
  }

  onServiceSslCertificateChange(event: MatCheckboxChange): void {
    if (event.checked) {
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsSecret', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsCert');
      this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsKey');
      this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsCa');
      this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsSecretReflectionEnabled');
    } else {
      this.clusterInfoFormGroup.removeControl('nginxIngressGatewayWildcardDomainTlsSecret');
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsCert', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsKey', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsCa', this.fb.control(''));
      this.clusterInfoFormGroup.addControl('nginxIngressGatewayWildcardDomainTlsSecretReflectionEnabled', this.fb.control(false));
    }
  }
  // Service Ingress End
  // Application Ingress Start
  onAppIngressChange(e: MatCheckboxChange): void {
    if (e.checked) {
      this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomain');
      this.clusterInfoFormGroup.removeControl('nginxIngressControllerTypeForApplication');
      this.clusterInfoFormGroup.removeControl('nginxIngressClassNameForApplication');
      if (this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsEnabled')?.value === true) this.onAppSslChange(false);
      this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsEnabled');
    } else {
      this.clusterInfoFormGroup.addControl(
        'applicationIngressGatewayWildcardDomain',
        this.fb.control('', [Validators.required, DnsValidator.host])
      );
      this.clusterInfoFormGroup.addControl('nginxIngressControllerTypeForApplication', this.fb.control('NGINX', [Validators.required]));
      this.clusterInfoFormGroup.addControl(
        'nginxIngressClassNameForApplication',
        this.fb.control('nginx', [Validators.required, Validators.pattern(/^[a-zA-Z0-9-]+$/)])
      );
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsEnabled', this.fb.control(false));
    }
  }
  onAppSslChange(checked: boolean): void {
    if (checked === true) {
      this.clusterInfoFormGroup.addControl(
        'applicationIngressGatewayWildcardDomainTlsAutoSSL',
        this.fb.control({
          value: false,
          disabled: this.clusterInfoFormGroup.get('clusterIssuerName').value ? false : true
        })
      );
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsCert', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsKey', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsCa', this.fb.control(''));
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsSecretReflectionEnabled', this.fb.control(false));
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsSecretEnabled', this.fb.control(false));
      this.clusterInfoFormGroup
        .get('applicationIngressGatewayWildcardDomainTlsAutoSSL')
        .valueChanges.pipe(distinctUntilChanged())
        .subscribe(bool => {
          if (bool) {
            if (this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsSecretEnabled')?.value === true) {
              this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsSecret');
            } else {
              this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsCert');
              this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsKey');
              this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsCa');
              this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsSecretReflectionEnabled');
            }
            this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsSecretEnabled');
          } else {
            this.clusterInfoFormGroup.addControl(
              'applicationIngressGatewayWildcardDomainTlsCert',
              this.fb.control('', [Validators.required])
            );
            this.clusterInfoFormGroup.addControl(
              'applicationIngressGatewayWildcardDomainTlsKey',
              this.fb.control('', [Validators.required])
            );
            this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsCa', this.fb.control(''));
            this.clusterInfoFormGroup.addControl(
              'applicationIngressGatewayWildcardDomainTlsSecretReflectionEnabled',
              this.fb.control(false)
            );
            this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsSecretEnabled', this.fb.control(false));
          }
        });
    } else {
      if (this.clusterInfoFormGroup.get('applicationIngressGatewayWildcardDomainTlsSecretEnabled')?.value === true) {
        this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsSecret');
      } else {
        this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsCert');
        this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsKey');
        this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsCa');
        this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsSecretReflectionEnabled');
      }
      this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsAutoSSL');
      this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsSecretEnabled');
    }
  }
  onAppSslCertificateChange(event: MatCheckboxChange): void {
    if (event.checked) {
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsSecret', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsCert');
      this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsKey');
      this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsCa');
      this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsSecretReflectionEnabled');
    } else {
      this.clusterInfoFormGroup.removeControl('applicationIngressGatewayWildcardDomainTlsSecret');
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsCert', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsKey', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsCa', this.fb.control(''));
      this.clusterInfoFormGroup.addControl('applicationIngressGatewayWildcardDomainTlsSecretReflectionEnabled', this.fb.control(false));
    }
  }
  // Application Ingress End

  // Additional Config Start
  onAddiVolumeSnapshotChange(e: MatCheckboxChange): void {
    if (e.checked) {
      this.clusterInfoFormGroup.addControl(
        'snapshotClassName',
        this.fb.control('', [Validators.required, Validators.pattern(/^[a-zA-Z0-9-]+$/)])
      );
      this.clusterInfoFormGroup.addControl(
        'snapshotClassRWMName',
        this.fb.control('', [Validators.required, Validators.pattern(/^[a-zA-Z0-9-]+$/)])
      );
    } else {
      this.clusterInfoFormGroup.removeControl('snapshotClassName');
      this.clusterInfoFormGroup.removeControl('snapshotClassRWMName');
    }
  }

  onAddiPrometheusChange(e: MatCheckboxChange): void {
    if (e.checked) {
      this.clusterInfoFormGroup.addControl('prometheusURL', this.fb.control('', [Validators.required, DnsValidator.url]));
      this.clusterInfoFormGroup.addControl('prometheusUsername', this.fb.control(''));
      this.clusterInfoFormGroup.addControl('prometheusPassword', this.fb.control(''));
    } else {
      this.clusterInfoFormGroup.removeControl('prometheusURL');
      this.clusterInfoFormGroup.removeControl('prometheusUsername');
      this.clusterInfoFormGroup.removeControl('prometheusPassword');
    }
  }

  onAddiGrafanaChange(e: MatCheckboxChange): void {
    if (e.checked) {
      this.clusterInfoFormGroup.addControl('grafanaURL', this.fb.control('', [Validators.required, DnsValidator.url]));
      this.clusterInfoFormGroup.addControl('grafanaUsername', this.fb.control(''));
      this.clusterInfoFormGroup.addControl('grafanaPassword', this.fb.control(''));
    } else {
      this.clusterInfoFormGroup.removeControl('grafanaURL');
      this.clusterInfoFormGroup.removeControl('grafanaUsername');
      this.clusterInfoFormGroup.removeControl('grafanaPassword');
    }
  }

  onAddiLokiChange(e: MatCheckboxChange): void {
    if (e.checked) {
      this.clusterInfoFormGroup.addControl('lokiURL', this.fb.control('', [Validators.required, DnsValidator.url]));
      this.clusterInfoFormGroup.addControl('lokiUsername', this.fb.control(''));
      this.clusterInfoFormGroup.addControl('lokiPassword', this.fb.control(''));
      this.clusterInfoFormGroup.addControl('lokiOrgId', this.fb.control(''));
    } else {
      this.clusterInfoFormGroup.removeControl('lokiURL');
      this.clusterInfoFormGroup.removeControl('lokiUsername');
      this.clusterInfoFormGroup.removeControl('lokiPassword');
      this.clusterInfoFormGroup.removeControl('lokiOrgId');
    }
  }

  onAddiArgoChange(e: MatCheckboxChange): void {
    if (e.checked) {
      this.clusterInfoFormGroup.addControl('argocdURL', this.fb.control('', [Validators.required, DnsValidator.url]));
      this.clusterInfoFormGroup.addControl('argocdUsername', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('argocdPassword', this.fb.control('', [Validators.required]));
      this.clusterInfoFormGroup.addControl('argocdPort', this.fb.control('', [Validators.required]));
    } else {
      this.clusterInfoFormGroup.removeControl('argocdURL');
      this.clusterInfoFormGroup.removeControl('argocdUsername');
      this.clusterInfoFormGroup.removeControl('argocdPassword');
      this.clusterInfoFormGroup.removeControl('argocdPort');
    }
  }
  onAddiServiceMeshChange(e: MatCheckboxChange): void {
    if (e.checked) {
      this.clusterInfoFormGroup.addControl(
        'serviceMeshGatewayWildcardDomain',
        this.fb.control('', [Validators.required, DnsValidator.host])
      );
      this.clusterInfoFormGroup.addControl('istioGatewayName', this.fb.control('', [Validators.required]));
    } else {
      if (this.enableKiali) {
        this.enableKiali = false;
        this.clusterInfoFormGroup.removeControl('kialiURL');
        this.clusterInfoFormGroup.removeControl('kialiToken');
      }
      if (this.serviceMeshSslEnabled) {
        this.serviceMeshSslEnabled = false;
        this.clusterInfoFormGroup.removeControl('serviceMeshGatewayWildcardDomainTls');
      }
      this.clusterInfoFormGroup.removeControl('serviceMeshGatewayWildcardDomain');
      this.clusterInfoFormGroup.removeControl('istioGatewayName');
      this.clusterInfoFormGroup.removeControl('kialiEnabled');
    }
  }

  onServiceMeshSslChange(e: MatCheckboxChange): void {
    if (e.checked) {
      this.clusterInfoFormGroup.addControl('serviceMeshGatewayWildcardDomainTls', this.fb.control('', [Validators.required]));
    } else {
      this.clusterInfoFormGroup.removeControl('serviceMeshGatewayWildcardDomainTls');
    }
  }

  onAddiKialiChange(e: MatCheckboxChange): void {
    if (e.checked) {
      this.clusterInfoFormGroup.addControl('kialiURL', this.fb.control('', [Validators.required, DnsValidator.url]));
      this.clusterInfoFormGroup.addControl('kialiToken', this.fb.control('', [Validators.required]));
    } else {
      this.clusterInfoFormGroup.removeControl('kialiURL');
      this.clusterInfoFormGroup.removeControl('kialiToken');
    }
  }

  // Additional Config End
  createCluster(): void {
    this.isSubmitting = true;
    this.clusterService.mcOnboardExistingCluster(this.clusterInfoFormGroup.value).subscribe({
      next: res => {
        if (res.status === 'CREATED') {
          if (res.data) {
            // Keep this delay to solve first cluster creation redirect 404 bug
            setTimeout(() => {
              this.toastr.success(res['message']);
              this.isSubmitting = false;
              this.router.navigate(['/clusters', res.data]);
            }, 5000);
          } else {
            this.isSubmitting = false;
            this.router.navigate(['/clusters']);
          }
          return;
        }
        this.toastr.error(res['message']);
      },
      error: error => {
        this.isSubmitting = false;
        this.toastr.error(error.message);
      }
    });
  }
}
