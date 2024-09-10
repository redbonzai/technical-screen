import { sort } from "./sort";

// Standard packages (neither bulky nor heavy)
console.log("STANDARD", sort(90, 90, 90, 10)); // STANDARD
console.log("STANDARD", sort(50, 50, 50, 5)); // STANDARD
console.log("STANDARD", sort(100, 100, 100, 15)); // STANDARD

// Special packages (bulky but not heavy)
console.log("SPECIAL", sort(200, 100, 100, 10)); // SPECIAL
console.log("SPECIAL", sort(150, 90, 90, 10)); // SPECIAL
console.log("SPECIAL", sort(90, 150, 90, 15)); // SPECIAL
console.log("SPECIAL", sort(90, 90, 150, 15)); // SPECIAL
console.log("SPECIAL", sort(100, 100, 1000, 15)); // SPECIAL (Volume = 1,000,000 cm³)

// Special packages (heavy but not bulky)
console.log("SPECIAL", sort(90, 90, 90, 20)); // SPECIAL (Mass is exactly 20 kg)
console.log("SPECIAL", sort(95, 95, 95, 25)); // SPECIAL (Mass is greater than 20 kg, but not bulky)

// Rejected packages (both bulky and heavy)
console.log("REJECTED", sort(200, 200, 200, 25)); // REJECTED
console.log("REJECTED", sort(150, 150, 150, 21)); // REJECTED (Dimensions exactly 150 cm, mass over 20 kg)

// Edge cases
try {
    console.log(sort(0, 90, 90, 10)); // Error (Invalid input: zero dimension)
} catch (error: any) {
    console.error(error.message);
}

try {
    console.log(sort(90, 90, 90, -5)); // Error (Invalid input: negative mass)
} catch (error: any) {
    console.error(error.message);
}

try {
    console.log(sort(90, 0, 90, 10)); // Error (Invalid input: zero dimension)
} catch (error: any) {
    console.error(error.message);
}

try {
    console.log(sort(90, 90, 90, 0)); // STANDARD (Mass = 0 is technically not heavy or bulky)
} catch (error: any) {
    console.error(error.message);
}
