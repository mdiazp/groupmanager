import { Component, OnInit, Input, ViewChild, AfterViewInit, OnDestroy, ElementRef } from '@angular/core';
import { Group, GroupADUser } from '../../../../models/core';
import {
  GroupADUsersDataSource,
  APIGroupService,
  ErrorHandlerService,
  FeedbackHandlerService,
  Paginator,
  GroupADUserFilter
} from '../../../../services/core';
import { MatPaginator, MatDialog } from '@angular/material';
import { tap, debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { CheckDeleteDialogComponent } from '../../../../dialogs/core';
import { isNullOrUndefined } from 'util';
import { ADUserSelectorComponent } from '../../../common/ad-user-selector/ad-user-selector.component';
import { Observable, fromEvent } from 'rxjs';
import { ADUserSelectorDialogComponent } from '../aduser-selector-dialog/aduser-selector-dialog.component';

@Component({
  selector: 'app-group-aduser-list',
  templateUrl: './group-aduser-list.component.html',
  styleUrls: ['./group-aduser-list.component.css']
})
export class GroupADUserListComponent implements OnInit, AfterViewInit, OnDestroy {

  @Input() group: Group;
  groupADUsers: GroupADUser[] = [];
  dataSource: GroupADUsersDataSource;
  // displayedColumns= ['id', 'name', 'description', 'actived', 'operations'];
  displayedColumns= ['aduser', 'adname', 'operations'];

  @ViewChild('adUserPrefixFilter') adUserPrefixFilter: ElementRef;
  @ViewChild(MatPaginator) paginator: MatPaginator;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIGroupService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService,
              private dialog: MatDialog) {}

  ngOnInit() {
    this.dataSource = new GroupADUsersDataSource(this.api, this.eh);
    console.log('ngOnInit()');
    this.dataSource.load(
      true,
      this.group.ID,
      new GroupADUserFilter(
        null,
        new Paginator(
          0,
          this.initialPageSize,
        ),
      ),
    );
  }

  ngAfterViewInit(): void {
    fromEvent(this.adUserPrefixFilter.nativeElement, 'keyup')
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

    this.paginator.page
      .pipe(
        tap(
          () => {
            this.load(false);
          },
        ),
      )
      .subscribe();
  }

  onNew(): void {
    const ref = this.dialog.open(ADUserSelectorDialogComponent);

    ref.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) && result !== '' ) {
        this.api.PutGroupADUser(
          this.group.ID,
          result,
        )
        .subscribe(
          (group) => {
            this.fh.ShowFeedback('The user was inserted on group successfully');
            this.paginator.pageIndex = 0;
            this.load(true);
          },
          (e) => {
            this.eh.HandleError(e);
          }
        );
      }
    });
  }

  onDelete(aduser: string, adname: string): void {
    const dialogRef = this.dialog.open(CheckDeleteDialogComponent, {
      data: {
        msg: `Are you shure to remove user ${aduser} from this group?`,
      },
    });

    dialogRef.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) && result === true ) {
        this.api.DeleteGroupADUser(this.group.ID, aduser).subscribe(
          (_) => {
            this.fh.ShowFeedback('User was removed from group succesfully');
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
      this.group.ID,
      new GroupADUserFilter(
        this.adUserPrefixFilter.nativeElement.value,
        new Paginator(
          this.paginator.pageIndex * this.paginator.pageSize,
          this.paginator.pageSize
        ),
      ),
    );
  }

  ngOnDestroy(): void {
  }
}
