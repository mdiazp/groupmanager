import { Routes } from '@angular/router';
import { GroupListComponent } from './group-list/group-list.component';
import { GroupOneComponent } from './group-one/group-one.component';
import { RolAdminGuard } from '../../guards/core';

export const GroupsRoutes: Routes = [
  {
    path: '',
    redirectTo: 'list',
    pathMatch: 'full',
  },
  {
    path: 'list',
    component: GroupListComponent,
    canActivate: [RolAdminGuard],
    canLoad: [RolAdminGuard],
  },
  {
    path: 'showone/:id',
    component: GroupOneComponent,
  }
];
