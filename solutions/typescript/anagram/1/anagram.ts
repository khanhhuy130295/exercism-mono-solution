
export class Anagram {
  constructor(private readonly subject: string) {}

  public matches(...potentials: string[]): string[] {
    return potentials.filter(potential =>
      this.checkAnagrams(this.subject, potential)
    );
  }

  private checkAnagrams(str1: string, str2: string): boolean {
    const a = str1.toLowerCase();
    const b = str2.toLowerCase();

    if (a === b || a.length !== b.length) {
      return false;
    }

    const seen = new Map<string, number>();

    for (const char of a) {
      seen.set(char, (seen.get(char) ?? 0) + 1);
    }

    for (const char of b) {
      const count = seen.get(char) ?? 0;

      if (count === 0) {
        return false;
      }

      seen.set(char, count - 1);
    }

    return true;
  }
}
