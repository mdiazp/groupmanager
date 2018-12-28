import { Routes } from '@angular/router';
import { UserListComponent } from './user-list/user-list.component';
import { UserOneComponent } from './user-one/user-one.component';

export const UsersRoutes: Routes = [
  {
    path: '',
    redirectTo: 'list',
    pathMatch: 'full',
  },
  {
    path: 'list',
    component: UserListComponent,
  },
  {
    path: 'showone/:id',
    component: UserOneComponent,
  }
];
