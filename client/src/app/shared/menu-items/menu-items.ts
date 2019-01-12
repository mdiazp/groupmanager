import { Injectable } from '@angular/core';
import { SessionService } from '../../services/core';
import { RolAdmin } from '../../models/core';

export interface Menu {
  state: string[];
  name: string;
  type: string;
  icon: string;
}

const MENUITEMS_ROLADMIN = [
  { state: ['/', 'home'], name: 'Home', type: 'link', icon: 'home' },
  { state: ['/', 'users'], name: 'Users', type: 'link', icon: 'perm_contact_calendar' },
  { state: ['/', 'groups'], name: 'Groups', type: 'link', icon: 'date_range' },
];

@Injectable()
export class MenuItems {
  menu: Menu[] = [];

  constructor(private session: SessionService) {
    this.menu.push(
      { state: ['/', 'home'], name: 'Home', type: 'link', icon: 'home' }
    );

    const groups = this.session.GetGroupAdmins();
    for ( let i = 0; i < groups.length; i++ ) {
      const item: Menu = {
        state: ['/', 'groups', 'showone', `${groups[i].GroupID}`],
        // state: `groups/showone/${groups[i].GroupID}`,
        name: groups[i].GroupName,
        type: 'link',
        icon: 'view_list',
      };
      this.menu.push(item);
    }
  }

  getMenuitem(): Menu[] {
    if ( this.session.GetUser().Rol === RolAdmin ) {
      return MENUITEMS_ROLADMIN;
    } else {
      return this.menu;
    }
  }
}
