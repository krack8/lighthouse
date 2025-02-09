import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { UserService } from '../user.service';
import icMenu from '@iconify/icons-ic/twotone-menu';
import icArrowBack from '@iconify/icons-ic/arrow-back';
import { ToastrService } from '@sdk-ui/ui';
import { MatDialog } from '@angular/material/dialog';
@Component({
  selector: 'kc-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.scss']
})
export class UserDetailsComponent implements OnInit {
  icMenu = icMenu;
  icArrowBack = icArrowBack;

  user: any;
  isLoading!: boolean;
  isLoaded!: boolean;

  constructor(
    private _route: ActivatedRoute,
    private _userService: UserService,
    private toastr: ToastrService,
    private dialog: MatDialog
  ) {}

  ngOnInit() {
    this.getUser();
  }

  getUser(): void {
    this.isLoading = true;
    this._userService.mcGetUser(this._route.snapshot.params['id']).subscribe(
      data => {
        this.isLoading = false;
        this.user = data;
        if (!this.isLoaded) this.isLoaded = true;
      },
      err => {
        this.isLoading = false;
        this.user = null;
        this.toastr.error(err.message || "Can't fetch user");
      }
    );
  }

  updateRole(): void {
    import('./user-role-update-form/user-role-update-form.component').then(m => {
      const dialogRef = this.dialog.open(m.UserRoleUpdateFormComponent, {
        width: '100%',
        maxWidth: '600px',
        data: this.user
      });
      dialogRef.afterClosed().subscribe(res => {
        if (res) this.getUser();
      });
    });
  }
}
