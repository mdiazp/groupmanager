import { Component, OnInit, ViewChild } from '@angular/core';
import { UserSelectorComponent } from '../../../common/user-selector/user-selector.component';

@Component({
  selector: 'app-admin-selector-dialog',
  templateUrl: './admin-selector-dialog.component.html',
  styleUrls: ['./admin-selector-dialog.component.css']
})
export class AdminSelectorDialogComponent implements OnInit {

  @ViewChild(UserSelectorComponent) userSelector: UserSelectorComponent;

  constructor() { }

  ngOnInit() {}

  public GetUser(): number {
    return this.userSelector.autoUserSelection.ID;
  }

}
