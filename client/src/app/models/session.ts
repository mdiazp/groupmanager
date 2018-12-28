import { User as user } from './user';
import { GroupAdmin } from './groupAdmin';

export class Session {
    constructor(public Token: string,
                public User: user,
                public GroupAdmins: GroupAdmin[],
                public Permissions: string[]) {}
}
