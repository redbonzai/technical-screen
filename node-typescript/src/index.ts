import { sort } from "./sort";

console.log("STANDARD", sort(90, 90, 90, 10));
console.log("SPECIAL", sort(200, 200, 200, 10));
console.log("SPECIAL", sort(95, 95, 95, 25));
console.log("REJECTED", sort(200, 200, 200, 25));
