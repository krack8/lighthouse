export function flattenObject(obj: { [key: string | number]: any }, parentKey = '', result = {}) {
  // Iterate through each property in the object
  for (const key in obj) {
    if (obj.hasOwnProperty(key)) {
      // Create a new key combining the parent key with the current key
      const newKey = parentKey ? `${parentKey}.${key}` : key;

      // Check if the value is an object
      if (typeof obj[key] === 'object' && obj[key] !== null) {
        // Handle arrays and objects recursively
        if (Array.isArray(obj[key])) {
          obj[key].forEach((item, index) => {
            if (typeof item === 'string') {
              result[`${newKey}[${index}]`] = item;
            } else {
              flattenObject(item, `${newKey}[${index}]`, result);
            }
          });
        } else {
          flattenObject(obj[key], newKey, result);
        }
      } else {
        // Otherwise, directly assign the key-value pair to the result
        result[newKey] = obj[key];
      }
    }
  }
  return result;
}
