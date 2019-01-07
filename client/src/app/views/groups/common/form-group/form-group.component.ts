import { Component, OnInit, Input } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Group } from '../../../../models/core';
import { isNullOrUndefined } from 'util';

@Component({
  selector: 'app-form-group',
  templateUrl: './form-group.component.html',
  styleUrls: ['./form-group.component.css']
})
export class FormGroupComponent implements OnInit {

  @Input() initialValues = new Group(0, '', '', false);

  nameControl: FormControl;
  descriptionControl: FormControl;
  activedControl: FormControl;
  form: FormGroup;

  constructor() {}

  ngOnInit() {
    this.initForm();
  }

  initForm(): void {
    this.nameControl = new FormControl(
      this.initialValues.Name,
      [Validators.required, Validators.maxLength(20)]
    );
    this.descriptionControl = new FormControl(
      this.initialValues.Description,
      [Validators.required, Validators.maxLength(500)],
    );
    this.activedControl = new FormControl(this.initialValues.Actived, Validators.required);
    this.form = new FormGroup({
      'name': this.nameControl,
      'description': this.descriptionControl,
      'actived': this.activedControl,
    });
  }

  public Valid(): boolean {
    return this.form.valid;
  }

  public ResetValues(): void {
    this.nameControl.setValue(this.initialValues.Name);
    this.descriptionControl.setValue(this.initialValues.Description);
    this.activedControl.setValue(this.initialValues.Actived);
  }

  public GetGroup(): Group {
    return new Group(
      this.initialValues.ID,
      this.nameControl.value,
      this.descriptionControl.value,
      this.activedControl.value
    );
  }
}
