import * as $ from 'jquery';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
// import { HttpClientModule, HttpClient } from '@angular/common/http';
import { HttpModule } from '@angular/http';
import { LocationStrategy, PathLocationStrategy } from '@angular/common';
import { AppRoutes } from './app.routing';
import { AppComponent } from './app.component';

import { FlexLayoutModule } from '@angular/flex-layout';
import { FullComponent } from './layouts/full/full.component';
import { AppHeaderComponent } from './layouts/full/header/header.component';
import { AppSidebarComponent } from './layouts/full/sidebar/sidebar.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { DemoMaterialModule } from './demo-material-module';

import { MAT_DIALOG_DEFAULT_OPTIONS } from '@angular/material';
import { CheckDeleteDialogComponent } from './dialogs/core';

import { SharedModule } from './shared/shared.module';
import { SpinnerComponent } from './shared/spinner.component';
// import { LoginComponent } from './login/login.component';

@NgModule({
  declarations: [
    AppComponent,
    FullComponent,
    // LoginComponent,
    AppHeaderComponent,
    SpinnerComponent,
    AppSidebarComponent,

    CheckDeleteDialogComponent,
  ],
  entryComponents: [
    CheckDeleteDialogComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    DemoMaterialModule,
    // FormsModule,
    ReactiveFormsModule,
    FlexLayoutModule,
    // HttpClientModule,
    HttpModule,
    SharedModule,
    RouterModule.forRoot(AppRoutes, { useHash: true })
  ],
  providers: [
    /*{
      provide: LocationStrategy,
      useClass: PathLocationStrategy
    },*/
    {provide: MAT_DIALOG_DEFAULT_OPTIONS, useValue: {
      width: '400px',
      hasBackdrop: true,
      disableClose: true,
    }},
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
