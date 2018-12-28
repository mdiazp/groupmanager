import { Injectable } from '@angular/core';

@Injectable()
export class LocalStorageService {

    private prefix = 'xum-app-';

    constructor() {}

    setItem(key: string, value: any) {
        localStorage.setItem(`${this.prefix}${key}`, JSON.stringify(value));
    }

    getItem(key: string): any {
        try {
        return JSON.parse(localStorage.getItem(`${this.prefix}${key}`));
        } catch (error) {
        console.log(error);
        }
    }

    removeItem(key: string) {
        localStorage.removeItem(`${this.prefix}${key}`);
  }
}
