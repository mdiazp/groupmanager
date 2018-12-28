import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {Group} from '../../models/core';
import {APIGroupService, GroupFilter} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';

export class GroupsDataSource implements DataSource<Group> {

    private groupsSubject = new BehaviorSubject<Group[]>([]);
    private countSubject = new BehaviorSubject<number>(0);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    public count$ = this.countSubject.asObservable();

    constructor(private api: APIGroupService,
                private eh: ErrorHandlerService) {}

    load(loadCount: boolean, filter?: GroupFilter) {
      this.loadingSubject.next(true);
      if ( loadCount ) {
        this.loadCount(filter);
      } else {
        this.loadGroups(filter);
      }
    }

    private loadCount(filter?: GroupFilter) {
      this.api.GetGroupsCount(filter).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadGroups(filter);
        },
        (e) => {
          this.countSubject.next(0);
          this.groupsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadGroups(filter?: GroupFilter) {
      this.api.GetGroups(filter).subscribe(
        groups => this.groupsSubject.next(groups),
        (e) => {
          this.groupsSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    connect(collectionViewer: CollectionViewer): Observable<Group[]> {
        return this.groupsSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.groupsSubject.complete();
        this.loadingSubject.complete();
    }

}

