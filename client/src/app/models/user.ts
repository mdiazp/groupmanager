export class User {
    constructor(public ID: number,
                public Provider: string,
                public Username: string,
                public Name: string,
                public Rol: string,
                public Enabled: boolean) {}
}
