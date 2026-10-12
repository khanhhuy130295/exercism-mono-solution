
export function twoFer(name?: string): string {
  let result: string = "One for you, one for me."  
  if (name && name.length > 0){
    result  =`One for ${name}, one for me.`
  }
  return result;
}
