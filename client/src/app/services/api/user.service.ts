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
  User, GroupAdmin,
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';
import { Paginator, OrderBy } from './util';
import { isNullOrUndefined } from 'util';

@Injectable()
export class APIUserService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  public PostUser(user: UserToPost): Observable<User> {
    return this.post('/user', user);
  }

  public GetUser(id: number): Observable<User> {
    return this.get(`/user/${id}`);
  }

  public PatchUser(userID: number, user: UserToEdit): Observable<User> {
    return this.patch(`/user/${userID}`, user);
  }

  public DeleteUser(userID: number): Observable<Response> {
    return this.delete(`/user/${userID}`);
  }

  public GetUsers(filter?: UserFilter): Observable<User[]> {
    if ( filter && filter !== null ) {
      return this.get('/users', { params: filter.GetUSP() });
    } else {
      return this.get('/users');
    }
  }

  public GetUsersCount(filter?: UserFilter): Observable<number> {
    if ( filter && filter !== null ) {
      return this.get('/userscount', { params: filter.GetUSP() });
    } else {
      return this.get('/userscount');
    }
  }

  public GetUserGroups(userID: number, paginator: Paginator): Observable<GroupAdmin[]> {
    return this.get(`/user/${userID}/groups`, { params: paginator.GetUSP() });
  }

  public GetUserGroupsCount(userID: number): Observable<number> {
    return this.get(`/user/${userID}/groupscount`);
  }
}

export interface UserPublicInfo {
  ID: number;
  Username: string;
  Name: string;
}

export interface UserToPost {
  Username: string;
  Provider: string;
}

export interface UserToEdit {
  Rol: string;
  Enabled: boolean;
}

export class UserFilter {
  constructor(public UsernamePrefix: string,
              public Provider: string,
              public Rol: string,
              public NameSubstr: string,
              public Enabled: boolean,
              public GroupWhichAdmin: number,
              public paginator: Paginator,
              public orderby: OrderBy) {}

  public GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if ( !isNullOrUndefined(this.UsernamePrefix) && this.UsernamePrefix !== '' ) {
      usp.append('usernamePrefix', this.UsernamePrefix.toString());
    }
    if ( !isNullOrUndefined(this.Provider) && this.Provider !== '' ) {
      usp.append('provider', this.Provider.toString());
    }
    if ( !isNullOrUndefined(this.Rol) && this.Rol !== '' ) {
      usp.append('rol', this.Rol.toString());
    }
    if ( !isNullOrUndefined(this.NameSubstr) && this.NameSubstr !== '' ) {
      usp.append('nameSubstr', this.NameSubstr.toString());
    }
    if ( !isNullOrUndefined(this.Enabled) ) {
      usp.append('enabled', this.Enabled.toString());
    }
    if ( !isNullOrUndefined(this.GroupWhichAdmin) && this.GroupWhichAdmin !== 0 ) {
      usp.append('groupWhichAdmin', this.GroupWhichAdmin.toString());
    }
    if ( !isNullOrUndefined(this.paginator) ) {
      usp.appendAll(this.paginator.GetUSP());
    }
    if ( !isNullOrUndefined(this.orderby) ) {
      usp.appendAll(this.orderby.GetUSP());
    }
    return usp;
  }
}
