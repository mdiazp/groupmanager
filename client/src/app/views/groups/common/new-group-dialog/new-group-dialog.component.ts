import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Group } from '../../../../models/core';
import { FormGroupComponent } from '../form-group/form-group.component';

@Component({
  selector: 'app-new-group-dialog',
  templateUrl: './new-group-dialog.component.html',
  styleUrls: ['./new-group-dialog.component.css']
})
export class NewGroupDialogComponent implements OnInit {
  @ViewChild(FormGroupComponent) form: FormGroupComponent;

  constructor() { }

  ngOnInit() {
  }

  private getGroup(): Group {
    return this.form.GetGroup();
  }
}
