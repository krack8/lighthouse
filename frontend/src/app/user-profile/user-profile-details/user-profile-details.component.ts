import { Component, OnInit } from '@angular/core';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { RequesterService } from '@core-ui/services/requester.service';
import { MatDialog } from '@angular/material/dialog';
import { ToastrService } from '@sdk-ui/ui';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'kc-user-profile-details',
  templateUrl: './user-profile-details.component.html',
  styleUrls: ['./user-profile-details.component.scss']
})
export class UserProfileDetailsComponent implements OnInit {
  isLoading = true;
  userData: any;

  constructor(
    public dialog: MatDialog,
    public toastr: ToastrService,
    public toolbarService: ToolbarService,
    public snackBar: MatSnackBar,
    private requesterService: RequesterService
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: 'Account' });
    this.userData = this.requesterService.get();
  }

  changePassword(): void {
    import('../change-password/change-password.component').then(m => {
      this.dialog.open(m.ChangePasswordComponent, {
        disableClose: true,
        width: '500px',
        maxWidth: '500px'
      });
    });
  }
}
