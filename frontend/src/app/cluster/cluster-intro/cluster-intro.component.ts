import { Component } from '@angular/core';
import { fadeInUp400ms } from '@sdk-ui/animations/fade-in-up.animation';

@Component({
  selector: 'kc-cluster-intro',
  templateUrl: './cluster-intro.component.html',
  styleUrls: ['./cluster-intro.component.scss'],
  animations: [fadeInUp400ms]
})
export class ClusterIntroComponent {
  constructor() {}
}
