"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.sort = sort;
function sort(width, height, length, mass) {
    const volume = width * height * length;
    const isBulky = volume >= 1000000 || width >= 150 || height >= 150 || length >= 150;
    const isHeavy = mass >= 20;
    if (isBulky && isHeavy) {
        return "REJECTED";
    }
    else if (isBulky || isHeavy) {
        return "SPECIAL";
    }
    else {
        return "STANDARD";
    }
}
//# sourceMappingURL=sort.js.map