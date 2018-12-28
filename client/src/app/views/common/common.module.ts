import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpModule } from '@angular/http';
import { FlexLayoutModule } from '@angular/flex-layout';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { DemoMaterialModule, } from '../../demo-material-module';
import { ADUserSelectorComponent } from './ad-user-selector/ad-user-selector.component';
import { UserSelectorComponent } from './user-selector/user-selector.component';
// import { UserOneComponent } from './user-one/user-one.component';
// import { UserInfoComponent } from './common/user-info/user-info.component';
// import { UserSettingsComponent } from './common/user-settings/user-settings.component';
// import { NewUserDialogComponent } from './common/new-User-dialog/new-User-dialog.component';
// import { FormUserComponent } from './common/form-User/form-User.component';

@NgModule({
  imports: [
    CommonModule,
    HttpModule,
    DemoMaterialModule,
    FlexLayoutModule,
    ReactiveFormsModule,
    FormsModule,
  ],
  exports: [
    ADUserSelectorComponent,
    UserSelectorComponent,
  ],
  declarations: [ADUserSelectorComponent, UserSelectorComponent],
  entryComponents: [],
})
export class MyCommonModule { }
