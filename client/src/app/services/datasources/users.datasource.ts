import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import {Observable} from 'rxjs/Observable';
import {BehaviorSubject} from 'rxjs/BehaviorSubject';
import {catchError, finalize} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

import {User} from '../../models/core';
import {APIUserService, UserFilter} from '../api/core';
import {ErrorHandlerService} from '../error-handler.service';

export class UsersDataSource implements DataSource<User> {

    private usersSubject = new BehaviorSubject<User[]>([]);
    private countSubject = new BehaviorSubject<number>(0);

    private loadingSubject = new BehaviorSubject<boolean>(false);

    public loading$ = this.loadingSubject.asObservable();

    public count$ = this.countSubject.asObservable();

    constructor(private api: APIUserService,
                private eh: ErrorHandlerService) {}

    load(loadCount: boolean, filter?: UserFilter) {
      this.loadingSubject.next(true);
      if ( loadCount ) {
        this.loadCount(filter);
      } else {
        this.loadUsers(filter);
      }
    }

    private loadCount(filter?: UserFilter) {
      this.api.GetUsersCount(filter).subscribe(
        count => {
          this.countSubject.next(count);
          this.loadUsers(filter);
        },
        (e) => {
          this.countSubject.next(0);
          this.usersSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    private loadUsers(filter?: UserFilter) {
      this.api.GetUsers(filter).subscribe(
        xusers => this.usersSubject.next(xusers),
        (e) => {
          this.usersSubject.next([]);
          this.eh.HandleError(e);
        }
      );
    }

    connect(collectionViewer: CollectionViewer): Observable<User[]> {
        return this.usersSubject.asObservable();
    }

    disconnect(collectionViewer: CollectionViewer): void {
        this.usersSubject.complete();
        this.loadingSubject.complete();
    }

}
