"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const sort_1 = require("./sort");
describe('Package Sorting Function', () => {
    test('should return STANDARD for non-bulky and non-heavy package', () => {
        expect((0, sort_1.sort)(90, 90, 90, 10)).toBe("STANDARD");
    });
    test('should return SPECIAL for bulky package', () => {
        expect((0, sort_1.sort)(200, 200, 200, 10)).toBe("SPECIAL");
    });
    test('should return SPECIAL for heavy package', () => {
        expect((0, sort_1.sort)(95, 95, 95, 25)).toBe("SPECIAL");
    });
    test('should return REJECTED for both bulky and heavy package', () => {
        expect((0, sort_1.sort)(200, 200, 200, 25)).toBe("REJECTED");
    });
    test('should return SPECIAL for package with one dimension >= 150cm', () => {
        expect((0, sort_1.sort)(160, 50, 50, 10)).toBe("SPECIAL");
    });
});
//# sourceMappingURL=sort.spec.js.map