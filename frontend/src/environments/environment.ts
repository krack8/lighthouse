const LIGHTHOUSE_API_ENDPOINT = 'http://localhost:8082';

export const environment = {
  production: false,

  apiEndPoint: LIGHTHOUSE_API_ENDPOINT + '/api',
  multiClusterWsEndpoint: LIGHTHOUSE_API_ENDPOINT + '/clusterlog'
}; 