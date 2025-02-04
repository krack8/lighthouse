import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CdkClipboardModule } from '@cdk-ui/clipboard';
import { IconModule } from '@visurel/iconify-angular';
import icExpendeLess from '@iconify/icons-ic/expand-less';
import icExpendeMore from '@iconify/icons-ic/expand-more';
import icInfo from '@iconify/icons-fa-solid/info-circle';

import { AceEditorModule } from '@klovercloud/ace-editor/ace-editor.module';
import { CdkTooltipModule } from '@cdk-ui/tooltip';

import 'ace-builds';
import 'ace-builds/src-noconflict/mode-yaml';

@Component({
  selector: 'kc-onboard-cluster-prerequisite',
  standalone: true,
  imports: [CommonModule, CdkClipboardModule, IconModule, AceEditorModule, CdkTooltipModule],
  templateUrl: './onboard-cluster-prerequisite.component.html',
  styleUrls: ['./onboard-cluster-prerequisite.component.scss']
})
export class OnboardClusterPrerequisiteComponent {
  icExpendeLess = icExpendeLess;
  icExpendeMore = icExpendeMore;
  icInfo = icInfo;

  showEssentialRequirement = true;
  showOptionalRequirement = true;

  cartManagerConf = `apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: cluster-letsencrypt
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: support@klovercloud.com
    privateKeySecretRef:
      name: cluster-letsencrypt-key
    solvers:
      - http01:
          ingress:
            class: nginx
`;

  istioConf = `apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: default-istio-gateway
  namespace: istio-system
spec:
  selector:
    istio: ingressgateway
  servers:
    - hosts:
        - '*.<domain>'
      port:
        name: http
        number: 80
        protocol: HTTP`;
}
