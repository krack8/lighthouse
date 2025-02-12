import { Component, OnInit } from '@angular/core';
import { animate, style, transition, trigger } from '@angular/animations';

@Component({
  selector: 'kc-grafana-dashboard',
  templateUrl: './grafana-dashboard.component.html',
  styleUrls: ['./grafana-dashboard.component.scss'],
  animations: [
    transition(':enter', [style({ opacity: 0 }), animate('0.2s', style({ opacity: 1 }))]),
    transition(':leave', [animate('0.2s', style({ opacity: 0 }))])
  ]
})
export class GrafanaDashboardComponent implements OnInit {
  constructor() {}

  ngOnInit(): void {}
}
