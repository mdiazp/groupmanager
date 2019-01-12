import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable, BehaviorSubject } from 'rxjs';

import { Group, PermissionUpdateGroup, PermissionRetrieveGroupAdmin } from '../../../models/core';
import { APIGroupService, ErrorHandlerService, SessionService } from '../../../services/core';

@Component({
  selector: 'app-group-one',
  templateUrl: './group-one.component.html',
  styleUrls: ['./group-one.component.scss']
})
export class GroupOneComponent implements OnInit, OnDestroy {

  groupID: number;
  group: Group;

  private loadingSubject = new BehaviorSubject<boolean>(true);
  loading$ = this.loadingSubject.asObservable();

  private groupSubject = new BehaviorSubject<Group>(null);
  group$ = this.groupSubject.asObservable();

  constructor(private router: Router,
              private route: ActivatedRoute,
              private session: SessionService,
              private api: APIGroupService,
              private eh: ErrorHandlerService) {
    this.route.params.subscribe(
      params => {
        this.groupID = params.id;
        this.loadGroup();
      }
    );
  }

  ngOnInit() {
  }

  loadGroup(): void {
    this.loadingSubject.next(true);
    this.api.GetGroup(this.groupID).subscribe(
      (group) => {
        this.group = group;
        this.groupSubject.next(group);
        this.loadingSubject.next(false);
      },
      (e) => {
        this.router.navigate(['/', 'groups', 'list']);
        this.eh.HandleError(e);
      }
    );
  }

  enableSettings(): boolean {
    return this.session.HasPermission(PermissionUpdateGroup);
  }

  enableAdminList(): boolean {
    return this.session.HasPermission(PermissionRetrieveGroupAdmin);
  }

  ngOnDestroy(): void {
  }
}
