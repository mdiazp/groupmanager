import { Component, OnInit, Input, ViewChild, AfterViewInit } from '@angular/core';
import { Group, GroupAdmin } from '../../../../models/core';
import { GroupAdminsDataSource, APIGroupService, ErrorHandlerService, FeedbackHandlerService, Paginator } from '../../../../services/core';
import { MatPaginator, MatDialog } from '@angular/material';
import { tap } from 'rxjs/operators';
import { CheckDeleteDialogComponent } from '../../../../dialogs/core';
import { isNullOrUndefined } from 'util';
import { UserSelectorComponent } from '../../../common/user-selector/user-selector.component';

@Component({
  selector: 'app-group-admin-list',
  templateUrl: './group-admin-list.component.html',
  styleUrls: ['./group-admin-list.component.css']
})
export class GroupAdminListComponent implements OnInit, AfterViewInit {

  @Input() group: Group;
  groupadmins: GroupAdmin[] = [];
  dataSource: GroupAdminsDataSource;
  // displayedColumns= ['id', 'name', 'description', 'actived', 'operations'];
  displayedColumns= ['username', 'operations'];


  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(UserSelectorComponent) userSelector: UserSelectorComponent;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIGroupService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService,
              private dialog: MatDialog) {}

  ngOnInit() {
      this.dataSource = new GroupAdminsDataSource(this.api, this.eh);
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
        tap(() => this.load(false)),
      )
      .subscribe();
  }

  onNew(): void {
    console.log('onNew');
    this.api.PutGroupAdmin(
      this.group.ID,
      this.userSelector.autoUserSelection.ID
    )
    .subscribe(
      (group) => {
        console.log('New admin was inserted successfully');
        this.fh.ShowFeedback('New admin was inserted successfully');
        this.paginator.pageIndex = 0;
        this.load(true);

        this.userSelector.Clear();
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  onDelete(userID: number, username: string): void {
    const dialogRef = this.dialog.open(CheckDeleteDialogComponent, {
      data: {
        msg: `Are you shure to delete admin with name ${username}?`,
      },
    });

    dialogRef.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) && result === true ) {
        this.api.DeleteGroupAdmin(this.group.ID, userID).subscribe(
          (_) => {
            console.log('Admin was deleted successfully');
            this.fh.ShowFeedback('Admin was deleted succesfully');
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
}
