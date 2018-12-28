import { Component, ViewChild, OnInit, AfterViewInit, ElementRef } from '@angular/core';
import { MatPaginator, MatSort, MatSelect } from '@angular/material';
import { tap, debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { fromEvent } from 'rxjs';
import { MatDialog } from '@angular/material';

import { CheckDeleteDialogComponent } from '../../../dialogs/core';

import {
  APIUserService,
  UserFilter,
  UsersDataSource,
  ErrorHandlerService,
  FeedbackHandlerService,
  Paginator,
  OrderBy,
} from '../../../services/core';
import { User } from '../../../models/core';
import { isNullOrUndefined } from 'util';
import { NewUserDialogComponent } from '../common/new-user-dialog/new-user-dialog.component';
import { Router } from '@angular/router';
// import { NewUserDialogComponent } from '../common/new-User-dialog/new-User-dialog.component';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss']
})
export class UserListComponent implements OnInit, AfterViewInit {

  users: User[] = [];
  dataSource: UsersDataSource;
  displayedColumns= ['id', 'username', 'name', 'rol', 'operations'];

  @ViewChild(MatSort) sort: MatSort;
  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild('usernameFilter') usernameFilter: ElementRef;
  @ViewChild('nameFilter') nameFilter: ElementRef;
  @ViewChild('rolFilter') rolFilter: MatSelect;
  @ViewChild('enabledFilter') enabledFilter: MatSelect;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIUserService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService,
              private router: Router,
              private dialog: MatDialog) {}

  ngOnInit() {
      this.dataSource = new UsersDataSource(this.api, this.eh);
      this.dataSource.load(
        true,
        new UserFilter(
          null, null, null, null, null, null,
          new Paginator(
            0,
            this.initialPageSize,
          ),
          new OrderBy(
            'id',
            false,
          )
        ),
      );
  }

  ngAfterViewInit() {
    fromEvent(this.usernameFilter.nativeElement, 'keyup')
            .pipe(
                debounceTime(150),
                distinctUntilChanged(),
                tap(() => {
                    this.paginator.pageIndex = 0;
                    // this.paginator.page.emit();
                    this.load(true);
                })
            )
            .subscribe();

    fromEvent(this.nameFilter.nativeElement, 'keyup')
            .pipe(
                debounceTime(150),
                distinctUntilChanged(),
                tap(() => {
                    this.paginator.pageIndex = 0;
                    // this.paginator.page.emit();
                    this.load(true);
                })
            )
            .subscribe();

    this.rolFilter.valueChange.subscribe(
      () => {
        this.paginator.pageIndex = 0;
        // this.paginator.page.emit();
        this.load(true);
      },
    );

    this.enabledFilter.valueChange.subscribe(
      () => {
        this.paginator.pageIndex = 0;
        // this.paginator.page.emit();
        this.load(true);
      },
    );

    this.sort.sortChange.subscribe(
      () => {
        this.paginator.pageIndex = 0;
        // this.paginator.page.emit();
        this.load(false);
      }
    );

    this.paginator.page
      .pipe(
        tap(() => this.load(false)),
      )
      .subscribe();
  }

  onNew(): void {
    console.log('onNew');

    const ref = this.dialog.open(NewUserDialogComponent);

    ref.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) ) {
        this.api.PostUser(result).subscribe(
          (user) => {
            console.log('User was registered successfully');
            this.fh.ShowFeedback('User was registered successfully');
            // this.load(true);
            this.router.navigate(['/', 'users', 'showone', user.ID]);
          },
          (e) => {
            this.eh.HandleError(e);
          }
        );
      }
    });
  }

  onDelete(userID: number, username: string): void {
    const dialogRef = this.dialog.open(CheckDeleteDialogComponent, {
      data: {
        msg: `Are you shure to delete User with name ${username}?`,
      },
    });

    dialogRef.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) && result === true ) {
        this.api.DeleteUser(userID).subscribe(
          (_) => {
            console.log('The element was deleted successfully');
            this.fh.ShowFeedback('User was deleted succesfully');
            this.paginator.pageIndex = 0;
            // this.paginator.page.emit();
            this.load(true);
          },
          (e) => {
            this.eh.HandleError(e);
          }
        );
      }
    });
  }

  load(loadCount: boolean) {
    this.dataSource.load(
      loadCount,

      new UserFilter(
        this.usernameFilter.nativeElement.value,
        null,
        (this.rolFilter.value === 'all' ? null : this.rolFilter.value),
        this.nameFilter.nativeElement.value,
        (this.enabledFilter.value === 'all' ? null : this.enabledFilter.value),
        null,
        new Paginator(
          this.paginator.pageIndex * this.paginator.pageSize,
          this.paginator.pageSize
        ),
        new OrderBy(
          this.sort.active,
          (this.sort.direction !== 'asc'),
        ),
      ),
    );
  }
}
