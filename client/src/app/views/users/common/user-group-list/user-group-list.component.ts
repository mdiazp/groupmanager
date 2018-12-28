import { Component, OnInit, ViewChild, Input, AfterViewInit } from '@angular/core';
import {
  APIUserService,
  ErrorHandlerService,
  FeedbackHandlerService,
  UserGroupsDataSource,
  Paginator
} from '../../../../services/core';
import { MatPaginator } from '@angular/material';
import { User } from '../../../../models/user';
import { GroupAdmin } from '../../../../models/core';
import { tap } from 'rxjs/operators';

@Component({
  selector: 'app-user-group-list',
  templateUrl: './user-group-list.component.html',
  styleUrls: ['./user-group-list.component.css']
})
export class UserGroupListComponent implements OnInit, AfterViewInit {

  @Input() user: User;
  usergroups: GroupAdmin[] = [];
  dataSource: UserGroupsDataSource;
  // displayedColumns= ['id', 'name', 'description', 'actived', 'operations'];
  displayedColumns= ['groupID', 'groupname'];


  @ViewChild(MatPaginator) paginator: MatPaginator;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIUserService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService) {}

  ngOnInit() {
      this.dataSource = new UserGroupsDataSource(this.api, this.eh);
      this.dataSource.load(
        true,
        this.user.ID,
        new Paginator(
          0,
          this.initialPageSize,
        )
      );
  }

  ngAfterViewInit(): void {
    this.paginator.page
      .pipe(
        tap(() => this.load(false)),
      )
      .subscribe();
  }

  load(loadCount: boolean) {
    this.dataSource.load(
      loadCount,
      this.user.ID,
      new Paginator(
        this.paginator.pageIndex * this.paginator.pageSize,
        this.paginator.pageSize
      ),
    );
  }

}
