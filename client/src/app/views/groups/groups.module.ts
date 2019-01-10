import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpModule } from '@angular/http';
import { FlexLayoutModule } from '@angular/flex-layout';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { DemoMaterialModule, } from '../../demo-material-module';
import { MyCommonModule } from '../common/common.module';
import { GroupsRoutes } from './groups.routing';
import { GroupListComponent } from './group-list/group-list.component';
import { GroupOneComponent } from './group-one/group-one.component';
import { GroupInfoComponent } from './common/group-info/group-info.component';
import { GroupSettingsComponent } from './common/group-settings/group-settings.component';
import { NewGroupDialogComponent } from './common/new-group-dialog/new-group-dialog.component';
import { FormGroupComponent } from './common/form-group/form-group.component';
import { GroupAdminListComponent } from './common/group-admin-list/group-admin-list.component';
import { GroupADUserListComponent } from './common/group-aduser-list/group-aduser-list.component';
import { ADUserSelectorDialogComponent } from './common/aduser-selector-dialog/aduser-selector-dialog.component';
import { AdminSelectorDialogComponent } from './common/admin-selector-dialog/admin-selector-dialog.component';

@NgModule({
  imports: [
    CommonModule,
    HttpModule,
    DemoMaterialModule,
    FlexLayoutModule,
    ReactiveFormsModule,
    FormsModule,
    RouterModule.forChild(GroupsRoutes),
    MyCommonModule,
  ],
  declarations: [
    GroupListComponent,
    GroupOneComponent,
    GroupInfoComponent,
    GroupSettingsComponent,
    NewGroupDialogComponent,
    FormGroupComponent,
    GroupAdminListComponent,
    GroupADUserListComponent,
    ADUserSelectorDialogComponent,
    AdminSelectorDialogComponent,
  ],
  entryComponents: [
    NewGroupDialogComponent,
    ADUserSelectorDialogComponent,
    AdminSelectorDialogComponent,
  ],
})
export class GroupsModule { }
