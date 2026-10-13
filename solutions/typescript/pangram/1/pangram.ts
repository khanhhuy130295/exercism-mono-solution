export function isPangram(input:string) {
 let isPangram = false;

  const seen = new Set<string>();

  for (const char of input.toLowerCase()) {
    if (char < "a" || char > "z") {
      continue;
    }

    seen.add(char);

    if (seen.size === 26) {
      isPangram = true;
      break;
    }
  }

  return isPangram;
}
