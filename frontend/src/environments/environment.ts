const KLOVERCLOUD_API_ENDPOINT = 'http://localhost:8080';

export const environment = {
  production: false,

  apiEndPoint: KLOVERCLOUD_API_ENDPOINT + '/api',
  multiClusterWsEndpoint: KLOVERCLOUD_API_ENDPOINT + '/clusterlog'
};
