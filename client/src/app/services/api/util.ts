import { URLSearchParams } from '@angular/http';
import { isNullOrUndefined } from 'util';

export class Paginator {
  constructor(public offset: number,
              public limit: number) {}

  GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if (!isNullOrUndefined(this.limit) && this.limit !== 0 ) {
      usp.append('limit', this.limit.toString());
    }
    if (!isNullOrUndefined(this.offset) ) {
      usp.append('offset', this.offset.toString());
    }
    return usp;
  }
}

export class OrderBy {
  constructor(public by: string,
              public desc: boolean) {}

  GetUSP(): URLSearchParams {
    let usp: URLSearchParams;
    usp = new URLSearchParams();
    if ( !isNullOrUndefined(this.by) && this.by !== '' ) {
      usp.append('orderby', this.by);
    }
    if ( !isNullOrUndefined(this.desc) !== null) {
      usp.append('desc', this.desc.toString());
    }
    return usp;
  }
}
