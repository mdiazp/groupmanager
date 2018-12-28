import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpModule } from '@angular/http';
import { FlexLayoutModule } from '@angular/flex-layout';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { DemoMaterialModule, } from '../../demo-material-module';
import { MyCommonModule } from '../common/common.module';

import { HomeRoutes } from './home.routing';
import { HomeComponent } from './home.component';
// import { FormUserComponent } from './common/form-User/form-User.component';

@NgModule({
  imports: [
    CommonModule,
    HttpModule,
    DemoMaterialModule,
    FlexLayoutModule,
    ReactiveFormsModule,
    FormsModule,
    RouterModule.forChild(HomeRoutes),

    MyCommonModule,
  ],
  declarations: [
    HomeComponent,
  ],
})
export class HomeModule { }
