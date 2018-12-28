import { Component, OnInit, Input, Output, EventEmitter, ViewChild, OnDestroy } from '@angular/core';

import { FormGroupComponent } from '../form-group/form-group.component';
import { Group } from '../../../../models/core';
import { APIGroupService, ErrorHandlerService, FeedbackHandlerService } from '../../../../services/core';


@Component({
  selector: 'app-group-settings',
  templateUrl: './group-settings.component.html',
  styleUrls: ['./group-settings.component.scss']
})
export class GroupSettingsComponent implements OnInit, OnDestroy {

  @Input() group: Group;
  @Output() change = new EventEmitter<boolean>();
  @ViewChild(FormGroupComponent) form: FormGroupComponent;

  constructor(private api: APIGroupService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService) {}

  ngOnInit() {
  }

  onSave(): void {
    const g = this.form.GetGroup();

    this.api.UpdateGroup(g).subscribe(
      (group) => {
        this.group.Name = group.Name;
        this.group.Description = group.Description;
        this.group.Actived = group.Actived;

        this.change.emit(true);

        this.fh.ShowFeedback('Group was updated succesfully');
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  ngOnDestroy(): void {
  }
}
