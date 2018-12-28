import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, CanLoad } from '@angular/router';
import { Observable } from 'rxjs';
import { Router, Route } from '@angular/router';

import { SessionService } from '../services/session.service';
import { RolAdmin } from '../models/core';

@Injectable()
export class RolAdminGuard implements CanActivate, CanLoad {
  constructor(private session: SessionService,
              private router: Router) {}

  canActivate( next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean> | Promise<boolean> | boolean {
    if (this.can()) {
      return true;
    }
    this.router.navigate(['/login']);
    return false;
  }

  canLoad(route: Route): boolean | Observable<boolean> | Promise<boolean> {
    if (this.can()) {
      return true;
    }
    this.router.navigate(['/login']);
    return false;
  }

  can(): boolean {
    return (this.session.IsOpen() && this.session.GetUser().Rol === RolAdmin);
  }
}
