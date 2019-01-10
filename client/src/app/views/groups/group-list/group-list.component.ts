import { Component, ViewChild, OnInit, AfterViewInit, ElementRef } from '@angular/core';
import { MatPaginator, MatSort, MatSelect } from '@angular/material';
import { tap, debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { fromEvent } from 'rxjs';
import { MatDialog } from '@angular/material';

import { CheckDeleteDialogComponent } from '../../../dialogs/core';

import {
  APIGroupService,
  GroupFilter,
  GroupsDataSource,
  ErrorHandlerService,
  FeedbackHandlerService,
  Paginator,
  OrderBy,
} from '../../../services/core';
import { Group } from '../../../models/core';
import { isNullOrUndefined } from 'util';
import { NewGroupDialogComponent } from '../common/new-group-dialog/new-group-dialog.component';
import { Router } from '@angular/router';
import { ADUserSelectorComponent } from '../../common/ad-user-selector/ad-user-selector.component';

@Component({
  selector: 'app-group-list',
  templateUrl: './group-list.component.html',
  styleUrls: ['./group-list.component.scss']
})
export class GroupListComponent implements OnInit, AfterViewInit {

  groups: Group[] = [];
  dataSource: GroupsDataSource;
  // displayedColumns= ['id', 'name', 'description', 'actived', 'operations'];
  displayedColumns= ['name', 'operations'];


  @ViewChild(MatSort) sort: MatSort;
  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild('nameFilter') nameFilter: ElementRef;
  @ViewChild(MatSelect) activedFilter: MatSelect;
  @ViewChild(ADUserSelectorComponent) aduserSelector: ADUserSelectorComponent;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  constructor(private api: APIGroupService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService,
              private router: Router,
              private dialog: MatDialog) {}

  ngOnInit() {
      this.dataSource = new GroupsDataSource(this.api, this.eh);
      this.dataSource.load(
        true,
        new GroupFilter(
          null, null, null, null,
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

    this.activedFilter.valueChange.subscribe(
      () => {
        this.paginator.pageIndex = 0;
        // this.paginator.page.emit();
        this.load(true);
      },
    );

    this.aduserSelector.selectionChanges.subscribe(
      (_) => {
        this.paginator.pageIndex = 0;
        this.load(true);
      }
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
    const ref = this.dialog.open(NewGroupDialogComponent);

    ref.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) ) {
        this.api.PostGroup(result).subscribe(
          (group) => {
            console.log('The element was created successfully');
            this.fh.ShowFeedback('The element was created successfully');
            this.router.navigate(['/', 'groups', 'showone', group.ID]);
          },
          (e) => {
            this.eh.HandleError(e);
          }
        );
      }
    });
  }

  onDelete(groupID: number, groupName: string): void {
    const dialogRef = this.dialog.open(CheckDeleteDialogComponent, {
      data: {
        msg: `Are you shure to delete group with name ${groupName}?`,
      },
    });

    dialogRef.afterClosed().subscribe(result => {
      if ( !isNullOrUndefined(result) && result === true ) {
        this.api.DeleteGroup(groupID).subscribe(
          (_) => {
            console.log('The element was deleted successfully');
            this.fh.ShowFeedback('Group was deleted succesfully');
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
    let aduser: string; aduser = '';
    if ( this.aduserSelector.ValidSelection() ) {
      aduser = this.aduserSelector.autoUserSelection.Username;
    }

    this.dataSource.load(
      loadCount,
      new GroupFilter(
        this.nameFilter.nativeElement.value,
        null,
        (this.activedFilter.value === 'all' ? null : this.activedFilter.value),
        aduser,
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
