import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {GroupADUser} from '../../models/core';
import {APIGroupService, Paginator} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';

export class GroupADUsersDataSource implements DataSource<GroupADUser> {

    private groupADUserSubject = new BehaviorSubject<GroupADUser[]>([]);
    private countSubject = new BehaviorSubject<number>(0);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    public count$ = this.countSubject.asObservable();

    constructor(private api: APIGroupService,
                private eh: ErrorHandlerService) {}

    load(loadCount: boolean, groupID: number, paginator: Paginator) {
      this.loadingSubject.next(true);
      if ( loadCount ) {
        this.loadCount(groupID, paginator);
      } else {
        this.loadGroupADUsers(groupID, paginator);
      }
    }

    private loadCount(groupID: number, paginator: Paginator) {
      this.api.GetGroupADUsersCount(groupID).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadGroupADUsers(groupID, paginator);
        },
        (e) => {
          this.countSubject.next(0);
          this.groupADUserSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadGroupADUsers(groupID: number, paginator: Paginator) {
      this.api.GetGroupADUsers(groupID, paginator).subscribe(
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

