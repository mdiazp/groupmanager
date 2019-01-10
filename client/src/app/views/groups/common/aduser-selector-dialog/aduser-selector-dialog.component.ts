import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ADUserProvider, ADUserInfo } from '../../../../services/api/aduser-provider.service';
import { ErrorHandlerService, UserToPost } from '../../../../services/core';
import { MatAutocompleteSelectedEvent } from '@angular/material';
import { debounceTime, distinctUntilChanged, tap } from 'rxjs/operators';
import { ADUserSelectorComponent } from '../../../common/ad-user-selector/ad-user-selector.component';

@Component({
  selector: 'app-aduser-selector-dialog',
  templateUrl: './aduser-selector-dialog.component.html',
  styleUrls: ['./aduser-selector-dialog.component.css']
})
export class ADUserSelectorDialogComponent implements OnInit {

  @ViewChild(ADUserSelectorComponent) userSelector: ADUserSelectorComponent;

  constructor() { }

  ngOnInit() {}

  public GetADUser(): string {
    return this.userSelector.autoUserSelection.Username;
  }
}
