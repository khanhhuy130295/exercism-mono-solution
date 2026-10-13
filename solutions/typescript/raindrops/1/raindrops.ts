
const factors: {value: number, sound: string}[] = [
  { value:3, sound:"Pling"},
  { value:5, sound:"Plang"},
  { value:7, sound:"Plong"},
  
]
export function convert(count: number):string {
  let result:string = "";
  for (const factor of factors){
    if (count % factor.value === 0){
      result += factor.sound
    }
  }
  if (result.length == 0){
    return `${count}`
  }
  return result;
}
