import { Injectable } from '@angular/core';
import {
  Http,
  URLSearchParams,
  Response,
  RequestOptions,
  RequestOptionsArgs,
  Headers
} from '@angular/http';
import { BehaviorSubject, Observable, Operator } from 'rxjs';
import { map } from 'rxjs/operators';
import { isNullOrUndefined } from 'util';
import { APIService } from './api.service';
import { SessionService } from '../session.service';

@Injectable()
export class ADUserProvider extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  public GetUsers(search: string): Observable<ADUserInfo[]> {
    return this.get(`/btu/${search}`);
  }
}

export class ADUserInfo {
  constructor(public Username: string,
              public Name: string) {}
}
