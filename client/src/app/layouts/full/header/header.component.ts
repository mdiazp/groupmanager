import { Component } from '@angular/core';
import { Http } from '@angular/http';
import {
  SessionService,
  ErrorHandlerService,
  APIAccountService,
} from '../../../services/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: []
})
export class AppHeaderComponent {
  constructor(protected http: Http,
              protected router: Router,
              public session: SessionService,
              protected eh: ErrorHandlerService,
              protected api: APIAccountService) {}

  onLogout(): void {
    this.api.Logout().subscribe(
      (_) => {
        this.session.Close();
        this.router.navigate(['/login']);
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }
}
