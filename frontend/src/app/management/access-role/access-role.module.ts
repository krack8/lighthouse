import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AccessRoleRoutingModule } from './access-role-routing.module';
import { AccessRoleFormComponent } from './containers/access-role-form/access-role-form.component';
import { AccessRoleListComponent } from './containers/access-role-list/access-role-list.component';
import { AccessRoleService } from './access-role.service';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatStepperModule } from '@angular/material/stepper';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { MatTooltipModule } from '@angular/material/tooltip';
import { SharedModule } from '@shared-ui/shared.module';
import { IconModule } from '@visurel/iconify-angular';
import { PageLayoutModule } from '@sdk-ui/ui/page-layout/page-layout.module';
import { RoleDeleteConformationComponent } from './common/components/role-delete-conformation/role-delete-conformation.component';
import { MatDialogModule } from '@angular/material/dialog';

@NgModule({
  declarations: [AccessRoleFormComponent, AccessRoleListComponent, RoleDeleteConformationComponent],
  imports: [
    CommonModule,
    AccessRoleRoutingModule,
    ReactiveFormsModule,
    FormsModule,
    MatSnackBarModule,
    MatProgressSpinnerModule,
    MatButtonModule,
    MatIconModule,
    MatFormFieldModule,
    MatInputModule,
    MatExpansionModule,
    MatStepperModule,
    MatCheckboxModule,
    MatProgressBarModule,
    MatTooltipModule,
    MatDialogModule,
    Ng2SearchPipeModule,
    IconModule,
    PageLayoutModule,
    SharedModule
  ],
  providers: [AccessRoleService]
})
export class AccessRoleModule {}
