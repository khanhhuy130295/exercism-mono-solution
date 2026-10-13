export class Gigasecond {

  constructor(private _time_convert: Date){
  }
  
  public date() {
     const gigasecondInMilliseconds = 1_000_000_000 * 1000;
    return new Date(this._time_convert.getTime()  + gigasecondInMilliseconds);
  }
}
