/**
 * K8s routes resources map
 * This map needs to be updated if any new resource is added in the Lighthouse in future
 */
export const k8sRoutesMap = new Map<string, string>([
  ['ConfigMap', 'config-maps'],
  ['PersistentVolumeClaim', 'pvcs'],
  ['Pod', 'pods'],
  ['ResourceQuota', 'resource-quota'],
  ['Secret', 'secrets'],
  ['ServiceAccount', 'service-accounts'],
  ['Service', 'service'],
  ['DaemonSet', 'daemon-sets'],
  ['Deployment', 'deployments'],
  ['ReplicaSet', 'replica-sets'],
  ['StatefulSet', 'stateful-sets'],
  ['CronJob', 'cron-job'],
  ['Job', 'job'],
  ['Certificate', 'certificates'],
  ['Ingress', 'ingresses'],
  ['NetworkPolicy', 'network-policy'],
  ['RoleBinding', 'role-binding'],
  ['Role', 'role'],
  ['ControllerRevision', 'controller-revision'],
  ['PodDisruptionBudget', 'PDB'],
  ['ReplicationController', 'replication-controller'],
  ['Endpoints', 'endpoints'],
  ['EndpointSlice', 'endpoints-slice'],
  ['VirtualService', 'virtual-service']
]);

/**
 * K8s routes permission map
 * This map contains the permission name for each resource routes in lighthouse namespace
 */
export const k8sRoutesPermissionMap = new Map<string, string>([
  ['config-maps', 'VIEW_NAMESPACE_CONFIG_MAP'],
  ['pvcs', 'VIEW_NAMESPACE_PERSISTENT_VOLUME'],
  ['pods', 'VIEW_NAMESPACE_POD'],
  ['resource-quota', 'VIEW_NAMESPACE_RESOURCE_QUOTA'],
  ['secrets', 'VIEW_NAMESPACE_SECRET'],
  ['service-accounts', 'VIEW_NAMESPACE_SERVICE_ACCOUNT'],
  ['service', 'VIEW_NAMESPACE_SERVICE'],
  ['daemon-sets', 'VIEW_NAMESPACE_DAEMON_SET'],
  ['deployments', 'VIEW_NAMESPACE_DEPLOYMENT'],
  ['replica-sets', 'VIEW_NAMESPACE_REPLICA_SET'],
  ['stateful-sets', 'VIEW_NAMESPACE_STATEFUL_SET'],
  ['cron-job', 'VIEW_NAMESPACE_CRON_JOB'],
  ['job', 'VIEW_NAMESPACE_JOB'],
  ['certificates', 'VIEW_NAMESPACE_CERTIFICATE'],
  ['ingresses', 'VIEW_NAMESPACE_INGRESS'],
  ['network-policy', 'VIEW_NAMESPACE_NETWORK_POLICY'],
  ['role-binding', 'VIEW_NAMESPACE_ROLE_BINDING'],
  ['role', 'VIEW_NAMESPACE_ROLE'],
  ['controller-revision', 'VIEW_NAMESPACE_CONTROLLER_REVISION'],
  ['PDB', 'VIEW_NAMESPACE_PDB'],
  ['replication-controller', 'VIEW_NAMESPACE_REPLICATION_CONTROLLER'],
  ['endpoints', 'VIEW_NAMESPACE_ENDPOINTS'],
  ['endpoints-slice', 'VIEW_NAMESPACE_ENDPOINT_SLICE'],
  ['virtual-service', 'VIEW_NAMESPACE_VIRTUAL_SERVICE']
]);
