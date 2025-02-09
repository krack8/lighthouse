import { Component, OnInit } from '@angular/core';
import icBackspace from '@iconify/icons-ic/arrow-back';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'kc-error403',
  templateUrl: './error-403.component.html',
  styleUrls: ['./error-403.component.scss']
})
export class Error403Component implements OnInit {
  icBackspace = icBackspace;

  constructor(
    private _location: Location,
    private activatedRoute: ActivatedRoute
  ) {}

  ngOnInit() {
    const replaceUrl = this.activatedRoute.snapshot.queryParams['url'];
    if (replaceUrl) {
      this._location.replaceState(replaceUrl);
    }
  }

  goBack(): void {
    this._location.back();
  }
}
