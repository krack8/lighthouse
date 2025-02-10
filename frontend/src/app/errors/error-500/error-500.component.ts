import { Component, OnInit } from '@angular/core';
import icSearch from '@iconify/icons-ic/twotone-search';

@Component({
  selector: 'kc-error-500',
  templateUrl: './error-500.component.html',
  styleUrls: ['./error-500.component.scss']
})
export class Error500Component implements OnInit {
  icSearch = icSearch;

  constructor() {}

  ngOnInit() {}
}
