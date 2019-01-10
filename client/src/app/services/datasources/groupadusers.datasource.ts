import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {GroupADUser} from '../../models/core';
import {APIGroupService, Paginator, GroupADUserFilter} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';

export class GroupADUsersDataSource implements DataSource<GroupADUser> {

    private groupADUserSubject = new BehaviorSubject<GroupADUser[]>([]);
    private countSubject = new BehaviorSubject<number>(0);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    public count$ = this.countSubject.asObservable();

    constructor(private api: APIGroupService,
                private eh: ErrorHandlerService) {}

    load(loadCount: boolean, groupID: number, filter: GroupADUserFilter) {
      this.loadingSubject.next(true);
      if ( loadCount ) {
        this.loadCount(groupID, filter);
      } else {
        this.loadGroupADUsers(groupID, filter);
      }
    }

    private loadCount(groupID: number, filter: GroupADUserFilter) {
      this.api.GetGroupADUsersCount(groupID, filter).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadGroupADUsers(groupID, filter);
        },
        (e) => {
          this.countSubject.next(0);
          this.groupADUserSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadGroupADUsers(groupID: number, filter: GroupADUserFilter) {
      this.api.GetGroupADUsers(groupID, filter).subscribe(
        groupADUsers => this.groupADUserSubject.next(groupADUsers),
        (e) => {
          this.groupADUserSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    connect(collectionViewer: CollectionViewer): Observable<GroupADUser[]> {
        return this.groupADUserSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.groupADUserSubject.complete();
        this.loadingSubject.complete();
    }

}

