export function toParams(query: string) {
  const q = query.replace(/^\??\//, '');

  return q.split('&').reduce((values: any, param) => {
    const [key, value] = param.split('=');

    values[key] = value;

    return values;
  }, {});
}
