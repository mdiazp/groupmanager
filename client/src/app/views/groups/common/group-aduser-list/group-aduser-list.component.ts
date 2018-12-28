import { Component, OnInit, Input, ViewChild, AfterViewInit, OnDestroy } from '@angular/core';
import { Group, GroupADUser } from '../../../../models/core';
import { GroupADUsersDataSource, APIGroupService, ErrorHandlerService, FeedbackHandlerService, Paginator } from '../../../../services/core';
import { MatPaginator, MatDialog } from '@angular/material';
import { tap } from 'rxjs/operators';
import { CheckDeleteDialogComponent } from '../../../../dialogs/core';
import { isNullOrUndefined } from 'util';
import { ADUserSelectorComponent } from '../../../common/ad-user-selector/ad-user-selector.component';
import { Observable } from 'rxjs';

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


  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(ADUserSelectorComponent) ADUserSelector: ADUserSelectorComponent;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIGroupService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService,
              private dialog: MatDialog) {}

  ngOnInit() {
    this.dataSource = new GroupADUsersDataSource(this.api, this.eh);
    this.dataSource.load(
      true,
      this.group.ID,
      new Paginator(
        0,
        this.initialPageSize,
      ),
    );
  }

  ngAfterViewInit(): void {
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
    this.api.PutGroupADUser(
      this.group.ID,
      this.ADUserSelector.autoUserSelection.Username
    )
    .subscribe(
      (group) => {
        this.fh.ShowFeedback('The user was inserted on group successfully');
        this.ADUserSelector.Clear();
        this.paginator.pageIndex = 0;
        this.load(true);
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  onDelete(aduser: string, adname: string): void {
    const dialogRef = this.dialog.open(CheckDeleteDialogComponent, {
      data: {
        msg: `Are you shure to delete user ${aduser} - ${adname}?`,
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
      new Paginator(
        this.paginator.pageIndex * this.paginator.pageSize,
        this.paginator.pageSize
      ),
    );
  }

  ngOnDestroy(): void {
  }
}
