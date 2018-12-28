import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ADUserProvider, ADUserInfo } from '../../../../services/api/aduser-provider.service';
import { ErrorHandlerService, UserToPost } from '../../../../services/core';
import { MatAutocompleteSelectedEvent } from '@angular/material';
import { debounceTime, distinctUntilChanged, tap } from 'rxjs/operators';
import { ADUserSelectorComponent } from '../../../common/ad-user-selector/ad-user-selector.component';

@Component({
  selector: 'app-new-user-dialog',
  templateUrl: './new-user-dialog.component.html',
  styleUrls: ['./new-user-dialog.component.css']
})
export class NewUserDialogComponent implements OnInit {

  @ViewChild(ADUserSelectorComponent) userSelector: ADUserSelectorComponent;

  constructor(private adApi: ADUserProvider,
              private eh: ErrorHandlerService) { }

  ngOnInit() {}

  public GetUserToPost(): UserToPost {
    return {
      Username: this.userSelector.autoUserSelection.Username,
      Provider: 'AD',
    };
  }
}
