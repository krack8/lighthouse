// new apis
export const NAMESPACE = '/v1/namespace';
export const NAMESPACE_DEPLOYMENT = '/v1/deployment';
export const NAMESPACE_DEPLOYMENT_STATS = '/v1/deployment/stats';
export const NAMESPACE_POD = '/v1/pod';
export const NAMESPACE_POD_STATS = '/v1/pod/stats';
export const NAMESPACE_REPLICA_SET = '/v1/replicaset';
export const NAMESPACE_STATEFUL_SET = '/v1/statefulset';
export const NAMESPACE_STATEFUL_SET_STATS = '/v1/statefulset/stats';
export const NAMESPACE_DAEMON_SET = '/v1/daemonset';
export const NAMESPACE_DAEMON_SET_STATS = '/v1/daemonset/stats';

export const NAMESPACE_SECRET = '/v1/secret';
export const NAMESPACE_CONFIG_MAP = '/v1/config-map';

export const NAMESPACE_SERVICE_ACCOUNT = '/v1/service-account';
export const NAMESPACE_SERVICE = '/v1/service';

export const NAMESPACE_INGRESS = '/v1/ingress';
export const NAMESPACE_CERTIFICATE = '/v1/certificate';
export const NAMESPACE_ROLE_BINDING = '/v1/role-binding';
export const NAMESPACE_ROLE = '/v1/role';
export const NAMESPACE_JOB = '/v1/job';
export const NAMESPACE_CRON_JOB = '/v1/cronjob';
export const NAMESPACE_NETWORK_POLICY = '/v1/network-policy';
export const NAMESPACE_RESOURCE_QUOTA = '/v1/resource-quota';
export const NAMESPACE_PERSISTENT_VOLUME = '/v1/pvc';
export const NAMESPACE_GATEWAY = '/v1/gateway';
export const NAMESPACE_VIRTUAL_SERVICE = '/v1/virtual-service';
export const NAMESPACE_NAME_LIST = '/v1/namespace/names';

export const GET_LOGS = '/log-metrics/container';
export const GET_LOGS_V1 = '/v1/pod/logs/';

export const NAMESPACE_DEPLOYMENT_POD_LIST = '/v1/deployment/{0}/pods';
export const NAMESPACE_STATEFUL_SETS_POD_LIST = '/v1/statefulset/{0}/pods';

export const ATHENTICATE_USER = '/v1/authenticate';
export const NAMESPACE_ENDPOINTS = '/v1/endpoints';
export const NAMESPACE_ENDPOINT_SLICE = '/v1/endpoint-slice';
export const NAMESPACE_PDB = '/v1/PDB';
export const NAMESPACE_CONTROLLER_REVISION = '/v1/controller-revision';
export const NAMESPACE_REPLICATION_CONTROLLER = '/v1/replication-controller';

export const NAMESPACE_EVENTS = '/v1/event';
export const POD_TERMINAL_URL = '/ws/v1/pod/{0}/exec';
