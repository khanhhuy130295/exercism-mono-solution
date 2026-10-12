export function isLeap(year: number): boolean {
    let is_leap: boolean = false;
    if (year % 4 === 0) is_leap = true;
    if (year % 100 === 0) is_leap = false;
    if (year % 400 === 0) is_leap = true;

    return is_leap;
}