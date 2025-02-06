import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import icSearch from '@iconify/icons-ic/twotone-search';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'kc-error404',
  templateUrl: './error-404.component.html',
  styleUrls: ['./error-404.component.scss']
})
export class Error404Component implements OnInit {
  icSearch = icSearch;

  constructor(
    private activatedRoute: ActivatedRoute,
    private location: Location
  ) {}

  ngOnInit() {
    const replaceUrl = this.activatedRoute.snapshot.queryParams['url'];
    if (replaceUrl) {
      this.location.replaceState(replaceUrl);
    }
  }
}
