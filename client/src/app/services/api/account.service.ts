import { Injectable } from '@angular/core';
import {
  Http,
  URLSearchParams,
  Response,
  RequestOptions,
  RequestOptionsArgs,
  Headers,
} from '@angular/http';
import { BehaviorSubject, Observable, Operator } from 'rxjs';
import { map } from 'rxjs/operators';

import {
  Credentials,
  Session,
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';

@Injectable()
export class APIAccountService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  Login(cred: Credentials): Observable<Session> {
    return this.post('/session', cred);
  }
  Logout(): Observable<Response> {
    return this.delete('/session');
  }
  Providers(): Observable<String[]> {
    return this.get('/usersproviders');
  }
}
