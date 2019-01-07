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

@Injectable()
export class ADUserProvider {
  private bpath = 'http://api.gm.local:1234/btu';

  constructor(protected http: Http) {}

  public GetUsers(search: string): Observable<ADUserInfo[]> {
    return this.http.get(`${this.bpath}/${search}`)
    .pipe(
      map(res => res.json()),
    );
  }
}

export class ADUserInfo {
  constructor(public Username: string,
              public Name: string) {}
}
