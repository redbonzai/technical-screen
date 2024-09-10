export function sort(width: number, height: number, length: number, mass: number): string {
    if (width <= 0 || height <= 0 || length <= 0) {
        throw new Error("Invalid input: dimensions must be positive");
    }
    if (mass < 0) {
        throw new Error("Invalid input: mass must be positive");
    }

    const volume = width * height * length;
    const isBulky = volume > 1000000 || width >= 150 || height >= 150 || length >= 150; // Changed to > 1,000,000
    const isHeavy = mass >= 20;

    if (isBulky && isHeavy) {
        return "REJECTED";
    } else if (isBulky || isHeavy) {
        return "SPECIAL";
    } else {
        return "STANDARD";
    }
}
