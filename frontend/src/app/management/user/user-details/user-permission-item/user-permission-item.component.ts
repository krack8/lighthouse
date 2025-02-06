import { animate, style, transition, trigger } from '@angular/animations';
import { Component, Input, OnInit } from '@angular/core';
import icArrowDropDown from '@iconify/icons-ic/arrow-drop-down';
import icArrow from '@iconify/icons-ic/twotone-arrow-right';

export const fadeAnimation = trigger('fadeAnimation', [
  transition(':enter', [
    style({ height: '0px', opacity: 0, visibility: 'hidden' }),
    animate('300ms cubic-bezier(.37,1.04,.68,.98)', style({ height: '*', opacity: 1, visibility: 'visible' }))
  ]),
  transition(':leave', [animate('300ms cubic-bezier(.37,1.04,.68,.98)', style({ height: '0px', opacity: 0, visibility: 'hidden' }))])
]);

@Component({
  selector: 'kc-user-permission-item',
  templateUrl: './user-permission-item.component.html',
  styleUrls: ['./user-permission-item.component.scss'],
  animations: [fadeAnimation]
})
export class UserPermissionItemComponent implements OnInit {
  @Input() role: any;
  icArrowDropDown = icArrowDropDown;
  icArrow = icArrow;
  showPermissions: boolean = false;

  constructor() {}

  ngOnInit(): void {
    console.log(this.role);
  }

  togglePermission(): void {
    this.showPermissions = !this.showPermissions;
  }
}
