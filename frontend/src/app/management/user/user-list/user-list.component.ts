import { ChangeDetectorRef, Component, Inject, OnInit, Optional } from '@angular/core';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icMail from '@iconify/icons-ic/twotone-mail';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icAdd from '@iconify/icons-ic/twotone-add';
import icInfo from '@iconify/icons-ic/twotone-info';
import icLock from '@iconify/icons-ic/lock';
import icMoreHoriz from '@iconify/icons-ic/twotone-more-horiz';
import icMoreVert from '@iconify/icons-ic/twotone-more-vert';
import people_outline from '@iconify/icons-ic/twotone-people-outline';
import { UserService } from '../user.service';
import { User } from '../user';
import { MatDialog } from '@angular/material/dialog';
import { UserFormComponent } from '../user-form/user-form.component';
import { BaseTableComponenet } from '@sdk-ui/components';
import { ToolbarService } from '@sdk-ui/services';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui/confirm-dialog-static/confirm-dialog-static.component';
import { ToastrService } from '@sdk-ui/ui';
import { RequesterService } from '@core-ui/services/requester.service';

@Component({
  selector: 'kc-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss']
})
export class UserListComponent extends BaseTableComponenet implements OnInit {
  icEdit = icEdit;
  icDelete = icDelete;
  icAdd = icAdd;
  icLock = icLock;
  icMoreHoriz = icMoreHoriz;
  icMoreVert = icMoreVert;
  icInfo = icInfo;
  icMail = icMail;
  people_outline = people_outline;

  isLoading!: boolean;
  tableColumns = ['checkbox', 'username', 'name', 'type', 'status', 'created', 'action'];
  users: User[];
  searchUser!: string;
  currentUser: any;

  constructor(
    cdr: ChangeDetectorRef,
    public toolbarService: ToolbarService,
    public dialog: MatDialog,
    public toastr: ToastrService,
    private _userService: UserService,
    private requester: RequesterService
  ) {
    super(cdr);
  }

  ngOnInit() {
    this.currentUser = this.requester.get();
    this.toolbarService.changeData({ title: 'Users' });
    this.getUsers();
  }

  getUsers(): void {
    this.isLoading = true;
    this._userService.mcGetUsers().subscribe({
      next: data => {
        this.users = data;
        this.isLoading = false;
      },
      error: err => {
        this.isLoading = false;
        this.toastr.error(err.message || 'Something wrong on fetching user');
      }
    });
  }

  userForm(user?: any): void {
    const dialogRef = this.dialog.open(UserFormComponent, {
      disableClose: true,
      width: '800px',
      maxWidth: '800px',
      maxHeight: '95vh',
      data: user || null
    });
    dialogRef.afterClosed().subscribe(isUpdated => {
      if (isUpdated) this.getUsers();
    });
  }

  updatePassword(user?: any): void {
    import('../update-password/update-password.component').then(m => {
      const dialogRef = this.dialog.open(m.UpdatePasswordComponent, {
        disableClose: true,
        width: '500px',
        maxWidth: '500px',
        data: user || null
      });
      dialogRef.afterClosed().subscribe(isUpdated => {
        if (isUpdated) this.getUsers();
      });
    });
  }

  deleteUser(user) {
    const dialogRef = this.dialog.open(ConfirmDialogStaticComponent, {
      width: '350px',
      data: {
        message: `Are you sure! want to delete "${user.username}" user?`,
        icon: '/assets/img/bin.svg'
      }
    });

    dialogRef.afterClosed().subscribe((result: any) => {
      if (result) {
        this._userService.mcDeleteUser(user.id).subscribe({
          next: _ => {
            this.toastr.success('User deleted.');
            this.getUsers();
          },
          error: err => {
            this.toastr.error(err.message || 'Something Wrong on deleting team');
          }
        });
      }
    });
  }
}
