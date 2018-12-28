import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { User } from '../../../../models/user';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { APIUserService, ErrorHandlerService, FeedbackHandlerService } from '../../../../services/core';

@Component({
  selector: 'app-user-settings',
  templateUrl: './user-settings.component.html',
  styleUrls: ['./user-settings.component.css']
})
export class UserSettingsComponent implements OnInit {

  @Input() user: User;
  @Output() change = new EventEmitter<boolean>();

  form: FormGroup;
  rolControl: FormControl;
  enabledControl: FormControl;

  constructor(private api: APIUserService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService) { }

  ngOnInit() {
    this.rolControl = new FormControl(this.user.Rol, Validators.required);
    this.enabledControl = new FormControl(this.user.Enabled, Validators.required);
    this.form = new FormGroup({
      'rol': this.rolControl,
      'enabled': this.enabledControl,
    });
  }

  public ResetValues(): void {
    this.rolControl.setValue(this.user.Rol);
    this.enabledControl.setValue(this.user.Enabled);
  }

  onSave(): void {
    const ue = {
      Rol: this.rolControl.value,
      Enabled: this.enabledControl.value,
    };

    console.log('onSave()');
    this.api.PatchUser(this.user.ID, ue).subscribe(
      (user) => {
        this.user.ID = user.ID;
        this.user.Provider = user.Provider;
        this.user.Username = user.Username;
        this.user.Name = user.Name;
        this.user.Rol = user.Rol;
        this.user.Enabled = user.Enabled;

        this.change.emit(true);

        this.fh.ShowFeedback('User was updated succesfully');
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }
}
