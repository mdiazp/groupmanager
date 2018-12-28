import { Component, OnInit, AfterViewInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormControl, FormGroup, Validators } from '@angular/forms';

import {
  APIAccountService,
  SessionService,
  ErrorHandlerService,
} from '../../services/core';
import { Credentials } from '../../models/core';

/** @title Responsive sidenav */
@Component({
  selector: 'app-login',
  templateUrl: 'login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup;
  username: FormControl;
  password: FormControl;
  provider: FormControl;

  loading = true;
  providers: String[] = [];

  constructor(private api: APIAccountService,
              private session: SessionService,
              private eh: ErrorHandlerService,
              private router: Router) {
    this.loading = true;
  }

  ngOnInit() {
    this.initForm();
    this.loadProviders();
  }

  initForm() {
    this.username = new FormControl('', Validators.required);
    this.password = new FormControl('', Validators.required);
    this.provider = new FormControl('', Validators.required);

    this.loginForm = new FormGroup({
      'username': this.username,
      'password': this.password,
      'provider': this.provider,
    });
  }

  onLogin() {
    this.api.Login(
      new Credentials(this.username.value, this.password.value, this.provider.value),
    ).subscribe(
      (session) => {
        this.session.Open(session);
        this.router.navigate(['']);
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  loadProviders(): void {
    this.api.Providers().subscribe(
      (providers) => {
        this.providers = providers;
        if (this.providers.length > 0) {
          this.provider.setValue(this.providers[0]);
        }
        this.loading = false;
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }
}
