import { Injectable } from '@angular/core';

import { LocalStorageService } from './local-storage.service';
import { Session, User, GroupAdmin } from '../models/core';

const ssname = 'session-info';

@Injectable()
export class SessionService {
  private session: Session;

  constructor(private storage: LocalStorageService) {
    if ( !this.Open(storage.getItem(ssname)) ) {
      this.Close();
    }
  }

  public Open(session: Session): boolean {
    if ( session === null || !session ) {
      return false;
    }

    this.storage.setItem(ssname, session);
    this.session = session;
    return true;
  }

  public Close(): void {
    this.storage.removeItem(ssname);
    this.session = null;
  }

  public IsOpen(): boolean {
    return !(this.session === null || !this.session);
  }

  public GetToken(): string {
    if (!this.IsOpen()) {
      return '';
    }
    return this.session.Token;
  }

  public GetUser(): User {
    if (!this.IsOpen()) {
      return new User(0, '', '', '', '', false);
    }
    return this.session.User;
  }

  public GetGroupAdmins(): GroupAdmin[] {
    if (!this.IsOpen()) {
      return [];
    }
    return this.session.GroupAdmins;
  }

  public HasPermission(permission: string): boolean {
    if (!this.IsOpen()) {
      return false;
    }
    for ( let i = 0; i < this.session.Permissions.length; i++ ) {
      if (this.session.Permissions[i] === permission) {
        return true;
      }
    }
    return false;
  }
}
