import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FlexLayoutModule } from '@angular/flex-layout';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
// Materail Module
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatDialogModule } from '@angular/material/dialog';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatSortModule } from '@angular/material/sort';
import { MatTableModule } from '@angular/material/table';
import { MatTooltipModule } from '@angular/material/tooltip';
// kc Module
import { ContainerModule } from '@sdk-ui/directives/container/container.module';
import { PageLayoutModule } from '@sdk-ui/ui/page-layout/page-layout.module';
import { SecondaryToolbarModule } from '@sdk-ui/ui/secondary-toolbar/secondary-toolbar.module';
// Icon
import { IconModule } from '@visurel/iconify-angular';
import { SharedModule } from '@shared-ui/shared.module';
import { UserRoutingModule } from './user-routing.module';
import { UserService } from './user.service';
import { UserListComponent } from './user-list/user-list.component';
import { UserFormComponent } from './user-form/user-form.component';
import { UserDetailsComponent } from './user-details/user-details.component';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { UserPermissionItemComponent } from './user-details/user-permission-item/user-permission-item.component';
import { AccessRoleService } from '@management/access-role/access-role.service';
import { CdkHintModule } from '@cdk-ui/hint';
import { MatRadioModule } from '@angular/material/radio';
import { ClusterService } from '@cluster/cluster.service';

@NgModule({
  declarations: [UserListComponent, UserFormComponent, UserDetailsComponent, UserPermissionItemComponent],
  imports: [
    CommonModule,
    ContainerModule,
    PageLayoutModule,
    FlexLayoutModule,
    MatInputModule,
    MatTableModule,
    MatSortModule,
    MatIconModule,
    MatButtonModule,
    MatCheckboxModule,
    MatMenuModule,
    MatTooltipModule,
    MatSnackBarModule,
    MatDialogModule,
    MatProgressSpinnerModule,
    MatSelectModule,
    MatRadioModule,
    IconModule,
    FormsModule,
    ReactiveFormsModule,
    SecondaryToolbarModule,
    SharedModule,
    UserRoutingModule,
    Ng2SearchPipeModule,
    MatSlideToggleModule,
    MatProgressBarModule,
    CdkHintModule
  ],
  providers: [UserService, AccessRoleService, ClusterService]
})
export class UserModule {}
