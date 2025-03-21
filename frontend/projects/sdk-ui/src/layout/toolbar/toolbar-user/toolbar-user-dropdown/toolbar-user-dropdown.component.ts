import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { PopoverRef } from '@sdk-ui/ui';
import { RequesterService, PermissionService } from '@core-ui/services';
import { ToolbarUserMenuItem } from '@sdk-ui/interfaces';
import { trackById } from '@core-ui/utils';
import { SdkConfigService } from '@sdk-ui/services';

@Component({
  selector: 'kc-toolbar-user-dropdown',
  templateUrl: './toolbar-user-dropdown.component.html',
  styleUrls: ['./toolbar-user-dropdown.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ToolbarUserDropdownComponent implements OnInit {
  items: ToolbarUserMenuItem[] = [];

  trackById = trackById;
  fullName!: string;
  userData: any;

  constructor(
    private popoverRef: PopoverRef<ToolbarUserDropdownComponent>,
    private router: Router,
    public requester: RequesterService,
    private permissionService: PermissionService,
    private SdkonfigService: SdkConfigService
  ) {}

  ngOnInit() {
    this.userData = this.requester.get();
    if (this.userData?.userInfo?.first_name || this.userData?.userInfo?.last_name) {
      this.fullName = `${this.userData.userInfo.first_name} ${this.userData.userInfo.last_name}`;
    } else {
      this.fullName = this.userData?.userInfo?.username?.split('@')?.[0];
    }

    this.items = this.SdkonfigService.initializeToolbarUserDropdown.filter(item => {
      if (item.permissions === undefined || this.userData?.userInfo?.user_type === 'ADMIN') return true;
      return this.permissionService.userPermissionsSnapshot.includes(item.permissions);
    });
  }

  close() {
    this.popoverRef.close();
  }

  logout() {
    this.requester.logoutUser().subscribe({
      next: () => {
        this.requester.clear();
        this.router.navigate(['/auth/login']);
        this.permissionService.loadUserPermissions([]);
        this.popoverRef.close();
      },
      error: () => {
        console.warn('Error while logging out');
      }
    });
  }
}
