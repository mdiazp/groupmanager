import { Component, OnInit } from '@angular/core';
import { User } from '../../../models/user';
import { BehaviorSubject } from 'rxjs';
import { ActivatedRoute, Router } from '@angular/router';
import { APIUserService, ErrorHandlerService } from '../../../services/core';

@Component({
  selector: 'app-user-one',
  templateUrl: './user-one.component.html',
  styleUrls: ['./user-one.component.css']
})
export class UserOneComponent implements OnInit {

  userID: number;
  user: User;

  private loadingSubject = new BehaviorSubject<boolean>(true);
  public loading$ = this.loadingSubject.asObservable();


  constructor(private router: Router,
              private route: ActivatedRoute,
              private api: APIUserService,
              private eh: ErrorHandlerService) {
    this.route.params.subscribe(
      params => {
        this.userID = params.id;
        this.loadUser();
      }
    );
  }

  ngOnInit() {
    this.loadUser();
  }

  refresh(): void {
    this.router.navigate(['/', 'users', 'showone', this.user.ID]);
  }

  loadUser(): void {
    this.loadingSubject.next(true);
    this.api.GetUser(this.userID).subscribe(
      (user) => {
        this.user = user;
        this.loadingSubject.next(false);
      },
      (e) => {
        this.router.navigate(['/', 'users', 'list']);
        this.eh.HandleError(e);
      }
    );
  }
}
