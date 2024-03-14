function expandedForm(num) {
    // Your code here
    let numStr = num.toString();
    let result = [];
    for (let i = 0; i < numStr.length; i++) {
        if (numStr[i] !== '0') {
            result.push(numStr[i] + '0'.repeat(numStr.length - i - 1));
        }
    }
    return result.join(' + ');
}