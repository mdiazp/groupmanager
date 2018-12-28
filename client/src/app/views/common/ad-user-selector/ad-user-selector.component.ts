import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import {
  ADUserInfo,
  ADUserProvider,
  ErrorHandlerService
} from '../../../services/core';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { MatAutocompleteSelectedEvent } from '@angular/material';

@Component({
  selector: 'app-ad-user-selector',
  templateUrl: './ad-user-selector.component.html',
  styleUrls: ['./ad-user-selector.component.css']
})
export class ADUserSelectorComponent implements OnInit {

  form: FormGroup;
  usernameControl: FormControl;

  @Output() selectionChanges = new EventEmitter<ADUserInfo>();

  users: ADUserInfo[] = [];
  autoUserSelection = new ADUserInfo('', '');

  constructor(private adApi: ADUserProvider,
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

          this.selectionChanges.emit(this.autoUserSelection);
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
    this.adApi.GetUsers(value).subscribe(
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

  displayUser(user?: ADUserInfo): string | undefined {
    return user ? user.Username : undefined;
  }
}
