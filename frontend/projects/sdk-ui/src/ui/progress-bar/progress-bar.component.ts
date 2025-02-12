import { Component, OnInit } from '@angular/core';
import { LoadingBarService } from '../loading-bar';

@Component({
  selector: 'kc-progress-bar',
  templateUrl: './progress-bar.component.html',
  styleUrls: ['./progress-bar.component.scss']
})
export class ProgressBarComponent implements OnInit {
  constructor(public loader: LoadingBarService) {}

  ngOnInit() {}
}
