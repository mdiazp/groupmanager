import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpModule } from '@angular/http';
import { FlexLayoutModule } from '@angular/flex-layout';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { DemoMaterialModule, } from '../../demo-material-module';
import { UsersRoutes } from './users.routing';
import { UserListComponent } from './user-list/user-list.component';
import { MyCommonModule } from '../common/common.module';

import { NewUserDialogComponent } from './common/new-user-dialog/new-user-dialog.component';
import { UserOneComponent } from './user-one/user-one.component';
import { UserInfoComponent } from './common/user-info/user-info.component';
import { UserSettingsComponent } from './common/user-settings/user-settings.component';
import { UserGroupListComponent } from './common/user-group-list/user-group-list.component';
// import { FormUserComponent } from './common/form-User/form-User.component';

@NgModule({
  imports: [
    CommonModule,
    HttpModule,
    DemoMaterialModule,
    FlexLayoutModule,
    ReactiveFormsModule,
    FormsModule,
    RouterModule.forChild(UsersRoutes),

    MyCommonModule,
  ],
  declarations: [
    UserListComponent,
    NewUserDialogComponent,
    UserOneComponent,
    UserInfoComponent,
    UserSettingsComponent,
    UserGroupListComponent,
    /*
    FormUserComponent
    */
  ],
  entryComponents: [
    NewUserDialogComponent,
  ],
})
export class UsersModule { }
