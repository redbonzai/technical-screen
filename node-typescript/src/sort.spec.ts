import { sort } from "./sort";

describe("sort function", () => {
    // Standard packages (neither bulky nor heavy)
    it("should classify standard packages correctly", () => {
        expect(sort(90, 90, 90, 10)).toBe("STANDARD");
        expect(sort(50, 50, 50, 5)).toBe("STANDARD");
        expect(sort(95, 95, 95, 15)).toBe("STANDARD");
    });

    // Special packages (bulky but not heavy)
    it("should classify bulky but not heavy packages as SPECIAL", () => {
        expect(sort(200, 100, 100, 10)).toBe("SPECIAL");
        expect(sort(150, 90, 90, 10)).toBe("SPECIAL");
        expect(sort(90, 150, 90, 15)).toBe("SPECIAL");
        expect(sort(90, 90, 150, 15)).toBe("SPECIAL");
        expect(sort(100, 100, 1000, 15)).toBe("SPECIAL"); // Volume = 1,000,000 cm³
    });

    // Special packages (heavy but not bulky)
    it("should classify heavy but not bulky packages as SPECIAL", () => {
        expect(sort(90, 90, 90, 20)).toBe("SPECIAL"); // Mass is exactly 20 kg
        expect(sort(95, 95, 95, 25)).toBe("SPECIAL"); // Mass is greater than 20 kg, but not bulky
    });

    // Rejected packages (both bulky and heavy)
    it("should reject packages that are both bulky and heavy", () => {
        expect(sort(200, 200, 200, 25)).toBe("REJECTED");
        expect(sort(150, 150, 150, 21)).toBe("REJECTED"); // Dimensions exactly 150 cm, mass over 20 kg
    });

    // Edge cases
    it("should handle edge cases and invalid inputs", () => {
        // Zero or negative dimensions should throw errors
        expect(() => sort(0, 90, 90, 10)).toThrow(
            "Invalid input: dimensions must be positive",
        );
        expect(() => sort(90, 90, 90, -5)).toThrow(
            "Invalid input: mass must be positive",
        );
        expect(() => sort(90, 0, 90, 10)).toThrow(
            "Invalid input: dimensions must be positive",
        );

        // Mass of 0 is technically allowed and considered STANDARD
        expect(sort(90, 90, 90, 0)).toBe("STANDARD");
    });
});
