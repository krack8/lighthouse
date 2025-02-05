import { Component, OnInit } from '@angular/core';
import { NavigationService } from '@sdk-ui/services';

@Component({
  selector: 'kc-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {
  items = this.navigationService.items;

  constructor(private navigationService: NavigationService) {}

  ngOnInit() {}
}
