import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import {
  ErrorHandlerService,
  APIUserService, UserPublicInfo,
  UserFilter,
  Paginator
} from '../../../services/core';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { MatAutocompleteSelectedEvent } from '@angular/material';

@Component({
  selector: 'app-user-selector',
  templateUrl: './user-selector.component.html',
  styleUrls: ['./user-selector.component.css']
})
export class UserSelectorComponent implements OnInit {

  form: FormGroup;
  usernameControl: FormControl;

  users: UserPublicInfo[] = [];
  autoUserSelection = { ID: 0, Username: '', Name: '' };

  constructor(private api: APIUserService,
              private eh: ErrorHandlerService) { }

  ngOnInit() {
    this.usernameControl = new FormControl('');
    this.form = new FormGroup({
      'username': this.usernameControl,
    });

    this.usernameControl.valueChanges
      .pipe(
        debounceTime(150),
        distinctUntilChanged(),
      )
      .subscribe(
        (value) => {
          if (value !== '') {
            this.loadUsers(value);
          } else {
            this.users = [];
          }
        }
      );
  }

  ValidSelection(): boolean {
    return (this.autoUserSelection.Username !== '' &&
            this.autoUserSelection === this.usernameControl.value);
  }

  Clear(): void {
    this.usernameControl.setValue('');
  }

  loadUsers(value: string): void {
    this.api.GetUsers(
      new UserFilter(value, '', '', '', null, null,
      new Paginator(0, 10), null)
    ).subscribe(
      (data) => {
        this.users = data;
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  onSelectUser(ev: MatAutocompleteSelectedEvent) {
    this.autoUserSelection = ev.option.value;
  }

  displayUser(user?: UserPublicInfo): string | undefined {
    return user ? user.Username : undefined;
  }
}
