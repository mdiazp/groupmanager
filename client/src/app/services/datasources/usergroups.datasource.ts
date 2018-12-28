import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';

import {GroupAdmin} from '../../models/core';
import {Paginator, APIUserService} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';

export class UserGroupsDataSource implements DataSource<GroupAdmin> {

    private usergroupsSubject = new BehaviorSubject<GroupAdmin[]>([]);
    private countSubject = new BehaviorSubject<number>(0);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    public count$ = this.countSubject.asObservable();

    constructor(private api: APIUserService,
                private eh: ErrorHandlerService) {}

    load(loadCount: boolean, userID: number, paginator: Paginator) {
      this.loadingSubject.next(true);
      if ( loadCount ) {
        this.loadCount(userID, paginator);
      } else {
        this.loadUserGroups(userID, paginator);
      }
    }

    private loadCount(userID: number, paginator: Paginator) {
      this.api.GetUserGroupsCount(userID).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadUserGroups(userID, paginator);
        },
        (e) => {
          this.countSubject.next(0);
          this.usergroupsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadUserGroups(userID: number, paginator: Paginator) {
      this.api.GetUserGroups(
        userID,
        paginator,
      ).subscribe(
        groupadmins => this.usergroupsSubject.next(groupadmins),
        (e) => {
          this.usergroupsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    connect(collectionViewer: CollectionViewer): Observable<GroupAdmin[]> {
        return this.usergroupsSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.usergroupsSubject.complete();
        this.loadingSubject.complete();
    }

}

