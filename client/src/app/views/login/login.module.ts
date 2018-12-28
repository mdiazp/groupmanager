import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpModule } from '@angular/http';
import { FlexLayoutModule } from '@angular/flex-layout';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';

import { DemoMaterialModule, } from '../../demo-material-module';
import { LoginRoutes } from './login.routing';
import { LoginComponent } from './login.component';

@NgModule({
  imports: [
    CommonModule,
    HttpModule,
    DemoMaterialModule,
    FlexLayoutModule,
    ReactiveFormsModule,
    RouterModule.forChild(LoginRoutes)
  ],
  declarations: [
    LoginComponent,
  ]
})
export class LoginModule { }
