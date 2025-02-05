export function flattenDeep<T>(array: T[]): T[] {
  return array.reduce((acc: T[], val) => (Array.isArray(val) ? acc.concat(flattenDeep(val)) : acc.concat(val)), []);
}
