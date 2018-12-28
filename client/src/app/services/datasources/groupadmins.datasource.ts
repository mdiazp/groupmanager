import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {GroupAdmin} from '../../models/core';
import {APIGroupService, Paginator} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';

export class GroupAdminsDataSource implements DataSource<GroupAdmin> {

    private groupadminsSubject = new BehaviorSubject<GroupAdmin[]>([]);
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
        this.loadGroupAdmins(groupID, paginator);
      }
    }

    private loadCount(groupID: number, paginator: Paginator) {
      this.api.GetGroupAdminsCount(groupID).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadGroupAdmins(groupID, paginator);
        },
        (e) => {
          this.countSubject.next(0);
          this.groupadminsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadGroupAdmins(groupID: number, paginator: Paginator) {
      this.api.GetGroupAdmins(
        groupID,
        paginator,
      ).subscribe(
        groupadmins => this.groupadminsSubject.next(groupadmins),
        (e) => {
          this.groupadminsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    connect(collectionViewer: CollectionViewer): Observable<GroupAdmin[]> {
        return this.groupadminsSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.groupadminsSubject.complete();
        this.loadingSubject.complete();
    }

}

