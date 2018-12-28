import { Component, OnInit, Input } from '@angular/core';
import { Group } from '../../../../models/core';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-group-info',
  templateUrl: './group-info.component.html',
  styleUrls: ['./group-info.component.css']
})
export class GroupInfoComponent implements OnInit {

  @Input() group: Group;

  constructor() {}

  ngOnInit() {}

}
