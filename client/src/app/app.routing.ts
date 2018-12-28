import { Routes } from '@angular/router';

import { FullComponent } from './layouts/full/full.component';
import { AuthGuard, RolAdminGuard } from './guards/core';
// import { LoginComponent } from './login/login.component';

export const AppRoutes: Routes = [
  {
    path: 'login',
    loadChildren: './views/login/login.module#LoginModule'
  },
  {
    path: '',
    component: FullComponent,
    canActivate: [AuthGuard],
    canLoad: [AuthGuard],
    children: [
      {
        path: '',
        redirectTo: '/home',
        pathMatch: 'full'
      },
      {
        path: 'home',
        loadChildren: './views/home/home.module#HomeModule'
      },
      {
        path: 'users',
        canActivate: [RolAdminGuard],
        canLoad: [RolAdminGuard],
        loadChildren: './views/users/users.module#UsersModule'
      },
      {
        path: 'groups',
        loadChildren: './views/groups/groups.module#GroupsModule'
      }
    ]
  }
];
