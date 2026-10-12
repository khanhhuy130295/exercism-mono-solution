
export const colorCode = (color: string) => {
    const idx: number = COLORS.findIndex(value => value == color)
    return idx
}

export const COLORS: string[] = [
    "black",
    "brown",
    "red",
    "orange",
    "yellow",
    "green",
    "blue",
    "violet",
    "grey",
    "white"
]
