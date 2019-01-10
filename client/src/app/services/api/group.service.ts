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
  Group, GroupAdmin, GroupADUser
} from '../../models/core';

import { SessionService } from '../session.service';
import { APIService } from './api.service';
import { Paginator, OrderBy } from './util';
import { isNullOrUndefined } from 'util';

@Injectable()
export class APIGroupService extends APIService {
  constructor(protected http: Http,
              protected session: SessionService) {
    super(http, session);
  }

  public PostGroup(group: Group): Observable<Group> {
    return this.post('/group', group);
  }

  public GetGroup(id: number): Observable<Group> {
    return this.get(`/group/${id}`);
  }

  public UpdateGroup(group: Group): Observable<Group> {
    return this.patch(`/group/${group.ID}`, group);
  }

  public DeleteGroup(groupID: number): Observable<Response> {
    return this.delete(`/group/${groupID}`);
  }

  public GetGroups(filter?: GroupFilter): Observable<Group[]> {
    if ( filter && filter !== null ) {
      return this.get('/groups', { params: filter.GetUSP() });
    } else {
      return this.get('/groups');
    }
  }

  public GetGroupsCount(filter?: GroupFilter): Observable<number> {
    if ( filter && filter !== null ) {
      return this.get('/groupscount', { params: filter.GetUSP() });
    } else {
      return this.get('/groupscount');
    }
  }

  public PutGroupAdmin(groupID: number, userID: number): Observable<GroupAdmin> {
    return this.put(`/group/${groupID}/admins/${userID}`, null);
  }

  public DeleteGroupAdmin(groupID: number, userID: number): Observable<Response> {
    return this.delete(`/group/${groupID}/admins/${userID}`, null);
  }

  public GetGroupAdmins(groupID: number, paginator: Paginator): Observable<GroupAdmin[]> {
    return this.get(`/group/${groupID}/admins`, { params: paginator.GetUSP() });
  }

  public GetGroupAdminsCount(groupID: number): Observable<number> {
    return this.get(`/group/${groupID}/adminscount`);
  }

  public PutGroupADUser(groupID: number, aduser: string): Observable<GroupADUser> {
    return this.put(`/group/${groupID}/adusers/${aduser}`, null);
  }

  public DeleteGroupADUser(groupID: number, aduser: string): Observable<Response> {
    return this.delete(`/group/${groupID}/adusers/${aduser}`, null);
  }

  public GetGroupADUsers(groupID: number, filter?: GroupADUserFilter): Observable<GroupADUser[]> {
    if ( filter && filter !== null ) {
      return this.get(`/group/${groupID}/adusers`, { params: filter.GetUSP() });
    } else {
      return this.get(`/group/${groupID}/adusers`);
    }
  }

  public GetGroupADUsersCount(groupID: number, filter?: GroupADUserFilter): Observable<number> {
    if ( filter && filter !== null ) {
      return this.get(`/group/${groupID}/aduserscount`, { params: filter.GetUSP() });
    } else {
      return this.get(`/group/${groupID}/aduserscount`);
    }
  }
}

export class GroupFilter {
  constructor(public NameSubstr: string,
              public AdminID: number,
              public Actived: boolean,
              public ADUser: string,
              public paginator: Paginator,
              public orderby: OrderBy) {}

  public GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if ( !isNullOrUndefined(this.NameSubstr) && this.NameSubstr !== '' ) {
      usp.append('nameSubstr', this.NameSubstr.toString());
    }
    if ( !isNullOrUndefined(this.Actived) ) {
      usp.append('actived', this.Actived.toString());
    }
    if ( !isNullOrUndefined(this.AdminID) && this.AdminID !== 0 ) {
      usp.append('adminID', this.AdminID.toString());
    }
    if ( !isNullOrUndefined(this.ADUser) && this.ADUser !== '' ) {
      usp.append('aduser', this.ADUser.toString());
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

export class GroupADUserFilter {
  constructor(public ADUserPrefix: string,
              public paginator: Paginator) {}

  public GetUSP(): URLSearchParams {
    let usp: URLSearchParams; usp = new URLSearchParams();
    if ( !isNullOrUndefined(this.ADUserPrefix) && this.ADUserPrefix !== '' ) {
      usp.append('adUserPrefix', this.ADUserPrefix);
    }
    if ( !isNullOrUndefined(this.paginator) ) {
      usp.appendAll(this.paginator.GetUSP());
    }
    return usp;
  }
}
