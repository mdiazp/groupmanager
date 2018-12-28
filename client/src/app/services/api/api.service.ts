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

import { SessionService } from '../session.service';
import { isNullOrUndefined } from 'util';

export class APIService {
    protected APIPath = 'http://api.gm.local:1234';
    protected AuthTokenHeader = 'AuthToken';

    constructor(protected http: Http,
                protected session: SessionService) {
    }

    private GetCommonOpts(options?: RequestOptionsArgs): RequestOptionsArgs {
      if ( !options || options === null ) {
        options = new RequestOptions();
      }

      // console.log('va a preguntar');
      if (isNullOrUndefined(options.headers)) {
        // console.log('headers is null');
        options.headers = new Headers();
      }
      options.headers.append('Content-Type', 'application/json');
      options.headers.append('Accept', 'application/json');
      if (this.session.IsOpen()) {
        options.headers.append(this.AuthTokenHeader, this.session.GetToken());
      }
      return options;
    }

    protected get(path: string, options?: RequestOptionsArgs): Observable<any> {
      return this.http.get(
        `${this.APIPath}${path}`, this.GetCommonOpts(options)).pipe(
        map(res => res.json())
      );
    }

    protected post(path: string, body: any, options?: RequestOptionsArgs): Observable<any> {
      return this.http.post(
        `${this.APIPath}${path}`, body, this.GetCommonOpts(options)).pipe(
        map(res => res.json())
      );
    }

    protected put(path: string, body: any, options?: RequestOptionsArgs): Observable<any> {
      return this.http.put(
        `${this.APIPath}${path}`, body, this.GetCommonOpts(options)).pipe(
        map(res => res.json())
      );
    }

    protected patch(path: string, body: any, options?: RequestOptionsArgs): Observable<any> {
      return this.http.patch(
        `${this.APIPath}${path}`, body, this.GetCommonOpts(options)).pipe(
          map(res => res.json())
      );
    }

    protected delete(path: string, options?: RequestOptionsArgs): Observable<Response> {
      return this.http.delete(`${this.APIPath}${path}`, this.GetCommonOpts(options));
    }
}
